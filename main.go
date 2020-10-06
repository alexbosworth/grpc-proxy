package main

import (
	"github.com/gin-gonic/gin"
	"os"
)

var (
	defaultServerPort     = "8080"
	listenPortEnvVar      = "PORT"
	tlsCertFilePathEnvVar = "TLS_CERT_FILE_PATH"
	tlsKeyFilePathEnvVar  = "TLS_KEY_FILE_PATH"
)

// Start the gRPC proxy service
func main() {
	app := gin.Default()
	listenPort := ":" + os.Getenv(listenPortEnvVar)
	tlsCertFilePath := os.Getenv(tlsCertFilePathEnvVar)
	tlsKeyFilePath := os.Getenv(tlsKeyFilePathEnvVar)

	if listenPort == ":" {
		listenPort = ":" + defaultServerPort
	}

	// Handle requests to get a Loop In quote
	app.POST("/v0/loopInQuote", func(context *gin.Context) {
		loopInQuote(context)
	})

	// Handle requests to get Loop In terms
	app.POST("/v0/loopInTerms", func(context *gin.Context) {
		loopInTerms(context)
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

	if tlsCertFilePath != "" && tlsKeyFilePath != "" {
		app.RunTLS(listenPort, tlsCertFilePath, tlsKeyFilePath)
	} else {
		app.Run(listenPort)
	}
}
