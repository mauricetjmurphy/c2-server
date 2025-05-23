package tasks

import (
	"github.com/google/uuid"
)

// Task interface represents a generic task
type Task interface {
	Run() Result
}

// Result represents the result of a task execution
type Result struct {
	ID      uuid.UUID
	Output  string
	Success bool
}

// PingTask represents a ping task
type PingTask struct {
	ID uuid.UUID
}

// ExecuteTask represents an execute task
type ExecuteTask struct {
	ID      uuid.UUID
	Command string
}

// Configuration represents the configuration settings for the implant
type Configuration struct {
	MeanDwell float64
	IsRunning bool
}

// Payload represents the structure of the data sent to the listener
type Payload struct {
	Command string `json:"command"`
	Data    string `json:"data"`
}

// ListenerResponse represents the structure of the response from the listener
type ListenerResponse struct {
	ID       int    `json:"id"`
	TaskID   string `json:"task_id"`
	TaskType string `json:"task_type"`
	Command  string `json:"command,omitempty"`
}

// ResultPayload represents the structure of the result to be sent back to the listener
type ResultPayload struct {
	Success  bool   `json:"success"`
	TaskID   string `json:"task_id"`
	Contents string `json:"contents"`
}
