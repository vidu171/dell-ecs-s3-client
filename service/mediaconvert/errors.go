// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediaconvert

import (
	"github.com/vidu171/dell-ecs-s3-client/private/protocol"
)

const (

	// ErrCodeBadRequestException for service response error code
	// "BadRequestException".
	ErrCodeBadRequestException = "BadRequestException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	ErrCodeConflictException = "ConflictException"

	// ErrCodeForbiddenException for service response error code
	// "ForbiddenException".
	ErrCodeForbiddenException = "ForbiddenException"

	// ErrCodeInternalServerErrorException for service response error code
	// "InternalServerErrorException".
	ErrCodeInternalServerErrorException = "InternalServerErrorException"

	// ErrCodeNotFoundException for service response error code
	// "NotFoundException".
	ErrCodeNotFoundException = "NotFoundException"

	// ErrCodeTooManyRequestsException for service response error code
	// "TooManyRequestsException".
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
