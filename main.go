package main

import (
	"github.com/gin-gonic/gin"
)

// Start the gRPC proxy service
func main() {
	r := gin.Default()

	// Handle requests to get Loop Out terms
	r.POST("/v0/loopOutTerms", func(c *gin.Context) {
		loopOutTerms(c)
	})

	r.Run()
}
