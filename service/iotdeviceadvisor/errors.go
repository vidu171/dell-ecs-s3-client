// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package iotdeviceadvisor

import (
	"github.com/vidu171/dell-ecs-s3-client/private/protocol"
)

const (

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// Sends Conflict Exception.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// Sends Internal Failure Exception.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// Sends Resource Not Found Exception.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// Sends invalid request exception.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"ConflictException":         newErrorConflictException,
	"InternalServerException":   newErrorInternalServerException,
	"ResourceNotFoundException": newErrorResourceNotFoundException,
	"ValidationException":       newErrorValidationException,
}
