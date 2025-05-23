package services

import (
	"github.com/mauricetjmurphy/mr-elephant/internal/api/repositories/dao"
)

type ResultService struct {
	resultDAO   dao.IResultDAO
	taskDAO     dao.ITaskDAO
	historyDAO  dao.IHistoryDAO
	taskService *TaskService
}

func NewResultService(resultDAO dao.IResultDAO, taskDAO dao.ITaskDAO, historyDAO dao.IHistoryDAO, taskService *TaskService) *ResultService {
	return &ResultService{
		resultDAO:   resultDAO,
		taskDAO:     taskDAO,
		historyDAO:  historyDAO,
		taskService: taskService,
	}
}

func (s *ResultService) GetAllResults() ([]dao.Result, error) {
	return s.resultDAO.GetAllResults()
}

func (s *ResultService) CreateNewResult(result *dao.Result, requestData map[string]interface{}) (*dao.Result, error) {
	tx := s.resultDAO.BeginTransaction()

	task, err := s.taskDAO.GetTask(result.TaskID)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	taskOptions, err := s.taskService.CreateTaskOptions(requestData)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	history := &dao.History{
		TaskID:      task.TaskID,
		TaskType:    task.TaskType,
		TaskOptions: taskOptions,
		TaskResult:  result.Contents,
		Success:     result.Success,
	}

	if err := s.historyDAO.CreateHistory(history); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := s.resultDAO.SaveResult(result); err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := s.taskDAO.ClearTask(result.TaskID); err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return result, nil
}
