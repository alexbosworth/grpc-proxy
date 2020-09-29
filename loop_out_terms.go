package main

import (
	"context"
	"github.com/alexbosworth/grpc-proxy/looprpc"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Arguments for get loop out terms request
type loopOutTermsRequest struct {
	ProtocolVersion string `json:"protocol_version" binding:"required"`
}

// Get the terms for loop out
func loopOutTerms(c *gin.Context) {
	// Map the request into its type
	var req loopOutTermsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	_, hasProtocol := looprpc.ProtocolVersion_value[req.ProtocolVersion]

	// Exit early when the protocol is not defined
	if !hasProtocol {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"message": "ExpectedKnownProtocolVersionToGetTerms"},
		)
		return
	}

	// Get a connection to the server
	client, connectionError := swapClient()

	// Exit early when there is a connection issue
	if connectionError != nil {
		c.JSON(
			http.StatusServiceUnavailable,
			gin.H{"message": "FailedToConnectToSwapServer"},
		)
		return
	}

	// Map the arguments into a request
	loopReq := looprpc.ServerLoopOutTermsRequest{
		ProtocolVersion: looprpc.ProtocolVersion(
			looprpc.ProtocolVersion_value[req.ProtocolVersion],
		),
	}

	response, err := client.LoopOutTerms(context.Background(), &loopReq)

	// Pass back errors when present
	if err != nil {
		c.JSON(
			http.StatusServiceUnavailable,
			gin.H{"message": "UnexpectedErrorFromSwapServer"},
		)
		return
	}

	// Return the response from the server
	c.JSON(http.StatusOK, response)
}
