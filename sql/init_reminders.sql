PRAGMA foreign_keys=on;
BEGIN TRANSACTION;
CREATE TABLE reminders(id int, name text, description text, created_date timestamp, start_date timestamp, end_date timestamp, timeMinute int, timeHour int, repeatSeconds int, repeatMinutes int, repeatHours int, repeatDays int, repeatWeeks int, repeatMonths int);
INSERT INTO reminders VALUES(1,'Clean','Clean all the things',datetime('now'),datetime('now'), datetime('now'),0,0,30,0,0,0,0,0);
COMMIT
