package models

import "time"

//GameRequest ...
type GameRequest struct {
	Name        string    `json:"name"`
	Description string    `json:"description"`
	URLImage    string    `json:"url_image"`
	StartAt     time.Time `json:"start_at"`
	EndAt       time.Time `json:"end_at"`
}
