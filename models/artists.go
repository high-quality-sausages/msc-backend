package models

type Artist struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Nation  string `json:"nation"`
	Picture string `json:"picture"`
}
