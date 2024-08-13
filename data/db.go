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
		log.Printf("Failed to update reminder: %+v\n", err)
		return &models.Reminder{}, err
	}

	log.Printf("db - CompleteReminder - &updatedReminder: %+v", &updatedReminder)

	return &updatedReminder, err
}

func GetReminders() ([]models.Reminder, error) {
	reminders := []models.Reminder{}

	rows, err := remindersDb.Query("select * from reminders")
	defer rows.Close()

	if err != nil {
		log.Printf("Failed to get reminders: %+v\n", err)
		return reminders, err
	}

	for rows.Next() {

		var reminder models.Reminder
		err = rows.Scan(
			&reminder.Id, &reminder.Title, &reminder.Description,
			&reminder.Completed, &reminder.CreatedDate, &reminder.StartDate,
			&reminder.EndDate, &reminder.TimeHour, &reminder.TimeMinute,
			&reminder.RepeatSeconds, &reminder.RepeatMinutes, &reminder.RepeatHours,
			&reminder.RepeatDays, &reminder.RepeatWeeks, &reminder.RepeatMonths,
		)

		if err != nil {
			log.Printf("Failed to scan row while getting reminders: %+v\n", err)
			return reminders, err
		}

		log.Printf("GetReminders - Reminder %d - %+v", reminder.Id, reminder)

		reminders = append(reminders, reminder)
	}

	return reminders, err
}
