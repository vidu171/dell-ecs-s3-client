// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemakerruntime

import (
	"github.com/vidu171/dell-ecs-s3-client/private/protocol"
)

const (

	// ErrCodeInternalFailure for service response error code
	// "InternalFailure".
	//
	// An internal failure occurred.
	ErrCodeInternalFailure = "InternalFailure"

	// ErrCodeModelError for service response error code
	// "ModelError".
	//
	// Model (owned by the customer in the container) returned 4xx or 5xx error
	// code.
	ErrCodeModelError = "ModelError"

	// ErrCodeServiceUnavailable for service response error code
	// "ServiceUnavailable".
	//
	// The service is unavailable. Try your call again.
	ErrCodeServiceUnavailable = "ServiceUnavailable"

	// ErrCodeValidationError for service response error code
	// "ValidationError".
	//
	// Inspect your request and try again.
	ErrCodeValidationError = "ValidationError"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"InternalFailure":    newErrorInternalFailure,
	"ModelError":         newErrorModelError,
	"ServiceUnavailable": newErrorServiceUnavailable,
	"ValidationError":    newErrorValidationError,
}
