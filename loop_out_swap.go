package main

import (
	"encoding/hex"
	"errors"
	"github.com/alexbosworth/grpc-proxy/looprpc"
)

// Arguments for a Loop Out request
type loopOutSwap struct {
	Amt             uint64 `json:"amt,string,omitempty" binding:"required"`
	Deadline        int64  `json:"swap_publication_deadline"`
	Expiry          int32  `json:"expiry" binding:"required"`
	ProtocolVersion string `json:"protocol_version" binding:"required"`
	ReceiverKey     string `json:"receiver_key" binding:"required"`
	SwapHash        string `json:"swap_hash" binding:"required"`
}

// Derive gRPC request details for a loopOutSwap
func (swap loopOutSwap) asGrpcRequest() (*looprpc.ServerLoopOutRequest, error) {
	_, hasProtocol := looprpc.ProtocolVersion_value[swap.ProtocolVersion]

	// Exit early when the protocol is not defined
	if !hasProtocol {
		return nil, errors.New("ExpectedKnownProtocolVersionForLoopOut")
	}

	protocolVersion := looprpc.ProtocolVersion(
		looprpc.ProtocolVersion_value[swap.ProtocolVersion],
	)

	// Exit early when the receiver key does not map to expected bytes
	receiverKey, decodeReceiverKeyErr := hex.DecodeString(swap.ReceiverKey)
	if decodeReceiverKeyErr != nil {
		return nil, errors.New("ExpectedValidReceiverKeyHexValue")
	}

	// Exit early when the swap hash does not map to expected bytes
	swapHash, decodeSwapHashErr := hex.DecodeString(swap.SwapHash)
	if decodeSwapHashErr != nil {
		return nil, errors.New("ExpectedValidSwapHashHexValue")
	}

	request := looprpc.ServerLoopOutRequest{
		Amt:                     swap.Amt,
		Expiry:                  swap.Expiry,
		ProtocolVersion:         protocolVersion,
		ReceiverKey:             receiverKey,
		SwapHash:                swapHash,
		SwapPublicationDeadline: swap.Deadline,
	}

	return &request, nil
}
