package dto

import "time"

type EventResponse struct {
	EventID uint `json:"eventId"`
	Category    string    `json:"category"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	Date        time.Time `json:"date"`
	Time        string    `json:"time"`
	Organizer   string    `json:"organizer"`
	UserID      int       `json:"creator"`
}

