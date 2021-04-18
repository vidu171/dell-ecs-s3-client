// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ivs

import (
	"github.com/vidu171/dell-ecs-s3-client/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeChannelNotBroadcasting for service response error code
	// "ChannelNotBroadcasting".
	ErrCodeChannelNotBroadcasting = "ChannelNotBroadcasting"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodePendingVerification for service response error code
	// "PendingVerification".
	ErrCodePendingVerification = "PendingVerification"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServiceQuotaExceededException for service response error code
	// "ServiceQuotaExceededException".
	ErrCodeServiceQuotaExceededException = "ServiceQuotaExceededException"

	// ErrCodeStreamUnavailable for service response error code
	// "StreamUnavailable".
	ErrCodeStreamUnavailable = "StreamUnavailable"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":         newErrorAccessDeniedException,
	"ChannelNotBroadcasting":        newErrorChannelNotBroadcasting,
	"ConflictException":             newErrorConflictException,
	"InternalServerException":       newErrorInternalServerException,
	"PendingVerification":           newErrorPendingVerification,
	"ResourceNotFoundException":     newErrorResourceNotFoundException,
	"ServiceQuotaExceededException": newErrorServiceQuotaExceededException,
	"StreamUnavailable":             newErrorStreamUnavailable,
	"ThrottlingException":           newErrorThrottlingException,
	"ValidationException":           newErrorValidationException,
}
