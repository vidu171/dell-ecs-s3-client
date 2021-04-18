// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package apigatewaymanagementapi

import (
	"github.com/vidu171/dell-ecs-s3-client/private/protocol"
)

const (

	// ErrCodeForbiddenException for service response error code
	// "ForbiddenException".
	//
	// The caller is not authorized to invoke this operation.
	ErrCodeForbiddenException = "ForbiddenException"

	// ErrCodeGoneException for service response error code
	// "GoneException".
	//
	// The connection with the provided id no longer exists.
	ErrCodeGoneException = "GoneException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The client is sending more than the allowed number of requests per unit of
	// time or the WebSocket client side buffer is full.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodePayloadTooLargeException for service response error code
	// "PayloadTooLargeException".
	//
	// The data has exceeded the maximum size allowed.
	ErrCodePayloadTooLargeException = "PayloadTooLargeException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"ForbiddenException":       newErrorForbiddenException,
	"GoneException":            newErrorGoneException,
	"LimitExceededException":   newErrorLimitExceededException,
	"PayloadTooLargeException": newErrorPayloadTooLargeException,
}