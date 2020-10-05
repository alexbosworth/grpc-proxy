package main

import (
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net/http"
)

// Return a gRPC error as a REST response
func returnGrpcError(context *gin.Context, err error, trailers metadata.MD) {
	serviceErr, _ := status.FromError(err)

	// Exit early when there is a special payment required error
	if serviceErr.Message() == "payment required" {
		context.JSON(
			http.StatusPaymentRequired,
			gin.H{
				"authenticate": trailers["www-authenticate"],
				"details":      serviceErr.Message(),
			},
		)
		return
	}

	context.JSON(
		http.StatusServiceUnavailable,
		gin.H{"details": serviceErr.Message()},
	)

	return
}
