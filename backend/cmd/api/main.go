// backend/cmd/api/main.go

// main.go is the entry point for the backend REST API service.
// It loads application configuration, registers Gin routes and Swagger documentation,
// and starts the local HTTP server.

package main

import (
	_ "ai-observability/backend/docs"
	"ai-observability/backend/internal/config"
	"ai-observability/backend/internal/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

// -------------------- Swagger Documentation --------------------
// Regenerate the Swagger/OpenAPI specification after changing route annotations:
// 		swag init -g ./cmd/api/main.go --parseInternal --parseDependency
//
// Swagger UI:
// 		http://localhost:8080/swagger/index.html

// -------------------- Running Backend Service Locally --------------------
// 	- cd to project root (backend)
// 	- Run the file
// 		go run ./cmd/api
//	- Test endpoints
//		http://localhost:8080

// -------------------- main --------------------
func main() {
	// -------------------- Setup --------------------
	// REST API Endpoints with Gin Web Framework

	// Load secrets and port number
	openApiKey, portNumber, serverSideError := loadSecrets()

	// To be removed later
	_ = openApiKey

	// Create a Gin router with the default logging and recovery middleware.
	router := gin.Default()

	// -------------------- REST API Endpoints --------------------
	// GET - http://localhost:8080/health
	router.GET("/health", func(c *gin.Context) {
		handler.Health(c, serverSideError)
	})

	// POST - http://localhost:8080/message
	router.POST("/message", handler.Message)

	// GET - http://localhost:8080/swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// -------------------- Server --------------------
	// Start the HTTP server
	fmt.Println("Starting server on port:", "127.0.0.1:"+portNumber)

	// Run on 127.0.0.1 for local development. Adding the "127.0.0.1"
	//  stops Windows Security from flagging.
	router.Run("127.0.0.1:" + portNumber)
}

// -------------------- Helper Functions --------------------

// loadSecrets loads the application configuration required by the API.
//
// It retrieves the OpenAI API key and server port from the configuration
// package. If configuration loading fails, the error is logged and the
// returned serverSideError flag is set to true so the health endpoint can
// report that the application did not initialize successfully.
//
// Returns:
//   - openApiKey: API key used to authenticate with the OpenAI API.
//   - portNumber: port on which the HTTP server should listen.
//   - serverSideError: true when configuration loading failed.
func loadSecrets() (string, string, bool) {
	// Load secrets
	openApiKey, portNumber, err := config.GetSecrets()

	// Variable initialized for if secrets loaded successfully
	serverSideError := err != nil

	// If there was an error loading the secrets
	if err != nil {
		log.Println("Failed to load configuration:", err)
		serverSideError = true
	}

	return openApiKey, portNumber, serverSideError
}
