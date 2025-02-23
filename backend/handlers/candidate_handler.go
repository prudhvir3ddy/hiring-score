package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/prudhvir3ddy/hiring-score/models"
	"net/http"
	"strconv"
)

type CandidateServiceInterface interface {
	GetFilteredCandidates(query string) []models.Candidate
	PaginateCandidates(candidates []models.Candidate, page, pageCount int) (*models.PaginatedResponse, error)
}

type CandidateHandler struct {
	service CandidateServiceInterface
}

func NewCandidateHandler(service CandidateServiceInterface) *CandidateHandler {
	return &CandidateHandler{service: service}
}

func (h *CandidateHandler) GetCandidates(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page parameter"})
		return
	}

	pageCount, err := strconv.Atoi(c.DefaultQuery("page_count", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid page_count parameter"})
		return
	}

	query := c.DefaultQuery("query", "")
	candidates := h.service.GetFilteredCandidates(query)

	response, err := h.service.PaginateCandidates(candidates, page, pageCount)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Header("Access-Control-Allow-Origin", "*") // Temp: For CORS
	c.JSON(http.StatusOK, response)
}
