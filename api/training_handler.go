package api

import (
	"SCMZU_Party_Building/entity"
	"SCMZU_Party_Building/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TrainingHandler struct {
	trainingService *service.TrainingService
}

func NewTrainingHandler(trainingService *service.TrainingService) *TrainingHandler {
	return &TrainingHandler{trainingService: trainingService}
}

func (h *TrainingHandler) ListTrainings(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	trainings, total, err := h.trainingService.ListTrainings(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  trainings,
		"total": total,
		"page":  page,
	})
}

func (h *TrainingHandler) GetTraining(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid training ID"})
		return
	}

	training, err := h.trainingService.GetTraining(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Training not found"})
		return
	}

	c.JSON(http.StatusOK, training)
}

func (h *TrainingHandler) CreateTraining(c *gin.Context) {
	var training entity.Training
	if err := c.ShouldBindJSON(&training); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.trainingService.CreateTraining(&training)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, training)
}

func (h *TrainingHandler) UpdateTraining(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid training ID"})
		return
	}

	var training entity.Training
	if err := c.ShouldBindJSON(&training); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	training.ID = uint(id)
	err = h.trainingService.UpdateTraining(&training)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, training)
}

func (h *TrainingHandler) DeleteTraining(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid training ID"})
		return
	}

	err = h.trainingService.DeleteTraining(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Training deleted successfully"})
}
