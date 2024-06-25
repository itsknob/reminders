package main

import (
    "testing"
)

var Reminders = [3]Reminder{
    {Id: "1", Name: "First"}, 
    {Id: "2", Name: "Second"}, 
    {Id: "3", Name: "Third"},
}
func TestGetReminders(t *testing.T) {
    want := Reminders
}
