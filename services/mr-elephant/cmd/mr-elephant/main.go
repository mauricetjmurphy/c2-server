package main

import (
	"log"

	"github.com/mauricetjmurphy/mr-elephant/internal/api/config"
	"github.com/mauricetjmurphy/mr-elephant/internal/api/database"
	"github.com/mauricetjmurphy/mr-elephant/internal/api/repositories/dao"
	"github.com/mauricetjmurphy/mr-elephant/internal/api/services"
	"github.com/mauricetjmurphy/mr-elephant/internal/api/v1/handler"
	"github.com/mauricetjmurphy/mr-elephant/internal/api/v1/router"
)

//	@title			Swagger Mr Elephant API
//	@version		1.0
//	@description	This is a server R2 server listener.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

// @host		http://127.0.0.1:4000
// @BasePath	/v1
func main() {
	// Load configuration
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Error loading configuration: %v", err)
	}

	// Initialize the database connection
	database.InitDB()

	// Initialize DAOs
	taskDAO := dao.NewTaskDAO(database.DB)
	resultDAO := dao.NewResultDAO(database.DB)
	historyDAO := dao.NewHistoryDAO(database.DB)

	// Initialize Services
	taskService := services.NewTaskService(taskDAO)
	resultService := services.NewResultService(resultDAO, taskDAO, historyDAO, taskService)
	historyService := services.NewHistoryService(historyDAO)

	// Setup handlers
	taskHandler := handler.NewTaskHandler(taskService)
	resultHandler := handler.NewResultHandler(resultService)
	historyHandler := handler.NewHistoryHandler(historyService)

	// Setup routes
	r := router.SetupRouter(taskHandler, resultHandler, historyHandler)
	err := r.Run(config.AppConfig.ServerAddress)
	if err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}
