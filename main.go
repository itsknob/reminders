package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Reminder struct {
    Id string `json:"id,string"`
    Title string `json:"title"`
    Completed bool `json:"completed"`
    Description string `json:"description,omitempty"`
}

var RemindersDb []Reminder = []Reminder{
    {Id: "0", Title: "Clean", Completed: false, Description: "Clean stuff"},
    {Id: "1", Title: "Wipe", Completed: false, Description: "Wipe stuff"},
    {Id: "2", Title: "Dust", Completed: false, Description: "Dust stuff"},
}

func main() {
    fmt.Println("Hello, world!")

    http.HandleFunc("/", IndexPage)
    http.HandleFunc("/reminders/{id}", GetReminder)
    http.HandleFunc("/reminders", GetReminders)
    http.HandleFunc("POST /reminders", CreateReminder)
    http.HandleFunc("POST /reminders/{id}/complete", CompleteReminder)

    log.Fatal(http.ListenAndServe(":8080", nil))
}

func IndexPage(w http.ResponseWriter, r *http.Request) {
    log.Printf("GET - Index.html")
    tmpl := template.Must(template.ParseFiles("index.html"))
    tmpl.Execute(w, RemindersDb)
}

func CompleteReminder(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")

    if id == "" {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    reminder, err := findReminder(id)
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(err)
    }
    reminder.Completed = true
    for idx, reminder := range RemindersDb {
        if reminder.Id == id {
            reminder.Completed = true
            RemindersDb[idx] = reminder
        }
    }
    log.Printf("Completed Reminder %s", id)

    tmpl := template.Must(template.ParseFiles("index.html"))
    tmpl.ExecuteTemplate(w, "list-item", reminder)
}

func CreateReminder(w http.ResponseWriter, r  *http.Request) {

    id := len(RemindersDb)
    title := r.PostFormValue("title")
    description := r.PostFormValue("description")

    reminder := &Reminder{
        Id: strconv.Itoa(id),
        Title: title,
        Description: description,
        Completed: false,
    }

    log.Printf("Creating Reminder %d with Title %s", id, title)
    RemindersDb = append(RemindersDb, *reminder)

    tmpl := template.Must(template.ParseFiles("index.html"))
    tmpl.ExecuteTemplate(w, "list-item", reminder)
}

func findReminder(id string) (*Reminder, error) {
    for _, reminder := range RemindersDb {
        if reminder.Id == id {
            return &reminder, nil
        }
    }
    return nil, errors.New(fmt.Sprintf("Reminder not found with Id: %s", id))
}

func  GetReminder(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")

    if id == "" {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    log.Printf("Found reminder with Id %s", id)
    reminder, err := findReminder(id)
    if err != nil {
        w.WriteHeader(http.StatusNotFound)
        json.NewEncoder(w).Encode(err)
    }
    json.NewEncoder(w).Encode(reminder)

}
func GetReminders(w http.ResponseWriter, r *http.Request) {
    log.Printf("Getting all Reminders")
    json.NewEncoder(w).Encode(RemindersDb)
}
