package main

import (
	"errors"
	"github.com/alexbosworth/grpc-proxy/looprpc"
)

// Arguments for a Loop In quote
type inQuote struct {
	Amt             uint64 `json:"amt,string,omitempty" binding:"required"`
	ProtocolVersion string `json:"protocol_version" binding:"required"`
}

// Derive gRPC request details
func (quote inQuote) asGrpcRequest() (
	*looprpc.ServerLoopInQuoteRequest,
	error,
) {
	_, hasProtocol := looprpc.ProtocolVersion_value[quote.ProtocolVersion]

	// Exit early when the protocol is not defined
	if !hasProtocol {
		return nil, errors.New("ExpectedKnownProtocolVersionForLoopInQuote")
	}

	// Map the protocol version name to a numeric value
	protocolVersion := looprpc.ProtocolVersion(
		looprpc.ProtocolVersion_value[quote.ProtocolVersion],
	)

	request := looprpc.ServerLoopInQuoteRequest{
		Amt:             quote.Amt,
		ProtocolVersion: protocolVersion,
	}

	return &request, nil
}
