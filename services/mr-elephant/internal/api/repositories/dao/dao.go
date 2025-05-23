package dao

import (
	"gorm.io/gorm"
)

// DAO aggregates all DAO interfaces
type DAO struct {
	IHistoryDAO
	IResultDAO
	ITaskDAO
}

// New returns a new DAO instance
func New(db *gorm.DB) *DAO {
	historyDAO := NewHistoryDAO(db)
	resultDAO := NewResultDAO(db)
	taskDAO := NewTaskDAO(db)
	return NewDAOs(historyDAO, resultDAO, taskDAO)
}

// NewDAOs returns a DAO struct with provided interfaces
func NewDAOs(
	historyDAO IHistoryDAO,
	resultDAO IResultDAO,
	taskDAO ITaskDAO,
) *DAO {
	return &DAO{
		IHistoryDAO: historyDAO,
		IResultDAO:  resultDAO,
		ITaskDAO:    taskDAO,
	}
}
