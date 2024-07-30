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
    err := row.Scan(
            &schedule.TimeHour,
            &schedule.TimeMinute,
            &schedule.RepeatSeconds,
            &schedule.RepeatMinutes,
            &schedule.RepeatHours,
            &schedule.RepeatDays,
            &schedule.RepeatWeeks,
            &schedule.RepeatMonths,
        )

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
    var newReminder models.Reminder
    err := row.Scan(newReminder)
    if err != nil {
        return nil, err
    }

    log.Printf("Created Reminder %s with Title %s", newReminder.Id, newReminder.Title)

    return &newReminder, nil
}

func FindReminder(id string) (*models.Reminder, error) {
    var reminder models.Reminder
    row := remindersDb.QueryRow("SELECT * FROM reminders WHERE id=?", id)
    err := row.Scan(&reminder.Id, &reminder.Title, &reminder.Description, &reminder.Completed, &reminder.Schedule.Id)
    return &reminder, err
}

func CompleteReminder(id string) (*models.Reminder, error) {
    reminder, err := FindReminder(id)
    if err != nil {
        return nil, err
    }
    row := remindersDb.QueryRow("UPDATE reminders SET completed=? WHERE id=? RETURNING *", true, reminder.Id)
    var updatedReminder models.Reminder
    err = row.Scan(&updatedReminder.Id, &updatedReminder.Title, &updatedReminder.Description, &updatedReminder.Completed, &updatedReminder.Schedule.Id)

    log.Printf("db - CompleteReminder - &updatedReminder: %+v", &updatedReminder)

    return &updatedReminder, err

}

func GetReminders() ([]models.Reminder, error) {
    reminders := []models.Reminder{}

    rows, err := remindersDb.Query("select * from reminders")
    defer rows.Close()

    if err != nil {
        return reminders, err
    }

    for rows.Next() {
        var reminder models.Reminder
        err = rows.Scan(&reminder.Id, &reminder.Title, &reminder.Description, &reminder.Completed, &reminder.Schedule.Id)
        log.Printf("GetReminders - Reminder %s - %+v", reminder.Id, reminder)
        reminders = append(reminders, reminder)
    }

    return reminders, err
}
