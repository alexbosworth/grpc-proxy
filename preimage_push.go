package main

import (
	"encoding/hex"
	"errors"
	"github.com/alexbosworth/grpc-proxy/looprpc"
)

// Arguments for a Loop Out preimage push request
type preimagePush struct {
	Preimage        string `json:"preimage" binding:"required"`
	ProtocolVersion string `json:"protocol_version" binding:"required"`
}

// Derive gRPC request details
func (push preimagePush) asGrpcRequest() (
	*looprpc.ServerLoopOutPushPreimageRequest,
	error,
) {
	_, hasProtocol := looprpc.ProtocolVersion_value[push.ProtocolVersion]

	// Exit early when the protocol is not defined
	if !hasProtocol {
		return nil, errors.New("ExpectedKnownProtocolVersionForPreimagePush")
	}

	// Map the protocol version name to a numeric value
	protocolVersion := looprpc.ProtocolVersion(
		looprpc.ProtocolVersion_value[push.ProtocolVersion],
	)

	// Exit early when the preimage does not map to expected bytes
	preimage, decodePreimageErr := hex.DecodeString(push.Preimage)
	if decodePreimageErr != nil {
		return nil, errors.New("ExpectedValidPreimageHexValueForPreimagePush")
	}

	request := looprpc.ServerLoopOutPushPreimageRequest{
		Preimage:        preimage,
		ProtocolVersion: protocolVersion,
	}

	return &request, nil
}
