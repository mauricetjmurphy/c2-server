package router

import (
	"github.com/mauricetjmurphy/mr-elephant/internal/api/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mauricetjmurphy/mr-elephant/docs"
	"github.com/mauricetjmurphy/mr-elephant/internal/api/middleware"
	"github.com/mauricetjmurphy/mr-elephant/internal/api/v1/handler"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter initializes the Gin router
func SetupRouter(taskHandler *handler.TaskHandler, resultHandler *handler.ResultHandler, historyHandler *handler.HistoryHandler) *gin.Engine {
	r := gin.Default()

	// Set the trusted proxies
	if len(config.AppConfig.TrustedProxies) > 0 {
		if err := r.SetTrustedProxies(config.AppConfig.TrustedProxies); err != nil {
			log.Fatalf("Failed to set trusted proxies: %v", err)
		}
	} else {
		if err := r.SetTrustedProxies(nil); err != nil {
			log.Fatalf("Failed to set trusted proxies: %v", err)
		}
	}

	docs.SwaggerInfo.BasePath = "/api/v1"

	// Apply the logger middleware
	r.Use(middleware.Logger())

	// Group API v1 routes
	v1 := r.Group("/api/v1")
	{
		v1.GET("/tasks", taskHandler.GetTasks)
		v1.POST("/task", taskHandler.CreateTask)

		v1.GET("/results", resultHandler.GetResults)
		v1.POST("/result", resultHandler.CreateResult)

		v1.GET("/history", historyHandler.GetHistory)

		// Health check
		r.GET("/ping", func(c *gin.Context) {
			c.String(http.StatusOK, "pong")
		})
	}

	// Add swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return r
}
