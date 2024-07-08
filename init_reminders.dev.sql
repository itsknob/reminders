PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE reminders(id int, name text, description text, schedule REFERENCES schedules (id));
INSERT INTO reminders VALUES(1,'Clean','Clean all the things',1);
CREATE TABLE schedules(id int, timeMinute int, timeHour int, repeatSeconds int, repeatMinutes int, repeatHours int, repeatDays int, repeatWeeks int, repeatMonths int);
INSERT INTO schedules VALUES(1,0,9,30,0,0,0,0,0);
COMMIT
