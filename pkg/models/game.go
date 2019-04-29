package models

//Game ..
type Game struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	URLImage    string `json:"url_image"`
	StartAt     string `json:"start_at"`
	EndAt       string `json:"end_at"`
	Private     bool   `json:"private"`
}
