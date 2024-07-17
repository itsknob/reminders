package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Reminder struct {
    Id string `json:"id"`
    Title string `json:"title"`
}

var RemindersDb []Reminder = []Reminder{
       {Id: "0", Title: "Clean"},
       {Id: "1", Title: "Wipe"},
       {Id: "2", Title: "Dust"},
    }

func main() {
    fmt.Println("Hello, world!")

    http.HandleFunc("/", IndexPage)
    http.HandleFunc("/reminders/{id}", GetReminder)
    http.HandleFunc("/reminders", GetReminders)
    http.HandleFunc("POST /reminders", CreateReminder)

    log.Fatal(http.ListenAndServe(":8080", nil))
}

func IndexPage(w http.ResponseWriter, r *http.Request) {
    log.Printf("GET - Index.html")
    tmpl := template.Must(template.ParseFiles("index.html"))
    tmpl.Execute(w, RemindersDb)
}

func CreateReminder(w http.ResponseWriter, r  *http.Request) {

    id := len(RemindersDb)
    title := r.PostFormValue("title")

    reminder := &Reminder{
        Id: strconv.Itoa(id),
        Title: title,
    }

    log.Printf("Creating Reminder %d with Title %s", id, title)
    RemindersDb = append(RemindersDb, *reminder)

    tmpl := template.Must(template.ParseFiles("index.html"))
    tmpl.ExecuteTemplate(w, "list-item", reminder)
}

func  GetReminder(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")

    if id == "" {
        w.WriteHeader(http.StatusNotFound)
        return
    }

    for _, reminder := range RemindersDb {
        if reminder.Id == id {
            log.Printf("Found reminder with Id %s", id)
            json.NewEncoder(w).Encode(reminder)
        }
    }

}
func GetReminders(w http.ResponseWriter, r *http.Request) {
    log.Printf("Getting all Reminders")
    json.NewEncoder(w).Encode(RemindersDb)
}
