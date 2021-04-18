// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package serverlessapplicationrepository

import (
	"github.com/vidu171/dell-ecs-s3-client/private/protocol"
)

const (

	// ErrCodeBadRequestException for service response error code
	// "BadRequestException".
	//
	// One of the parameters in the request is invalid.
	ErrCodeBadRequestException = "BadRequestException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// The resource already exists.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeForbiddenException for service response error code
	// "ForbiddenException".
	//
	// The client is not authenticated.
	ErrCodeForbiddenException = "ForbiddenException"

	// ErrCodeInternalServerErrorException for service response error code
	// "InternalServerErrorException".
	//
	// The AWS Serverless Application Repository service encountered an internal
	// error.
	ErrCodeInternalServerErrorException = "InternalServerErrorException"

	// ErrCodeNotFoundException for service response error code
	// "NotFoundException".
	//
	// The resource (for example, an access policy statement) specified in the request
	// doesn't exist.
	ErrCodeNotFoundException = "NotFoundException"

	// ErrCodeTooManyRequestsException for service response error code
	// "TooManyRequestsException".
	//
	// The client is sending more than the allowed number of requests per unit of
	// time.
	ErrCodeTooManyRequestsException = "TooManyRequestsException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"BadRequestException":          newErrorBadRequestException,
	"ConflictException":            newErrorConflictException,
	"ForbiddenException":           newErrorForbiddenException,
	"InternalServerErrorException": newErrorInternalServerErrorException,
	"NotFoundException":            newErrorNotFoundException,
	"TooManyRequestsException":     newErrorTooManyRequestsException,
}
