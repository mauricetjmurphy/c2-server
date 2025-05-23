package models

type TaskDTO struct {
	ID       uint   `json:"id"`
	TaskID   string `json:"task_id"`
	TaskType string `json:"task_type"`
	Command  string `json:"command"`
}
