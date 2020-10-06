package main

import (
	"github.com/alexbosworth/grpc-proxy/looprpc"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"net/http"
)

// Arguments for get loop in terms request
type loopInTermsRequest struct {
	ProtocolVersion string `json:"protocol_version" binding:"required"`
}

// Get the terms for loop in
func loopInTerms(c *gin.Context) {
	headers := requestHeaders{}

	// Exit early when the headers cannot be bound
	if err := c.ShouldBindHeader(&headers); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"details": err.Error()})
		return
	}

	// Map the request into its type
	var req loopInTermsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"details": err.Error()})
		return
	}

	_, hasProtocol := looprpc.ProtocolVersion_value[req.ProtocolVersion]

	// Exit early when the protocol is not defined
	if !hasProtocol {
		c.JSON(
			http.StatusBadRequest,
			gin.H{"details": "ExpectedKnownProtocolVersionToGetTerms"},
		)
		return
	}

	// Get a connection to the server
	client, connectionError := swapClient()

	// Exit early when there is a connection issue
	if connectionError != nil {
		c.JSON(
			http.StatusServiceUnavailable,
			gin.H{"details": "FailedToConnectToSwapServerForLoopInTerms"},
		)
		return
	}

	// Map the arguments into a request
	loopReq := looprpc.ServerLoopInTermsRequest{
		ProtocolVersion: looprpc.ProtocolVersion(
			looprpc.ProtocolVersion_value[req.ProtocolVersion],
		),
	}

	// gRPC responses have "trailers" for metadata in addition to headers
	responseTrailers := metadata.MD{}

	grpcResponse, loopErr := client.LoopInTerms(
		headers.asGrpcContext(),
		&loopReq,
		grpc.Trailer(&responseTrailers),
	)

	// Exit early when an error is encountered
	if loopErr != nil {
		returnGrpcError(c, loopErr, responseTrailers)
		return
	}

	// Return the response from the server
	c.JSON(http.StatusOK, grpcResponse)
}
