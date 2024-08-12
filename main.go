package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"knob.dev/reminders/data"
	"knob.dev/reminders/models"
)

func main() {

	fmt.Println("Hello, world!")

	http.HandleFunc("/", IndexPage)
	http.HandleFunc("/reminders/{id}", GetReminder)
	http.HandleFunc("/reminders", GetReminders)
	http.HandleFunc("POST /reminders", NewReminder)
	http.HandleFunc("POST /reminders/{id}/complete", CompleteReminder)

	log.Printf("Registered Routes!")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func IndexPage(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET - Index.html\n")
	tmpl := template.Must(template.ParseFiles("index.html"))
	reminders, err := data.GetReminders()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(reminders)
		panic(fmt.Sprintf("Failed to get Reminders %+v", err))
	}
	log.Printf("IndexPage - reminders: \n%+v\n", reminders)
	tmpl.Execute(w, reminders)
}

// func NewSchedule(w http.ResponseWriter, r *http.Request) {
// 	timeHour, _ := strconv.Atoi(r.PostFormValue("timeHour"))
// 	timeMinute, _ := strconv.Atoi(r.PostFormValue("timeMinute"))
// 	repeatSeconds, _ := strconv.Atoi(r.PostFormValue("repeatSeconds"))
// 	repeatMinutes, _ := strconv.Atoi(r.PostFormValue("repeatMinutes"))
// 	repeatHours, _ := strconv.Atoi(r.PostFormValue("repeatHours"))
// 	repeatDays, _ := strconv.Atoi(r.PostFormValue("repeatDays"))
// 	repeatWeeks, _ := strconv.Atoi(r.PostFormValue("repeatWeeks"))
// 	repeatMonths, _ := strconv.Atoi(r.PostFormValue("repeatMonths"))
//
// 	newSchedule, err := data.CreateSchedule(&models.SchedulePostBody{
// 		TimeHour:      timeHour,
// 		TimeMinute:    timeMinute,
// 		RepeatSeconds: repeatSeconds,
// 		RepeatMinutes: repeatMinutes,
// 		RepeatHours:   repeatHours,
// 		RepeatDays:    repeatDays,
// 		RepeatWeeks:   repeatWeeks,
// 		RepeatMonths:  repeatMonths,
// 	})
//
// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		json.NewEncoder(w).Encode(err)
// 	}
//
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(newSchedule)
//
// }

func CompleteReminder(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	updatedReminder, err := data.CompleteReminder(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
	}

	log.Printf("Completed Reminder %s\n", id)

	tmpl := template.Must(template.ParseFiles("index.html"))
	log.Printf("*updatedReminder - %+v\n", *updatedReminder)
	tmpl.ExecuteTemplate(w, "list-item", updatedReminder)
}

func getPostFormValueAsInt(r *http.Request, key string) int {
	value := r.PostFormValue(key)
	intValue, err := strconv.Atoi(value)
	if err != nil {
		panic("Failed to convert String to Integer")
	}
	return intValue
}

func NewReminder(w http.ResponseWriter, r *http.Request) {

	fmt.Println("New Reminder")

	title := r.PostFormValue("title")
	description := r.PostFormValue("description")

	timeHour := getPostFormValueAsInt(r, "timeHour")
	timeMinute := getPostFormValueAsInt(r, "timeMinute")
	repeatSeconds := getPostFormValueAsInt(r, "repeatSeconds")
	repeatMinutes := getPostFormValueAsInt(r, "repeatMinutes")
	repeatHours := getPostFormValueAsInt(r, "repeatHours")
	repeatDays := getPostFormValueAsInt(r, "repeatDays")
	repeatWeeks := getPostFormValueAsInt(r, "repeatWeeks")
	repeatMonths := getPostFormValueAsInt(r, "repeatMonths")

	// scheduleInput := &models.SchedulePostBody{
	// 	TimeHour:      timeHour,
	// 	TimeMinute:    timeMinute,
	// 	RepeatSeconds: repeatSeconds,
	// 	RepeatMinutes: repeatMinutes,
	// 	RepeatHours:   repeatHours,
	// 	RepeatDays:    repeatDays,
	// 	RepeatWeeks:   repeatWeeks,
	// 	RepeatMonths:  repeatMonths,
	// }
	//
	// fmt.Printf("main - NewReminder - scheduleInput: "bu\n%+v\n", scheduleInput)
	//
	// schedule, err := data.CreateSchedule(scheduleInput)

	// if err != nil {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	json.NewEncoder(w).Encode(err)
	// 	return
	// }

	// fmt.Printf("main - NewReminder - Schedule: \n%+v\n", schedule)

	newReminder, err := data.CreateReminder(&models.ReminderPostBody{
		Title:         title,
		Completed:     false,
		Description:   description,
		CreatedDate:   "",
		StartDate:     "",
		EndDate:       new(string),
		TimeHour:      timeHour,
		TimeMinute:    timeMinute,
		RepeatSeconds: repeatSeconds,
		RepeatMinutes: repeatMinutes,
		RepeatHours:   repeatHours,
		RepeatDays:    repeatDays,
		RepeatWeeks:   repeatWeeks,
		RepeatMonths:  repeatMonths,
	})

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		panic(err)
	}

	tmpl := template.Must(template.ParseFiles("index.html"))

	fmt.Printf("main - NewReminder - newReminder:\n%+v\n", newReminder)

	tmpl.ExecuteTemplate(w, "list-item", newReminder)
}

func GetReminder(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	reminder, err := data.FindReminder(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
	}

	log.Printf("Found reminder with Id %s\n", id)

	json.NewEncoder(w).Encode(reminder)

}

func GetReminders(w http.ResponseWriter, r *http.Request) {
	log.Printf("Getting all Reminders\n")
	reminders, err := data.GetReminders()

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
	}

	json.NewEncoder(w).Encode(reminders)

}
