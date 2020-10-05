package main

import (
	"github.com/gin-gonic/gin"
)

// Start the gRPC proxy service
func main() {
	app := gin.Default()

	// Handle requests to get a Loop In quote
	app.POST("/v0/loopInQuote", func(context *gin.Context) {
		loopInQuote(context)
	})

	// Handle requests to get Loop Out terms
	app.POST("/v0/loopOutPushPreimage", func(context *gin.Context) {
		loopOutPushPreimage(context)
	})

	// Handle requests to get a Loop Out quote
	app.POST("/v0/loopOutQuote", func(context *gin.Context) {
		loopOutQuote(context)
	})

	// Handle requests to get Loop Out terms
	app.POST("/v0/loopOutTerms", func(context *gin.Context) {
		loopOutTerms(context)
	})

	// Handle request to start a new Loop In
	app.POST("/v0/newLoopInSwap", func(context *gin.Context) {
		newLoopInSwap(context)
	})

	// Handle request to start a new Loop Out
	app.POST("/v0/newLoopOutSwap", func(context *gin.Context) {
		newLoopOutSwap(context)
	})

	app.Run()
}
