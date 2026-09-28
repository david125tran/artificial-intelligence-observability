// backend/internal/handler/health.go

// health.go contains the HTTP handler for the application's health endpoint.
// It reports whether the backend initialized successfully and returns an
// appropriate HTTP status for healthy and unhealthy states.

package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Health handles GET requests to the health endpoint.
//
// @Summary Check API health
// @Description Returns the current health state of the backend.
// @Tags Health
// @Produce json
// @Success 200 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /health [get]
func Health(c *gin.Context, serverSideError bool) {
	if serverSideError {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": "unhealthy",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}
