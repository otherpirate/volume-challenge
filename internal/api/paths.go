package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/otherpirate/volume-challenge/internal/flights"
)

type PathRequest struct {
	Flights [][]string `json:"flights" binding:"required"`
	Start   string     `json:"start" binding:"required"`
	End     string     `json:"end" binding:"required"`
}

func Paths(c *gin.Context) {
	var req PathRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	g := flights.NewGraphService(req.Flights)
	paths := g.FindPaths(req.Start, req.End)
	c.JSON(http.StatusOK, gin.H{"paths": paths})
}
