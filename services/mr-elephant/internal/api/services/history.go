package services

import (
	"github.com/mauricetjmurphy/mr-elephant/internal/api/repositories/dao"
)

type HistoryService struct {
	historyDAO dao.IHistoryDAO
}

func NewHistoryService(historyDAO dao.IHistoryDAO) *HistoryService {
	return &HistoryService{historyDAO: historyDAO}
}

func (s *HistoryService) CreateHistory(history *dao.History) error {
	return s.historyDAO.CreateHistory(history)
}

func (s *HistoryService) GetAllHistory() ([]dao.History, error) {
	return s.historyDAO.GetAllHistory()
}
