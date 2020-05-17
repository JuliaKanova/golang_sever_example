package models

// TaskList ...
type TaskList struct {
	Tasks []Task `json:"tasks"`
}

//Task ...
type Task struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}
