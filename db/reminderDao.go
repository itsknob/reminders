package dao

import (
	"database/sql"
	"errors"
	"fmt"

	r "example/hello/reminders"

	_ "modernc.org/sqlite"
)

type IReminderDao interface {
	New() ReminderDao
	GetReminder() (r.Reminder, error)
}

type ReminderDao struct {
	Database *sql.DB
}

type Dao interface {
}

func newReminderDao(db *sql.DB) ReminderDao {
	return ReminderDao{
		Database: db,
	}
}

func GetDao(daoType string) (Dao, error) {
	db, err := sql.Open("sqlite", "reminders.db")
	if err != nil {
		panic(fmt.Sprintf("DB Could not be opened - %+v", err))
	}
	if daoType == "reminder" {
		return newReminderDao(db), nil
	}
	// add more types here
	return nil, errors.New(fmt.Sprintf("No Dao Type for: %s", daoType))
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
