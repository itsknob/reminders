package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"

	"knob.dev/reminders/data"
	"knob.dev/reminders/models"
)

func main() {

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{AddSource: true}))

	fmt.Println("Hello, world!")

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	http.HandleFunc("/", IndexPage)
	http.HandleFunc("/reminders/{id}", GetReminder)
	http.HandleFunc("/reminders", GetReminders)
	http.HandleFunc("POST /reminders", NewReminder)
	http.HandleFunc("POST /reminders/{id}/complete", CompleteReminder)
	logger.Info("Registered Routes!")

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func IndexPage(w http.ResponseWriter, r *http.Request) {
	log.Printf("GET - Index.html\n")
	tmpl := template.Must(template.ParseFiles("index.html"))
	reminders, err := data.GetReminders()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(reminders)
		log.Printf("Failed to get Reminders \n%+v\n", err)
		return
	}
	slog.Info("IndexPage", "reminders", reminders)
	tmpl.Execute(w, reminders)
}

func CompleteReminder(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	updatedReminder, err := data.CompleteReminder(id)
	if err != nil {
		slog.Error("Failed to update reminder", "id", updatedReminder.Id)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
		return
	}

	slog.Info("UpdatedReminder ", "id", id, "completed", updatedReminder.Completed)

	tmpl := template.Must(template.ParseFiles("index.html"))
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
		log.Printf("Failed to create reminder: %+v\n", err)
		return
	}

	tmpl := template.Must(template.ParseFiles("index.html"))
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

	slog.Info("Found Reminder", "id", id)
	json.NewEncoder(w).Encode(reminder)
}

func GetReminders(w http.ResponseWriter, r *http.Request) {
	log.Printf("Getting all Reminders\n")
	reminders, err := data.GetReminders()

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(err)
	}

	slog.Info("Found Reminders", "count", len(reminders))
	json.NewEncoder(w).Encode(reminders)
}
