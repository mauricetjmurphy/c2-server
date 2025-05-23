package dao

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TaskOption struct {
	OptionKey   string `json:"option_key"`
	OptionValue string `json:"option_value"`
}

type JSONTaskOptions []TaskOption

func (j *JSONTaskOptions) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), j)
}

func (j JSONTaskOptions) Value() (driver.Value, error) {
	return json.Marshal(j)
}

type History struct {
	ID          uint            `json:"id" gorm:"primaryKey;autoIncrement"`
	HistoryID   string          `json:"history_id"`
	TaskID      string          `json:"task_id"`
	TaskType    string          `json:"task_type"`
	TaskOptions JSONTaskOptions `json:"task_options" gorm:"type:json"`
	TaskResult  string          `json:"task_result"`
	Success     bool            `json:"success"`
}

// BeforeCreate Before creating a history, generate a UUID for HistoryID
func (history *History) BeforeCreate(tx *gorm.DB) (err error) {
	history.HistoryID = uuid.New().String()
	return
}

type IHistoryDAO interface {
	CreateHistory(history *History) error
	GetAllHistory() ([]History, error)
}

type HistoryDAO struct {
	db *gorm.DB
}

func NewHistoryDAO(db *gorm.DB) IHistoryDAO {
	return &HistoryDAO{db: db}
}

func (dao *HistoryDAO) CreateHistory(history *History) error {
	return dao.db.Create(history).Error
}

func (dao *HistoryDAO) GetAllHistory() ([]History, error) {
	var history []History
	err := dao.db.Find(&history).Error
	return history, err
}
