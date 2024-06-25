package main

import (
	// "fmt"
	"net/http"

    "github.com/labstack/echo/v4"
)

// type Reminder struct {
//     Id string
//     Name string
// }
//
// var reminders []string
//
// // e.GET('/reminders/:id', getReminder)
// func getReminder(c echo.Context) error {
//     id := c.Param("id")
//     return c.String(http.StatusOK, id)
// }
// func getReminders(c echo.Context) error {
//     return c.String(http.StatusOK, "[1, 2, 3]")
// }
// func postReminder(c echo.Context) error {
//     r := new(Reminder)
//     if err := c.Bind(r); err != nil {
//         return err
//     }
//     return c.JSON(http.StatusCreated, r)
// }

func main() {
    e := echo.New()

    // routing
    e.GET("/reminders", GetReminders)
    e.GET("/reminders/:id", GetReminder)
    e.POST("/reminder", PostReminder)

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, world!")
    })
    e.Logger.Fatal(e.Start(":3000"))


    // component := Index("John")
    // http.Handle("/", templ.Handler(component))
    // fmt.Println("Listening on :3000")
    // http.ListenAndServe(":3000", nil)
}


