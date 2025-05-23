package dao

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	ID       uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	TaskID   string `json:"task_id"`
	TaskType string `json:"task_type"`
	Command  string `json:"command"`
}

// BeforeCreate Before creating a task, generate a UUID for TaskID
func (task *Task) BeforeCreate(tx *gorm.DB) (err error) {
	task.TaskID = uuid.New().String()
	return
}

type ITaskDAO interface {
	GetTask(taskID string) (*Task, error)
	CreateTask(task *Task) error
	ClearTask(taskID string) error
	GetAllTasks(tasks *[]Task) error
}

type TaskDAO struct {
	db *gorm.DB
}

func NewTaskDAO(db *gorm.DB) ITaskDAO {
	return &TaskDAO{db: db}
}

func (dao *TaskDAO) GetTask(taskID string) (*Task, error) {
	var task Task
	if err := dao.db.Where("task_id = ?", taskID).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (dao *TaskDAO) CreateTask(task *Task) error {
	return dao.db.Create(task).Error
}

func (dao *TaskDAO) ClearTask(taskID string) error {
	return dao.db.Where("task_id = ?", taskID).Delete(&Task{}).Error
}

func (dao *TaskDAO) GetAllTasks(tasks *[]Task) error {
	return dao.db.Find(tasks).Error
}
