package models

type Reminder struct {
    Id string `json:"id,string"`
    Title string `json:"title"`
    Completed bool `json:"completed"`
    Description string `json:"description,omitempty"`
}

type ReminderPostBody struct {
    Title string `json:"title"`
    Completed bool `json:"completed"`
    Description string `json:"description,omitempty"`
}

type Schedule struct {
    Id     string `json:"id,string"`
	TimeHour   int `json:"timeHour"`
	TimeMinute int `json:"timeMinute"`
	RepeatSeconds int `json:"repeatSeconds"`
	RepeatMinutes int `json:"repeatMinutes"`
	RepeatHours   int `json:"repeatHours"`
	RepeatDays    int `json:"repeatDays"`
	RepeatWeeks   int `json:"repeatWeeks"`
	RepeatMonths  int `json:"repeatMonths"`
}

type SchedulePostBody struct {
	TimeHour      int    `json:"timeHour"`
	TimeMinute    int    `json:"timeMinute"`
	RepeatSeconds int    `json:"repeatSeconds"`
	RepeatMinutes int    `json:"repeatMinutes"`
	RepeatHours   int    `json:"repeatHours"`
	RepeatDays    int    `json:"repeatDays"`
	RepeatWeeks   int    `json:"repeatWeeks"`
	RepeatMonths  int    `json:"repeatMonths"`
}

