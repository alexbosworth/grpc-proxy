package main

import (
	"crypto/tls"
	"github.com/alexbosworth/grpc-proxy/looprpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"os"
)

var (
	remoteServerSocketKey = "REMOTE_SERVER_SOCKET"
)

// Get a connected swap client that can be used to make swap gRPC calls
func swapClient() (client looprpc.SwapServerClient, err error) {
	creds := credentials.NewTLS(&tls.Config{})
	opts := []grpc.DialOption{}

	// Add TLS to the gRPC options
	opts = append(opts, grpc.WithTransportCredentials(creds))

	// Connect to the service
	connnection, err := grpc.Dial(os.Getenv(remoteServerSocketKey), opts...)

	// Exit early when there was an error connecting
	if err != nil {
		return nil, err
	}

	return looprpc.NewSwapServerClient(connnection), nil
}
