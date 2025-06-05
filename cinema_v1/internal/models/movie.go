package models

type Movie struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Duration    int     `json:"duration"` // в минутах
	Rating      float64 `json:"rating"`
	Genre       string  `json:"genre"`
}
