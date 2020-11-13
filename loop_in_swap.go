package main

import (
	"encoding/hex"
	"errors"
	"github.com/alexbosworth/grpc-proxy/looprpc"
)

// Arguments for a Loop In request
type loopInSwap struct {
	Amt             uint64 `json:"amt,string,omitempty" binding:"required"`
	LastHop         string `json:"last_hop"`
	ProbeInvoice    string `json:"probe_invoice" binding:"required"`
	ProtocolVersion string `json:"protocol_version" binding:"required"`
	SenderKey       string `json:"sender_key" binding:"required"`
	SwapHash        string `json:"swap_hash" binding:"required"`
	SwapInvoice     string `json:"swap_invoice" binding:"required"`
	UserAgent       string `json:"user_agent" binding:"required"`
}

// Derive gRPC request details for a loopInSwap
func (swap loopInSwap) asGrpcRequest() (*looprpc.ServerLoopInRequest, error) {
	// Exit early when the swap hash does not map to expected bytes
	lastHop, decodeLastHopErr := hex.DecodeString(swap.LastHop)
	if decodeLastHopErr != nil {
		return nil, errors.New("ExpectedValidSwapHashHexValue")
	}

	_, hasProtocol := looprpc.ProtocolVersion_value[swap.ProtocolVersion]

	// Exit early when the protocol is not defined
	if !hasProtocol {
		return nil, errors.New("ExpectedKnownProtocolVersionForLoopIn")
	}

	protocolVersion := looprpc.ProtocolVersion(
		looprpc.ProtocolVersion_value[swap.ProtocolVersion],
	)

	// Exit early when the sender key does not map to expected bytes
	senderKey, decodeSenderKeyErr := hex.DecodeString(swap.SenderKey)
	if decodeSenderKeyErr != nil {
		return nil, errors.New("ExpectedValidSenderKeyHexValueForLoopIn")
	}

	// Exit early when the swap hash does not map to expected bytes
	swapHash, decodeSwapHashErr := hex.DecodeString(swap.SwapHash)
	if decodeSwapHashErr != nil {
		return nil, errors.New("ExpectedValidSwapHashHexValueForLoopIn")
	}

	request := looprpc.ServerLoopInRequest{
		Amt:             swap.Amt,
		LastHop:         lastHop,
		ProbeInvoice:    swap.ProbeInvoice,
		ProtocolVersion: protocolVersion,
		SenderKey:       senderKey,
		SwapHash:        swapHash,
		SwapInvoice:     swap.SwapInvoice,
		UserAgent:       swap.UserAgent,
	}

	return &request, nil
}
