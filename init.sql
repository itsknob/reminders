CREATE TABLE IF NOT EXISTS schedules (
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
    repeat_months  INT
);

INSERT INTO schedules (
    id,
    created_date,
    start_date,
    end_date,
    time_hour,
    time_minute,
    repeat_seconds,
    repeat_minutes,
    repeat_hours,
    repeat_days,
    repeat_weeks,
    repeat_months
) VALUES ("0", "2024-07-20", "2024-07-21", "2024-07-22", 0, 0, 30, 0, 0, 0, 0, 0);

---

CREATE TABLE IF NOT EXISTS reminders (
    id INT PRIMARY_KEY,
    title VARCHAR(64),
    description VARCHAR(255),
    completed,
    schedule_id INT,
    FOREIGN KEY (schedule_id) REFERENCES schedules(id)
);

INSERT INTO reminders (
    id, title, description, completed, schedule_id
) VALUES
    ("0", "Clean", "Clean stuff", false, 0),
    ("1", "Wipe", "Wipe stuff", false, 0),
    ("2", "Dust", "Dust stuff", false, 0);
