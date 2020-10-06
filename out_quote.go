package main

import (
	"errors"
	"github.com/alexbosworth/grpc-proxy/looprpc"
)

// Arguments for a Loop Out quote
type outQuote struct {
	Amt             uint64 `json:"amt,string,omitempty" binding:"required"`
	Expiry          int32  `json:"expiry" binding:"required"`
	Deadline        int64  `json:"swap_publication_deadline,string"`
	ProtocolVersion string `json:"protocol_version" binding:"required"`
}

// Derive gRPC request details
func (quote outQuote) asGrpcRequest() (
	*looprpc.ServerLoopOutQuoteRequest,
	error,
) {
	_, hasProtocol := looprpc.ProtocolVersion_value[quote.ProtocolVersion]

	// Exit early when the protocol is not defined
	if !hasProtocol {
		return nil, errors.New("ExpectedKnownProtocolVersionForLoopOutQuote")
	}

	// Map the protocol version name to a numeric value
	protocolVersion := looprpc.ProtocolVersion(
		looprpc.ProtocolVersion_value[quote.ProtocolVersion],
	)

	request := looprpc.ServerLoopOutQuoteRequest{
		Amt:                     quote.Amt,
		Expiry:                  quote.Expiry,
		ProtocolVersion:         protocolVersion,
		SwapPublicationDeadline: quote.Deadline,
	}

	return &request, nil
}
