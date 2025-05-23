package database

import (
	"log"

	"github.com/mauricetjmurphy/mr-elephant/internal/api/config"
	"github.com/mauricetjmurphy/mr-elephant/internal/api/repositories/dao"
	"github.com/mauricetjmurphy/mr-elephant/internal/api/services"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB             *gorm.DB
	TaskService    *services.TaskService
	ResultService  *services.ResultService
	HistoryService *services.HistoryService
)

func InitDB() {
	if config.AppConfig == nil {
		log.Fatal("AppConfig is not initialized.")
	}

	dsn := config.AppConfig.DSN()
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Migrate the schema
	err = DB.AutoMigrate(&dao.Task{}, &dao.Result{}, &dao.History{})
	if err != nil {
		log.Fatalf("Error migrating the database: %v", err)
	}

	// Initialize DAOs
	taskDAO := dao.NewTaskDAO(DB)
	resultDAO := dao.NewResultDAO(DB)
	historyDAO := dao.NewHistoryDAO(DB)

	// Initialize Services
	TaskService = services.NewTaskService(taskDAO)
	ResultService = services.NewResultService(resultDAO, taskDAO, historyDAO, TaskService)
	HistoryService = services.NewHistoryService(historyDAO)

	log.Println("Successfully connected to the database")
}
