package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
)

// Start a new Loop Out swap
func newLoopOutSwap(c *gin.Context) {
	headers := requestHeaders{}

	// Exit early when the headers cannot be bound
	if err := c.ShouldBindHeader(&headers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"details": err.Error()})
		return
	}

	var req loopOutSwap

	// Exit early when the JSON arguments are incorrect types
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"details": err.Error()})
		return
	}

	loopReq, reqErr := req.asGrpcRequest()

	// Exit early when there is an error validating the request arguments
	if reqErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"details": reqErr.Error()})
		return
	}

	// Get a connection to the server
	client, connectErr := swapClient()

	// Exit early when there is a connection issue
	if connectErr != nil {
		c.JSON(
			http.StatusServiceUnavailable,
			gin.H{"details": "FailedToConnectToSwapServer"},
		)
		return
	}

	// gRPC responses have "trailers" for metadata in addition to headers
	responseTrailers := metadata.MD{}

	// Initiate the gRPC request to the server
	grpcResponse, loopErr := client.NewLoopOutSwap(
		headers.asGrpcContext(),
		loopReq,
		grpc.Trailer(&responseTrailers),
	)

	// Exit early when there was an error returned by the gRPC response
	if loopErr != nil {
		returnGrpcError(c, loopErr, responseTrailers)
		return
	}

	// Return the response from the server
	c.JSON(http.StatusOK, grpcResponse)
}
