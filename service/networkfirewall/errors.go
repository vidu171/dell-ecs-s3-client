// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package networkfirewall

import (
	"github.com/vidu171/dell-ecs-s3-client/private/protocol"
)

const (

	// ErrCodeInsufficientCapacityException for service response error code
	// "InsufficientCapacityException".
	//
	// AWS doesn't currently have enough available capacity to fulfill your request.
	// Try your request later.
	ErrCodeInsufficientCapacityException = "InsufficientCapacityException"

	// ErrCodeInternalServerError for service response error code
	// "InternalServerError".
	//
	// Your request is valid, but Network Firewall couldn’t perform the operation
	// because of a system problem. Retry your request.
	ErrCodeInternalServerError = "InternalServerError"

	// ErrCodeInvalidOperationException for service response error code
	// "InvalidOperationException".
	//
	// The operation failed because it's not valid. For example, you might have
	// tried to delete a rule group or firewall policy that's in use.
	ErrCodeInvalidOperationException = "InvalidOperationException"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	//
	// The operation failed because of a problem with your request. Examples include:
	//
	//    * You specified an unsupported parameter name or value.
	//
	//    * You tried to update a property with a value that isn't among the available
	//    types.
	//
	//    * Your request references an ARN that is malformed, or corresponds to
	//    a resource that isn't valid in the context of the request.
	ErrCodeInvalidRequestException = "InvalidRequestException"

	// ErrCodeInvalidResourcePolicyException for service response error code
	// "InvalidResourcePolicyException".
	ErrCodeInvalidResourcePolicyException = "InvalidResourcePolicyException"

	// ErrCodeInvalidTokenException for service response error code
	// "InvalidTokenException".
	//
	// The token you provided is stale or isn't valid for the operation.
	ErrCodeInvalidTokenException = "InvalidTokenException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// Unable to perform the operation because doing so would violate a limit setting.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeLogDestinationPermissionException for service response error code
	// "LogDestinationPermissionException".
	//
	// Unable to send logs to a configured logging destination.
	ErrCodeLogDestinationPermissionException = "LogDestinationPermissionException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// Unable to locate a resource using the parameters that you provided.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeResourceOwnerCheckException for service response error code
	// "ResourceOwnerCheckException".
	ErrCodeResourceOwnerCheckException = "ResourceOwnerCheckException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// Unable to process the request due to throttling limitations.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeUnsupportedOperationException for service response error code
	// "UnsupportedOperationException".
	//
	// The operation you requested isn't supported by Network Firewall.
	ErrCodeUnsupportedOperationException = "UnsupportedOperationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"InsufficientCapacityException":     newErrorInsufficientCapacityException,
	"InternalServerError":               newErrorInternalServerError,
	"InvalidOperationException":         newErrorInvalidOperationException,
	"InvalidRequestException":           newErrorInvalidRequestException,
	"InvalidResourcePolicyException":    newErrorInvalidResourcePolicyException,
	"InvalidTokenException":             newErrorInvalidTokenException,
	"LimitExceededException":            newErrorLimitExceededException,
	"LogDestinationPermissionException": newErrorLogDestinationPermissionException,
	"ResourceNotFoundException":         newErrorResourceNotFoundException,
	"ResourceOwnerCheckException":       newErrorResourceOwnerCheckException,
	"ThrottlingException":               newErrorThrottlingException,
	"UnsupportedOperationException":     newErrorUnsupportedOperationException,
}
