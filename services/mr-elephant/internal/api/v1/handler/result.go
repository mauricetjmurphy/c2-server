package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mauricetjmurphy/mr-elephant/internal/api/repositories/dao"
	"github.com/mauricetjmurphy/mr-elephant/internal/api/services"
	utils "github.com/mauricetjmurphy/mr-elephant/internal/api/utils"
	"github.com/mauricetjmurphy/mr-elephant/internal/api/v1/models"
)

type ResultHandler struct {
	resultService *services.ResultService
}

func NewResultHandler(resultService *services.ResultService) *ResultHandler {
	return &ResultHandler{resultService: resultService}
}

// GetResults retrieves all results
//
//	@Summary		Get all results
//	@Description	Retrieve a list of all results
//	@Tags			results
//	@Accept			json
//	@Produce		json
//	@Success		200	{array}		models.ResultDTO
//	@Failure		500	{object}	models.ErrorResponse
//	@Router			/results [get]
func (h *ResultHandler) GetResults(c *gin.Context) {
	results, err := h.resultService.GetAllResults()
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	var resultDTOs []models.ResultDTO
	for _, result := range results {
		resultDTOs = append(resultDTOs, models.ResultDTO{
			ID:       result.ID,
			ResultID: result.ResultID,
			Contents: result.Contents,
			TaskID:   result.TaskID,
			Command:  result.Command,
			Success:  result.Success,
		})
	}

	c.JSON(http.StatusOK, resultDTOs)
}

// CreateResult creates a new result
//
//	@Summary		Create a new result
//	@Description	Create a new result with the given details
//	@Tags			results
//	@Accept			json
//	@Produce		json
//	@Param			result	body		models.ResultDTO	true	"Result to create"
//	@Success		200		{object}	models.ResultDTO
//	@Failure		400		{object}	models.ErrorResponse
//	@Failure		500		{object}	models.ErrorResponse
//	@Router			/result [post]
func (h *ResultHandler) CreateResult(c *gin.Context) {
	var resultDTO models.ResultDTO

	body, err := utils.ReadAndLogBody(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid request body"})
		return
	}

	if err := utils.BindJSON(c, body, &resultDTO); err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "Invalid JSON provided"})
		return
	}

	result := &dao.Result{
		Contents: resultDTO.Contents,
		TaskID:   resultDTO.TaskID,
		Command:  resultDTO.Command,
		Success:  resultDTO.Success,
	}

	createdResult, err := h.resultService.CreateNewResult(result, map[string]interface{}{
		"contents": resultDTO.Contents,
		"task_id":  resultDTO.TaskID,
		"command":  resultDTO.Command,
		"success":  resultDTO.Success,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResponse{Error: err.Error()})
		return
	}

	resultDTO.ID = createdResult.ID
	resultDTO.ResultID = createdResult.ResultID

	c.JSON(http.StatusOK, resultDTO)
}
