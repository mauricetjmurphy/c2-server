package models

type TaskOption struct {
	OptionKey   string `json:"option_key"`
	OptionValue string `json:"option_value"`
}

type HistoryDTO struct {
	ID          uint   `json:"id"`
	HistoryID   string `json:"history_id"`
	TaskID      string `json:"task_id"`
	TaskType    string `json:"task_type"`
	TaskOptions string `json:"task_options"`
	TaskResult  string `json:"task_result"`
	Success     bool   `json:"success"`
}
