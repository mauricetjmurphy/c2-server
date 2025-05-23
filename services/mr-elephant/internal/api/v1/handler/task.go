package handler

import (
	"github.com/mauricetjmurphy/mr-elephant/internal/api/repositories/dao"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mauricetjmurphy/mr-elephant/internal/api/services"
	"github.com/mauricetjmurphy/mr-elephant/internal/api/v1/models"
)

type TaskHandler struct {
	taskService *services.TaskService
}

func NewTaskHandler(taskService *services.TaskService) *TaskHandler {
	return &TaskHandler{taskService: taskService}
}

// GetTasks retrieves all tasks
//
//	@Summary		Get all tasks
//	@Description	Retrieve a list of all tasks
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.TaskDTO
//	@Failure		500	{object}	models.ErrorResponse
//	@Router			/tasks [get]
func (h *TaskHandler) GetTasks(c *gin.Context) {
	tasks, err := h.taskService.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to fetch tasks: " + err.Error()})
		return
	}

	var taskDTOs []models.TaskDTO
	for _, task := range tasks {
		taskDTOs = append(taskDTOs, models.TaskDTO{
			ID:       task.ID,
			TaskID:   task.TaskID,
			TaskType: task.TaskType,
			Command:  task.Command,
		})
	}

	c.JSON(http.StatusOK, taskDTOs)
}

// CreateTask creates a new task
//
//	@Summary		Create a new task
//	@Description	Create a new task with the given details
//	@Tags			tasks
//	@Accept			json
//	@Produce		json
//	@Param			task	body		models.TaskDTO	true	"Task to create"
//	@Success		201		{object}	models.TaskDTO
//	@Failure		400		{object}	models.ErrorResponse
//	@Failure		500		{object}	models.ErrorResponse
//	@Router			/task [post]
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var taskDTO models.TaskDTO
	if err := c.ShouldBindJSON(&taskDTO); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid JSON provided: " + err.Error()})
		return
	}

	if taskDTO.TaskType == "" {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "TaskType is required"})
		return
	}

	task := &dao.Task{
		TaskType: taskDTO.TaskType,
		Command:  taskDTO.Command,
	}

	if err := h.taskService.CreateTask(task); err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to create task: " + err.Error()})
		return
	}

	taskDTO.ID = task.ID
	taskDTO.TaskID = task.TaskID

	c.JSON(http.StatusCreated, taskDTO)
}
