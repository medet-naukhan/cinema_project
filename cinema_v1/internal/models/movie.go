package models

import "time"

type Movie struct {
	ID          string        `json:"id"`
	Title       string        `json:"title"`
	Description string        `json:"description"`
	Duration    time.Duration `json:"duration"` // реализовать через time.Time
	Rating      float64       `json:"rating"`
	Genre       string        `json:"genre"`
}
