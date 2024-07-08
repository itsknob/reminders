package reminders

import (
	"encoding/json"
	"errors"
	"example/hello/db/dao"
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Item interface {

}
type Reminder struct {
	Id          int      `json:"id"`
    Type        string   `json:"type"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Schedule    Schedule `json:"schedule"`
}

type Schedule struct {
    Type   String `json:"type"`
	Time   Time   `json:"time"`
	Repeat Repeat `json:"repeat"`
}

type Time struct {
    Type   int `json:"type"`
	Hour   int `json:"int"`
	Minute int `json:"minute"`
}

type Repeat struct {
    Type    int `json:"type"`
	Seconds int `json:"seconds"`
	Minutes int `json:"minutes"`
	Hours   int `json:"hours"`
	Days    int `json:"days"`
	Weeks   int `json:"weeks"`
	Months  int `json:"months"`
}

var ExampleReminder = Reminder{
	Name:        "Clean",
	Description: "Clean the things",
	Schedule: Schedule{
		Repeat: Repeat{
			Seconds: 30,
		},
	},
}

var AllExampleReminders []Reminder = []Reminder{ExampleReminder}

func GetReminders(c *gin.Context) {
	ReminderDao := dao.
	reminders, err := ReminderDao.GetReminders()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}

	responseJSON, err := json.Marshal(reminders)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, responseJSON)
}
func GetReminder(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	for _, r := range AllExampleReminders {
		if r.Id == int(id) {
			body, err := json.Marshal(r)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, err)
				return
			}
			c.JSON(http.StatusOK, body)
		}
	}
	c.AbortWithError(http.StatusNotFound, errors.New(fmt.Sprintf("No Reminder found with Id: %d", id)))
	return
}

func sortReminders(i, j int) bool { return AllExampleReminders[i].Id < AllExampleReminders[j].Id }

func getNextId() int {
	sort.Slice(
		AllExampleReminders,
		sortReminders,
	)
	return AllExampleReminders[len(AllExampleReminders)-1].Id + 1
}

func CreateReminder(c *gin.Context) {
	newReminder := &Reminder{}
	if err := c.MustBindWith(newReminder, binding.JSON); err == nil {
		persistedReminder := Reminder{
			Id:          getNextId(),
			Name:        newReminder.Name,
			Description: newReminder.Description,
			Schedule:    newReminder.Schedule,
		}
		fmt.Printf("MOCK: DB - Reminders - Create - Id: %d\n", newReminder.Id)

		c.JSON(http.StatusCreated, persistedReminder)
	} else {
		c.AbortWithError(http.StatusBadRequest, errors.New(fmt.Sprintf("Failed to create Reminder from body - %+v", err)))
		return
	}
	c.AbortWithError(http.StatusInternalServerError, errors.New("Something went wrong"))
	return
}

func UpdateReminder(c *gin.Context) {
	c.JSON(http.StatusNotImplemented, errors.New("Not yet implemented"))
	return
}

func DeleteReminder(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 0)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, errors.New(fmt.Sprintf("Unable to parse Id %s - %+v", c.Param("id"), err)))
	}
	for _, reminder := range AllExampleReminders {
		if reminder.Id == int(id) {
			json, err := json.Marshal(reminder)
			if err != nil {
				c.AbortWithError(http.StatusInternalServerError, errors.New(fmt.Sprintf("Failed to marshal Reminder Id: %d - %+v", id, err)))
			}
			c.JSON(http.StatusOK, json)
			return
		}
	}
	c.JSON(http.StatusNotFound, errors.New(fmt.Sprintf("No Reminder found with Id: %d", id)))
	return

}
