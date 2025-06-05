package models

type Movie struct {
	ID          string  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Duration    int     `json:"duration"` // реализовать через time.Time
	Rating      float64 `json:"rating"`
	Genre       string  `json:"genre"`
}
