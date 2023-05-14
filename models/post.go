package models

type (
	Post struct {
		Text    string `json:"text"`
		Watches int    `json:"watches"`
	}
)
