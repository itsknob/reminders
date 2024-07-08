package dao

import (
	"database/sql"
	"fmt"

	r "example/hello/reminders"
	_ "modernc.org/sqlite"
)

type IReminderDao interface {
	GetReminder() (r.Reminder, error)
}

type ReminderDao struct {
	Database *sql.DB
}

func (reminderDao *ReminderDao) New() ReminderDao {
	db, err := sql.Open("sqlite", "reminders.db")
	if err != nil {
		panic(fmt.Sprintf("DB Could not be opened - %+v", err))
	}
	return ReminderDao{
		Database: db,
	}
}

func (reminderDao ReminderDao) GetReminder(id int) (r.Reminder, error) {
	return r.Reminder{}, nil
}

func (reminderDao *ReminderDao) GetAllReminders() ([]r.Reminder, error) {
	return []r.Reminder{}, nil
}

func (reminderDao *ReminderDao) DeleteReminder(id int) error {
	return nil
}

func (reminderDao *ReminderDao) SaveReminder(reminder r.Reminder) (r.Reminder, error) {
	return r.Reminder{}, nil
}
