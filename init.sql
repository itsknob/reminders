CREATE TABLE schedules IF NOT EXISTS (
    id INT PRIMARY_KEY,
    created_date DATETIME,
    start_date DATETIME,
    end_date DATETIME,
    time_hour  INT,
    time_minute INT,
    repeat_seconds INT,
    repeat_minutes INT,
    repeat_hours   INT,
    repeat_days    INT,
    repeat_weeks   INT,
    repeat_months  INT,
);

CREATE TABLE reminders IF NOT EXISTS (
    id INT PRIMARY_KEY,
    title VARCHAR(64),
    description VARCHAR(255),
    completed,
    schedule_id INT,
    FOREIGN KEY (schedule_id) REFERENCES schedules(id)
);

INSERT INTO schedules(id, time_hour, time_minute, repeat_id, created_date, start_date, end_date, repeat_seconds, repeat_minutes, repeat_hours, repeat_days, repeat_weeks, repeat_months)
    VALUES
        ("0", "2024-07-20", "2024-07-21", "2024-07-22", 0, 0, 30, 0, 0, 0, 0, 0, 0);

INSERT INTO reminders(id, type, description, completed, schedule_id)
    VALUES
        ("0", "Clean", "Clean stuff", false, null)
        ("1", "Wipe", "Wipe stuff", false, null)
        ("2", "Dust", "Dust stuff", false, null);
