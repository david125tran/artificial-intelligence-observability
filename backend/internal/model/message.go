// backend/internal/model/message.go

// message.go defines the request and response data models used by the chat API.
// These structs establish the JSON contract exchanged between the frontend,
// HTTP handlers, and backend application logic.

package model

// The Message field is required and contains the user's input that will
// eventually be passed into the backend AI workflow.
type MessageRequest struct {
	Message string `json:"message" binding:"required"`
}

// The Response field contains the backend's processed response to the
// user's submitted message.
type MessageResponse struct {
	Response string `json:"response"`
}
