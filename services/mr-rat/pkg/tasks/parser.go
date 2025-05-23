package tasks

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
)

// ParseTasksFrom parses a JSON response from the listener and returns a slice of Tasks
func ParseTasksFrom(taskData []byte, setter func(Configuration)) ([]Task, error) {
	var responses []ListenerResponse
	if err := json.Unmarshal(taskData, &responses); err != nil {
		return nil, err
	}

	var tasks []Task
	for _, response := range responses {
		id, err := uuid.Parse(response.TaskID)
		if err != nil {
			return nil, err
		}

		switch response.TaskType {
		case "ping":
			tasks = append(tasks, PingTask{ID: id})
		case "execute":
			tasks = append(tasks, ExecuteTask{ID: id, Command: response.Command})
		default:
			return nil, errors.New("unknown task type")
		}
	}
	return tasks, nil
}
