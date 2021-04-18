// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backup

import (
	"github.com/vidu171/dell-ecs-s3-client/private/protocol"
)

const (

	// ErrCodeAlreadyExistsException for service response error code
	// "AlreadyExistsException".
	//
	// The required resource already exists.
	ErrCodeAlreadyExistsException = "AlreadyExistsException"

	// ErrCodeDependencyFailureException for service response error code
	// "DependencyFailureException".
	//
	// A dependent AWS service or resource returned an error to the AWS Backup service,
	// and the action cannot be completed.
	ErrCodeDependencyFailureException = "DependencyFailureException"

	// ErrCodeInvalidParameterValueException for service response error code
	// "InvalidParameterValueException".
	//
	// Indicates that something is wrong with a parameter's value. For example,
	// the value is out of range.
	ErrCodeInvalidParameterValueException = "InvalidParameterValueException"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	//
	// Indicates that something is wrong with the input to the request. For example,
	// a parameter is of the wrong type.
	ErrCodeInvalidRequestException = "InvalidRequestException"

	// ErrCodeInvalidResourceStateException for service response error code
	// "InvalidResourceStateException".
	//
	// AWS Backup is already performing an action on this recovery point. It can't
	// perform the action you requested until the first action finishes. Try again
	// later.
	ErrCodeInvalidResourceStateException = "InvalidResourceStateException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// A limit in the request has been exceeded; for example, a maximum number of
	// items allowed in a request.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeMissingParameterValueException for service response error code
	// "MissingParameterValueException".
	//
	// Indicates that a required parameter is missing.
	ErrCodeMissingParameterValueException = "MissingParameterValueException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// A resource that is required for the action doesn't exist.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceUnavailableException for service response error code
	// "ServiceUnavailableException".
	//
	// The request failed due to a temporary failure of the server.
	ErrCodeServiceUnavailableException = "ServiceUnavailableException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AlreadyExistsException":         newErrorAlreadyExistsException,
	"DependencyFailureException":     newErrorDependencyFailureException,
	"InvalidParameterValueException": newErrorInvalidParameterValueException,
	"InvalidRequestException":        newErrorInvalidRequestException,
	"InvalidResourceStateException":  newErrorInvalidResourceStateException,
	"LimitExceededException":         newErrorLimitExceededException,
	"MissingParameterValueException": newErrorMissingParameterValueException,
	"ResourceNotFoundException":      newErrorResourceNotFoundException,
	"ServiceUnavailableException":    newErrorServiceUnavailableException,
}
