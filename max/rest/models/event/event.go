package event

import "time"

type Event struct {
	ID          int       `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime" binding:"required"`
	UserID      int       `json:"userID"`
}

var events []Event

func (e *Event) Save() error {
	events = append(events, *e)
	return nil
}

func GetAllEvents() []Event {
	return events
}
