// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotthingsgraph

import (
	"github.com/vidu171/dell-ecs-s3-client/private/protocol"
)

const (

	// ErrCodeInternalFailureException for service response error code
	// "InternalFailureException".
	ErrCodeInternalFailureException = "InternalFailureException"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	ErrCodeInvalidRequestException = "InvalidRequestException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeResourceAlreadyExistsException for service response error code
	// "ResourceAlreadyExistsException".
	ErrCodeResourceAlreadyExistsException = "ResourceAlreadyExistsException"

	// ErrCodeResourceInUseException for service response error code
	// "ResourceInUseException".
	ErrCodeResourceInUseException = "ResourceInUseException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	ErrCodeThrottlingException = "ThrottlingException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"InternalFailureException":       newErrorInternalFailureException,
	"InvalidRequestException":        newErrorInvalidRequestException,
	"LimitExceededException":         newErrorLimitExceededException,
	"ResourceAlreadyExistsException": newErrorResourceAlreadyExistsException,
	"ResourceInUseException":         newErrorResourceInUseException,
	"ResourceNotFoundException":      newErrorResourceNotFoundException,
	"ThrottlingException":            newErrorThrottlingException,
}
