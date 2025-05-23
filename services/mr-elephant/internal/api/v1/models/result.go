package models

type ResultDTO struct {
	ID       uint   `json:"id"`
	ResultID string `json:"result_id"`
	Contents string `json:"contents"`
	TaskID   string `json:"task_id"`
	Command  string `json:"command"`
	Success  bool   `json:"success"`
}
