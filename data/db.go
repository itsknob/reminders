package data

import (
	"database/sql"
	// "fmt"
	"log"
	// "time"

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

// func CreateSchedule(schedule *models.SchedulePostBody) (*models.Schedule, error) {
// 	fmt.Printf("db - CreateSchedule - SchedulePostBody %+v\n", schedule)
//
// 	newSchedule := models.Schedule{}
// 	err := remindersDb.QueryRow(`
//         insert into schedules
//         values (NULL, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
//         returning *;`,
// 		time.Now().UTC().String(), // created_date
// 		time.Now().UTC().String(), // start_date
// 		"",
// 		&schedule.TimeHour,
// 		&schedule.TimeMinute,
// 		&schedule.RepeatSeconds,
// 		&schedule.RepeatMinutes,
// 		&schedule.RepeatHours,
// 		&schedule.RepeatDays,
// 		&schedule.RepeatWeeks,
// 		&schedule.RepeatMonths,
// 	).Scan(
// 		&newSchedule.Id,
// 		&newSchedule.CreatedDate,
// 		&newSchedule.StartDate,
// 		&newSchedule.EndDate,
// 		&newSchedule.TimeHour,
// 		&newSchedule.TimeMinute,
// 		&newSchedule.RepeatSeconds,
// 		&newSchedule.RepeatMinutes,
// 		&newSchedule.RepeatHours,
// 		&newSchedule.RepeatDays,
// 		&newSchedule.RepeatWeeks,
// 		&newSchedule.RepeatMonths,
// 	)
//
// 	fmt.Printf("db - CreateSchedule - newSchedule - %+v\n", newSchedule)
//
// 	if err != nil {
// 		fmt.Printf("db - CreateSchedule - err: %+v\n", err)
// 		panic(err)
// 	}
//
// 	return &newSchedule, err // either err will be nil, or newschedule will be nil
// }

func CreateReminder(reminder *models.ReminderPostBody) (*models.Reminder, error) {

	row := remindersDb.QueryRow(`
        INSERT INTO reminders(
            title, completed, description, created_date,
            start_date, end_date, time_hour, time_minute, repeat_seconds,
            repeat_minutes, repeat_hours, repeat_days, repeat_weeks, repeat_months
        )
        VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
        RETURNING *`,
		reminder.Title,
		reminder.Completed,
		reminder.Description,
		reminder.CreatedDate,
		reminder.StartDate,
		reminder.EndDate,
		reminder.TimeHour,
		reminder.TimeMinute,
		reminder.RepeatSeconds,
		reminder.RepeatMinutes,
		reminder.RepeatHours,
		reminder.RepeatDays,
		reminder.RepeatWeeks,
		reminder.RepeatMonths,
	)

	var newReminder models.Reminder
	err := row.Scan(
		&newReminder.Id,
		&newReminder.Title,
		&newReminder.Description,
		&newReminder.Completed,
		&newReminder.CreatedDate,
		&newReminder.StartDate,
		&newReminder.EndDate,
		&newReminder.TimeHour,
		&newReminder.TimeMinute,
		&newReminder.RepeatSeconds,
		&newReminder.RepeatMinutes,
		&newReminder.RepeatHours,
		&newReminder.RepeatDays,
		&newReminder.RepeatWeeks,
		&newReminder.RepeatMonths,
	)

	if err != nil {
		return nil, err
	}

	log.Printf("Created Reminder %d with Title %s", newReminder.Id, newReminder.Title)

	return &newReminder, nil
}

func FindReminder(id string) (*models.Reminder, error) {
	var reminder models.Reminder
	row := remindersDb.QueryRow("SELECT * FROM reminders WHERE id=?", id)
	err := row.Scan(&reminder.Id, &reminder.Title, &reminder.Description,
		&reminder.Completed, &reminder.CreatedDate, &reminder.StartDate,
		&reminder.EndDate, &reminder.TimeHour, &reminder.TimeMinute,
		&reminder.RepeatSeconds, &reminder.RepeatMinutes,
		&reminder.RepeatHours, &reminder.RepeatDays, &reminder.RepeatWeeks,
		&reminder.RepeatMonths)
	return &reminder, err
}

func CompleteReminder(id string) (*models.Reminder, error) {
	reminder, err := FindReminder(id)
	if err != nil {
		return nil, err
	}

	log.Printf("db - CompleteReminder - reminder: \n%+v\n", reminder)

	row := remindersDb.QueryRow("UPDATE reminders SET completed=? WHERE id=? RETURNING *", true, reminder.Id)

	var updatedReminder models.Reminder
	err = row.Scan(
		&updatedReminder.Id,
		&updatedReminder.Title,
		&updatedReminder.Description,
		&updatedReminder.Completed,
		&updatedReminder.CreatedDate,
		&updatedReminder.StartDate,
		&updatedReminder.EndDate,
		&updatedReminder.TimeHour,
		&updatedReminder.TimeMinute,
		&updatedReminder.RepeatSeconds,
		&updatedReminder.RepeatMinutes,
		&updatedReminder.RepeatHours,
		&updatedReminder.RepeatDays,
		&updatedReminder.RepeatWeeks,
		&updatedReminder.RepeatMonths,
	)

	if err != nil {
		panic(err)
	}

	log.Printf("db - CompleteReminder - &updatedReminder: %+v", &updatedReminder)

	return &updatedReminder, err
}

func GetReminders() ([]models.Reminder, error) {
	reminders := []models.Reminder{}

	log.Println("db - GetReminders - start")

	rows, err := remindersDb.Query("select * from reminders")
	defer rows.Close()

	log.Println("db - GetReminders - queried")

	if err != nil {
		panic(err)
		// return reminders, err
	}

	log.Println("db - GetReminders - no err")

	for rows.Next() {
		log.Println("db - GetReminders - rows.Next")

		var reminder models.Reminder
		err = rows.Scan(
			&reminder.Id, &reminder.Title, &reminder.Description,
			&reminder.Completed, &reminder.CreatedDate, &reminder.StartDate,
			&reminder.EndDate, &reminder.TimeHour, &reminder.TimeMinute,
			&reminder.RepeatSeconds, &reminder.RepeatMinutes, &reminder.RepeatHours,
			&reminder.RepeatDays, &reminder.RepeatWeeks, &reminder.RepeatMonths,
		)

		log.Println("db - GetReminders - rows.Next - scanned")

		if err != nil {
			panic(err)
		}

		log.Printf("reminder: \n%+v\n", &reminder)
		log.Printf("GetReminders - Reminder %d - %+v", reminder.Id, reminder)
		reminders = append(reminders, reminder)
	}

	return reminders, err
}
