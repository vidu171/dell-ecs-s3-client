// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package mediastore

import (
	"github.com/vidu171/dell-ecs-s3-client/private/protocol"
)

const (

	// ErrCodeContainerInUseException for service response error code
	// "ContainerInUseException".
	//
	// The container that you specified in the request already exists or is being
	// updated.
	ErrCodeContainerInUseException = "ContainerInUseException"

	// ErrCodeContainerNotFoundException for service response error code
	// "ContainerNotFoundException".
	//
	// The container that you specified in the request does not exist.
	ErrCodeContainerNotFoundException = "ContainerNotFoundException"

	// ErrCodeCorsPolicyNotFoundException for service response error code
	// "CorsPolicyNotFoundException".
	//
	// The CORS policy that you specified in the request does not exist.
	ErrCodeCorsPolicyNotFoundException = "CorsPolicyNotFoundException"

	// ErrCodeInternalServerError for service response error code
	// "InternalServerError".
	//
	// The service is temporarily unavailable.
	ErrCodeInternalServerError = "InternalServerError"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// A service limit has been exceeded.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodePolicyNotFoundException for service response error code
	// "PolicyNotFoundException".
	//
	// The policy that you specified in the request does not exist.
	ErrCodePolicyNotFoundException = "PolicyNotFoundException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"ContainerInUseException":     newErrorContainerInUseException,
	"ContainerNotFoundException":  newErrorContainerNotFoundException,
	"CorsPolicyNotFoundException": newErrorCorsPolicyNotFoundException,
	"InternalServerError":         newErrorInternalServerError,
	"LimitExceededException":      newErrorLimitExceededException,
	"PolicyNotFoundException":     newErrorPolicyNotFoundException,
}
