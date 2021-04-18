// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package sagemakerfeaturestoreruntime

import (
	"github.com/vidu171/dell-ecs-s3-client/private/protocol"
)

const (

	// ErrCodeAccessForbidden for service response error code
	// "AccessForbidden".
	//
	// You do not have permission to perform an action.
	ErrCodeAccessForbidden = "AccessForbidden"

	// ErrCodeInternalFailure for service response error code
	// "InternalFailure".
	//
	// An internal failure occurred. Try your request again. If the problem persists,
	// contact AWS customer support.
	ErrCodeInternalFailure = "InternalFailure"

	// ErrCodeResourceNotFound for service response error code
	// "ResourceNotFound".
	//
	// A resource that is required to perform an action was not found.
	ErrCodeResourceNotFound = "ResourceNotFound"

	// ErrCodeServiceUnavailable for service response error code
	// "ServiceUnavailable".
	//
	// The service is currently unavailable.
	ErrCodeServiceUnavailable = "ServiceUnavailable"

	// ErrCodeValidationError for service response error code
	// "ValidationError".
	//
	// There was an error validating your request.
	ErrCodeValidationError = "ValidationError"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessForbidden":    newErrorAccessForbidden,
	"InternalFailure":    newErrorInternalFailure,
	"ResourceNotFound":   newErrorResourceNotFound,
	"ServiceUnavailable": newErrorServiceUnavailable,
	"ValidationError":    newErrorValidationError,
}
