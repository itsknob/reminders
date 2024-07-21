package data

import (
	"database/sql"
	"log"

	"knob.dev/reminders/models"
	_ "modernc.org/sqlite"
)

var (
    remindersDb *sql.DB
)

/** Special Function Type */
func init() {
    db, err := sql.Open("sqlite", "reminders.db")
    remindersDb = db
    if err != nil {
        panic(err)
    }
}

func CreateSchedule(schedule *models.SchedulePostBody) (*models.Schedule, error) {
    row := remindersDb.QueryRow(`
        insert into schedules(
            TimeHour,
            TimeMinute,
            RepeatSeconds,
            RepeatMinutes,
            RepeatHours,
            RepeatDays,
            RepeatWeeks,
            RepeatMonths,
        ) values (?, ?, ?, ?, ?, ?, ?, ?, ?)`,
            schedule.TimeHour,
            schedule.TimeMinute,
            schedule.RepeatSeconds,
            schedule.RepeatMinutes,
            schedule.RepeatHours,
            schedule.RepeatDays,
            schedule.RepeatWeeks,
            schedule.RepeatMonths,
        )

    var newSchedule *models.Schedule = nil
    err := row.Scan(&newSchedule)
    return newSchedule, err // either err will be nil, or newschedule will be nil
}

func CreateReminder(reminder *models.ReminderPostBody, schedule *models.Schedule) (*models.Reminder, error) {
    row := remindersDb.QueryRow(`
        INSERT INTO reminders(Title, Completed, Schedule)
        VALUES(?, ?, ?)`,
        reminder.Title,
        reminder.Completed,
        reminder.Description,
        schedule,
    )
    newReminder := models.Reminder{}
    err := row.Scan(newReminder)
    if err != nil {
        return nil, err
    }

    log.Printf("Created Reminder %s with Title %s", newReminder.Id, newReminder.Title)

    return &newReminder, nil
}

func FindReminder(id string) (*models.Reminder, error) {
    reminder := &models.Reminder{}
    rows := remindersDb.QueryRow("SELECT * FROM reminders WHERE id=?", id)
    err := rows.Scan(reminder)
    if err != nil {
        return nil, err
    }
    return reminder, nil
}

func CompleteReminder(id string) (*models.Reminder, error) {
    reminder, err := FindReminder(id)
    if err != nil {
        return nil, err
    }
    row := remindersDb.QueryRow("UPDATE reminders SET Completed = ? WHERE id=?", true, reminder.Id)
    updatedReminder := models.Reminder{}
    row.Scan(updatedReminder)

    return &updatedReminder, nil

}
func GetReminders() (*[]models.Reminder, error) {
    reminders := &[]models.Reminder{}
    rows, err := remindersDb.Query("select * from reminders")
    if err != nil {
        return reminders, err
    }
    err = rows.Scan(reminders)
    log.Printf("GetReminders - %+v", reminders)
    return reminders, err
}
