// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package secretsmanager

import (
	"github.com/vidu171/dell-ecs-s3-client/private/protocol"
)

const (

	// ErrCodeDecryptionFailure for service response error code
	// "DecryptionFailure".
	//
	// Secrets Manager can't decrypt the protected secret text using the provided
	// KMS key.
	ErrCodeDecryptionFailure = "DecryptionFailure"

	// ErrCodeEncryptionFailure for service response error code
	// "EncryptionFailure".
	//
	// Secrets Manager can't encrypt the protected secret text using the provided
	// KMS key. Check that the customer master key (CMK) is available, enabled,
	// and not in an invalid state. For more information, see How Key State Affects
	// Use of a Customer Master Key (http://docs.aws.amazon.com/kms/latest/developerguide/key-state.html).
	ErrCodeEncryptionFailure = "EncryptionFailure"

	// ErrCodeInternalServiceError for service response error code
	// "InternalServiceError".
	//
	// An error occurred on the server side.
	ErrCodeInternalServiceError = "InternalServiceError"

	// ErrCodeInvalidNextTokenException for service response error code
	// "InvalidNextTokenException".
	//
	// You provided an invalid NextToken value.
	ErrCodeInvalidNextTokenException = "InvalidNextTokenException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	//
	// You provided an invalid value for a parameter.
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeInvalidRequestException for service response error code
	// "InvalidRequestException".
	//
	// You provided a parameter value that is not valid for the current state of
	// the resource.
	//
	// Possible causes:
	//
	//    * You tried to perform the operation on a secret that's currently marked
	//    deleted.
	//
	//    * You tried to enable rotation on a secret that doesn't already have a
	//    Lambda function ARN configured and you didn't include such an ARN as a
	//    parameter in this call.
	ErrCodeInvalidRequestException = "InvalidRequestException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// The request failed because it would exceed one of the Secrets Manager internal
	// limits.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeMalformedPolicyDocumentException for service response error code
	// "MalformedPolicyDocumentException".
	//
	// You provided a resource-based policy with syntax errors.
	ErrCodeMalformedPolicyDocumentException = "MalformedPolicyDocumentException"

	// ErrCodePreconditionNotMetException for service response error code
	// "PreconditionNotMetException".
	//
	// The request failed because you did not complete all the prerequisite steps.
	ErrCodePreconditionNotMetException = "PreconditionNotMetException"

	// ErrCodePublicPolicyException for service response error code
	// "PublicPolicyException".
	//
	// The BlockPublicPolicy parameter is set to true and the resource policy did
	// not prevent broad access to the secret.
	ErrCodePublicPolicyException = "PublicPolicyException"

	// ErrCodeResourceExistsException for service response error code
	// "ResourceExistsException".
	//
	// A resource with the ID you requested already exists.
	ErrCodeResourceExistsException = "ResourceExistsException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// We can't find the resource that you asked for.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"DecryptionFailure":                newErrorDecryptionFailure,
	"EncryptionFailure":                newErrorEncryptionFailure,
	"InternalServiceError":             newErrorInternalServiceError,
	"InvalidNextTokenException":        newErrorInvalidNextTokenException,
	"InvalidParameterException":        newErrorInvalidParameterException,
	"InvalidRequestException":          newErrorInvalidRequestException,
	"LimitExceededException":           newErrorLimitExceededException,
	"MalformedPolicyDocumentException": newErrorMalformedPolicyDocumentException,
	"PreconditionNotMetException":      newErrorPreconditionNotMetException,
	"PublicPolicyException":            newErrorPublicPolicyException,
	"ResourceExistsException":          newErrorResourceExistsException,
	"ResourceNotFoundException":        newErrorResourceNotFoundException,
}
