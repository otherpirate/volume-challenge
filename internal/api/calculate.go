package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/otherpirate/volume-challenge/internal/flights"
)

type CalculateRequest struct {
	Flights [][]string `json:"flights" binding:"required"`
}

func Calculate(c *gin.Context) {
	var req CalculateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	path, err := flights.FindPath(req.Flights)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"path": path})
}
