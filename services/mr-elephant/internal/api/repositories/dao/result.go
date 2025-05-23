package dao

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Result struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	ResultID string `json:"result_id"`
	Contents string `json:"contents"`
	TaskID   string `json:"task_id"`
	Command  string `json:"command"`
	Success  bool   `json:"success"`
}

// BeforeCreate Before creating a result, generate a UUID for ResultID
func (result *Result) BeforeCreate(tx *gorm.DB) (err error) {
	result.ResultID = uuid.New().String()
	return
}

type IResultDAO interface {
	GetAllResults() ([]Result, error)
	SaveResult(result *Result) error
	BeginTransaction() *gorm.DB
}

type ResultDAO struct {
	db *gorm.DB
}

func NewResultDAO(db *gorm.DB) IResultDAO {
	return &ResultDAO{db: db}
}

func (dao *ResultDAO) GetAllResults() ([]Result, error) {
	var results []Result
	err := dao.db.Find(&results).Error
	return results, err
}

func (dao *ResultDAO) SaveResult(result *Result) error {
	return dao.db.Create(result).Error
}

func (dao *ResultDAO) BeginTransaction() *gorm.DB {
	return dao.db.Begin()
}
