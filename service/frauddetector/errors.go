// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package frauddetector

import (
	"github.com/vidu171/dell-ecs-s3-client/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// An exception indicating Amazon Fraud Detector does not have the needed permissions.
	// This can occur if you submit a request, such as PutExternalModel, that specifies
	// a role that is not in your account.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeConflictException for service response error code
	// "ConflictException".
	//
	// An exception indicating there was a conflict during a delete operation. The
	// following delete operations can cause a conflict exception:
	//
	//    * DeleteDetector: A conflict exception will occur if the detector has
	//    associated Rules or DetectorVersions. You can only delete a detector if
	//    it has no Rules or DetectorVersions.
	//
	//    * DeleteDetectorVersion: A conflict exception will occur if the DetectorVersion
	//    status is ACTIVE.
	//
	//    * DeleteRule: A conflict exception will occur if the RuleVersion is in
	//    use by an associated ACTIVE or INACTIVE DetectorVersion.
	ErrCodeConflictException = "ConflictException"

	// ErrCodeInternalServerException for service response error code
	// "InternalServerException".
	//
	// An exception indicating an internal server error.
	ErrCodeInternalServerException = "InternalServerException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// An exception indicating the specified resource was not found.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeThrottlingException for service response error code
	// "ThrottlingException".
	//
	// An exception indicating a throttling error.
	ErrCodeThrottlingException = "ThrottlingException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// An exception indicating a specified value is not allowed.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":     newErrorAccessDeniedException,
	"ConflictException":         newErrorConflictException,
	"InternalServerException":   newErrorInternalServerException,
	"ResourceNotFoundException": newErrorResourceNotFoundException,
	"ThrottlingException":       newErrorThrottlingException,
	"ValidationException":       newErrorValidationException,
}
