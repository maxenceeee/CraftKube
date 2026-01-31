package main

import "github.com/gin-gonic/gin"

// Manager main entry point
func main() {
	router := gin.Default()

	// Start server
	router.Run("localhost:8080")
}
