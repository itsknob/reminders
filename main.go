package main

import (
	"example/hello/reminders"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(c *gin.Context) {
    c.Next()

    for _, err := range c.Errors {
        // log here
        fmt.Println(err.JSON)
    }

    c.JSON(http.StatusInternalServerError, "")
}

func main() {
    router := gin.Default()

    // Middleware
    router.Use(gin.Logger())
    router.Use(gin.Recovery())

    // Routes
    router.GET("/foo", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })

    reminderRoutes := router.Group("/reminders")

    {
        reminderRoutes.GET("", reminders.GetReminders)
        reminderRoutes.GET("/:id", reminders.GetReminder)
        reminderRoutes.POST("", reminders.CreateReminder)
        reminderRoutes.PUT("/:id", reminders.UpdateReminder)
        reminderRoutes.DELETE("/:id", reminders.DeleteReminder)
    }

    // Start
    router.Run()
}
