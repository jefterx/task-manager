package models

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"Description"`
	Completed   bool   `json:"completed"`
}
