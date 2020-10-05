package main

import (
	"context"
	"google.golang.org/grpc/metadata"
	"time"
)

var (
	serviceResponseTimeout = time.Minute
)

// Known headers for a REST request
type requestHeaders struct {
	Authorization string `header:"authorization"`
}

// Translate REST request headers into gRPC request headers
func (headers requestHeaders) asGrpcContext() context.Context {
	ctx, _ := context.WithTimeout(context.Background(), serviceResponseTimeout)

	// Create a context that the gRPC request will operate in
	requestContext := metadata.AppendToOutgoingContext(
		ctx,
		"authorization",
		headers.Authorization,
	)

	return requestContext
}
