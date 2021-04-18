// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package forecastqueryservice

import (
	"github.com/vidu171/dell-ecs-s3-client/private/protocol"
)

const (

	// ErrCodeInvalidInputException for service response error code
	// "InvalidInputException".
	//
	// The value is invalid or is too long.
	ErrCodeInvalidInputException = "InvalidInputException"

	// ErrCodeInvalidNextTokenException for service response error code
	// "InvalidNextTokenException".
	//
	// The token is not valid. Tokens expire after 24 hours.
	ErrCodeInvalidNextTokenException = "InvalidNextTokenException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The limit on the number of requests per second has been exceeded.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeResourceInUseException for service response error code
	// "ResourceInUseException".
	//
	// The specified resource is in use.
	ErrCodeResourceInUseException = "ResourceInUseException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// We can't find that resource. Check the information that you've provided and
	// try again.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"InvalidInputException":     newErrorInvalidInputException,
	"InvalidNextTokenException": newErrorInvalidNextTokenException,
	"LimitExceededException":    newErrorLimitExceededException,
	"ResourceInUseException":    newErrorResourceInUseException,
	"ResourceNotFoundException": newErrorResourceNotFoundException,
}
