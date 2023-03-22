package models

import "gorm.io/gorm"

type Event struct {
    gorm.Model
    Id int `json:"ID" gorm:"primary_key"`
    Category string `json:"category"`
    Title string `json:"title"`
    Description string `json:"description"`
    Location string `json:"location"`
    Date string `json:"date"`
    Time string `json:"time"`
    Organizer string `json:"organizer"`
}
