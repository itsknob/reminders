CREATE TABLE IF NOT EXISTS reminders (
    id INTEGER PRIMARY KEY NOT NULL,
    title TEXT,
    description TEXT,
    completed,
    created_date TEXT,
    start_date TEXT,
    end_date TEXT,
    time_hour  INT,
    time_minute INT,
    repeat_seconds INT,
    repeat_minutes INT,
    repeat_hours   INT,
    repeat_days    INT,
    repeat_weeks   INT,
    repeat_months  INT
);

INSERT INTO reminders (
    title, description, completed, created_date, start_date,
    end_date, time_hour, time_minute, repeat_seconds, repeat_minutes,
    repeat_hours, repeat_days, repeat_weeks, repeat_months
) VALUES
    ("Clean", "Clean stuff", false, "2024-07-20", "2024-07-21", "2024-07-22", 0, 0, 30, 0, 0, 0, 0, 0),
    ("Wipe", "Wipe stuff", false, "2024-07-20", "2024-07-21", "2024-07-22", 0, 0, 30, 0, 0, 0, 0, 0),
    ("Dust", "Dust stuff", false, "2024-07-20", "2024-07-21", "2024-07-22", 0, 0, 30, 0, 0, 0, 0, 0);

