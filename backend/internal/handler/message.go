// backend/internal/handler/message.go

// message.go contains the HTTP handler for the chat /message endpoint.
// It validates incoming JSON against the request model and returns
// a structured JSON response to the client.

package handler

import (
	"ai-observability/backend/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Message handles POST requests to the message endpoint.
//
// @Summary Send a message
// @Description Accepts a JSON message payload and returns a response.
// @Tags Message
// @Accept json
// @Produce json
// @Param request body model.MessageRequest true "Message payload"
// @Success 200 {object} model.MessageResponse
// @Failure 400 {object} map[string]string
// @Router /message [post]
func Message(c *gin.Context) {
	var request model.MessageRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request payload.",
		})
		return
	}

	response := model.MessageResponse{
		Response: "Message received: " + request.Message,
	}

	c.JSON(http.StatusOK, response)
}
