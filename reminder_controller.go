package main

import (
	"net/http"
    // "github.com/a-h/templ"
    "github.com/labstack/echo/v4"
)

type Reminder struct {
    Id string
    Name string
}

var reminders []Reminder = []Reminder{
    {
        Id: "1",
        Name: "First",
    }, 
    {
        Id: "2",
        Name: "Second",
    }, 
    {
        Id: "3",
        Name: "Third",
    },
}

// e.GET('/reminders/:id', getReminder)
func GetReminder(c echo.Context) error {
    id := c.Param("id")
    var found = Reminder{}
    for _, rem := range reminders {
        if rem.Id == id {
            found = rem
            return c.JSON(http.StatusOK, rem)
        }
    }
    if found == (Reminder{}) {
        return c.JSON(http.StatusNotFound, "{ \"error\": \"Reminder with Id " + id + " not found\"}")
    }
    return c.String(http.StatusTeapot, "You managed to send a GET Request to a Teapot.")
}
func GetReminders(c echo.Context) error {
    return c.JSON(http.StatusOK, reminders)
}
func PostReminder(c echo.Context) error {
    r := new(Reminder) 
    if err := c.Bind(r); err != nil {
        return err
    }
    return c.JSON(http.StatusCreated, r)
}
