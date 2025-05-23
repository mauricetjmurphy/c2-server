package services

import (
	"fmt"
	"github.com/mauricetjmurphy/mr-elephant/internal/api/repositories/dao"
)

type TaskService struct {
	taskDAO dao.ITaskDAO
}

func NewTaskService(taskDAO dao.ITaskDAO) *TaskService {
	return &TaskService{taskDAO: taskDAO}
}

func (s *TaskService) GetTask(taskID string) (*dao.Task, error) {
	return s.taskDAO.GetTask(taskID)
}

func (s *TaskService) CreateTask(task *dao.Task) error {
	return s.taskDAO.CreateTask(task)
}

func (s *TaskService) ClearTask(taskID string) error {
	return s.taskDAO.ClearTask(taskID)
}

func (s *TaskService) GetTasks() ([]dao.Task, error) {
	var tasks []dao.Task
	err := s.taskDAO.GetAllTasks(&tasks)
	return tasks, err
}

func (s *TaskService) CreateTaskOptions(requestData map[string]interface{}) ([]dao.TaskOption, error) {
	var taskOptions []dao.TaskOption
	for key, value := range requestData {
		if key != "task_type" && key != "task_id" {
			taskOptions = append(taskOptions, dao.TaskOption{
				OptionKey:   key,
				OptionValue: fmt.Sprintf("%v", value),
			})
		}
	}
	return taskOptions, nil
}
