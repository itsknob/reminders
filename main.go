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

    log.Fatal(http.ListenAndServe(":8080", nil))
}

func IndexPage(w http.ResponseWriter, r *http.Request) {
    log.Printf("GET - Index.html")
    tmpl := template.Must(template.ParseFiles("index.html"))
    reminders, err := data.GetReminders()
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(reminders)
    }
    tmpl.Execute(w, reminders)
}

func NewSchedule(w http.ResponseWriter, r  *http.Request) {
    timeHour, _ := strconv.Atoi(r.PostFormValue("timeHour"))
    timeMinute, _ := strconv.Atoi(r.PostFormValue("timeMinute"))
    repeatSeconds, _ := strconv.Atoi(r.PostFormValue("repeatSeconds"))
    repeatMinutes, _ := strconv.Atoi(r.PostFormValue("repeatMinutes"))
    repeatHours, _ := strconv.Atoi(r.PostFormValue("repeatHours"))
    repeatDays, _ := strconv.Atoi(r.PostFormValue("repeatDays"))
    repeatWeeks, _ := strconv.Atoi(r.PostFormValue("repeatWeeks"))
    repeatMonths, _ := strconv.Atoi(r.PostFormValue("repeatMonths"))

    newSchedule, err := data.CreateSchedule(&models.SchedulePostBody{
        TimeHour: timeHour,
        TimeMinute: timeMinute,
        RepeatSeconds: repeatSeconds,
        RepeatMinutes: repeatMinutes,
        RepeatHours: repeatHours,
        RepeatDays: repeatDays,
        RepeatWeeks: repeatWeeks,
        RepeatMonths: repeatMonths,
    })
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(err)
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(newSchedule)

}

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

    log.Printf("Completed Reminder %s", id)

    tmpl := template.Must(template.ParseFiles("index.html"))
    log.Printf("%+v", *updatedReminder)
    tmpl.ExecuteTemplate(w, "list-item", updatedReminder)
}


func NewReminder(w http.ResponseWriter, r  *http.Request) {

    title := r.PostFormValue("title")
    description := r.PostFormValue("description")

    newReminder, err := data.CreateReminder(&models.ReminderPostBody{
        Title: title,
        Completed: false,
        Description: description,
    }, nil)

    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(err)
    }


    tmpl := template.Must(template.ParseFiles("index.html"))
    tmpl.ExecuteTemplate(w, "list-item", newReminder)
}

func  GetReminder(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")

    if id == "" {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    log.Printf("Found reminder with Id %s", id)
    reminder, err := data.FindReminder(id)
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(err)
    }
    json.NewEncoder(w).Encode(reminder)

}

func GetReminders(w http.ResponseWriter, r *http.Request) {
    log.Printf("Getting all Reminders")
    reminders, err := data.GetReminders()

    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(err)
    }

    json.NewEncoder(w).Encode(reminders)

}
