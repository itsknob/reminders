package models

type Reminder struct {
	Id            int     `json:"id"`
	Title         string  `json:"title"`
	Completed     bool    `json:"completed"`
	Description   string  `json:"description,omitempty"`
	CreatedDate   string  `json:"createdDate"`
	StartDate     string  `json:"startDate"`
	EndDate       *string `json:"endDate"`
	TimeHour      int     `json:"timeHour"`
	TimeMinute    int     `json:"timeMinute"`
	RepeatSeconds int     `json:"repeatSeconds"`
	RepeatMinutes int     `json:"repeatMinutes"`
	RepeatHours   int     `json:"repeatHours"`
	RepeatDays    int     `json:"repeatDays"`
	RepeatWeeks   int     `json:"repeatWeeks"`
	RepeatMonths  int     `json:"repeatMonths"`
}

type ReminderPostBody struct {
	Title         string  `json:"title"`
	Completed     bool    `json:"completed"`
	Description   string  `json:"description,omitempty"`
	CreatedDate   string  `json:"createdDate"`
	StartDate     string  `json:"startDate"`
	EndDate       *string `json:"endDate"`
	TimeHour      int     `json:"timeHour"`
	TimeMinute    int     `json:"timeMinute"`
	RepeatSeconds int     `json:"repeatSeconds"`
	RepeatMinutes int     `json:"repeatMinutes"`
	RepeatHours   int     `json:"repeatHours"`
	RepeatDays    int     `json:"repeatDays"`
	RepeatWeeks   int     `json:"repeatWeeks"`
	RepeatMonths  int     `json:"repeatMonths"`
}
