package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mauricetjmurphy/mr-elephant/internal/api/services"
	"github.com/mauricetjmurphy/mr-elephant/internal/api/v1/models"
)

type HistoryHandler struct {
	historyService *services.HistoryService
}

func NewHistoryHandler(historyService *services.HistoryService) *HistoryHandler {
	return &HistoryHandler{historyService: historyService}
}

// GetHistory retrieves all history
//
//	@Summary		Get all history
//	@Description	Retrieve a list of all history entries
//	@Tags			history
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.HistoryDTO
//	@Failure		500	{object}	models.ErrorResponse
//	@Router			/history [get]
func (h *HistoryHandler) GetHistory(c *gin.Context) {
	history, err := h.historyService.GetAllHistory()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	var historyDTOs []models.HistoryDTO
	for _, h := range history {
		taskOptionsJSON, err := json.Marshal(h.TaskOptions)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: "Failed to marshal task options: " + err.Error()})
			return
		}
		historyDTOs = append(historyDTOs, models.HistoryDTO{
			ID:          h.ID,
			HistoryID:   h.HistoryID,
			TaskID:      h.TaskID,
			TaskType:    h.TaskType,
			TaskOptions: string(taskOptionsJSON),
			TaskResult:  h.TaskResult,
			Success:     h.Success,
		})
	}

	c.JSON(http.StatusOK, historyDTOs)
}
