// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package kinesisvideo

import (
	"github.com/vidu171/dell-ecs-s3-client/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	//
	// You do not have required permissions to perform this operation.
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeAccountChannelLimitExceededException for service response error code
	// "AccountChannelLimitExceededException".
	//
	// You have reached the maximum limit of active signaling channels for this
	// AWS account in this region.
	ErrCodeAccountChannelLimitExceededException = "AccountChannelLimitExceededException"

	// ErrCodeAccountStreamLimitExceededException for service response error code
	// "AccountStreamLimitExceededException".
	//
	// The number of streams created for the account is too high.
	ErrCodeAccountStreamLimitExceededException = "AccountStreamLimitExceededException"

	// ErrCodeClientLimitExceededException for service response error code
	// "ClientLimitExceededException".
	//
	// Kinesis Video Streams has throttled the request because you have exceeded
	// the limit of allowed client calls. Try making the call later.
	ErrCodeClientLimitExceededException = "ClientLimitExceededException"

	// ErrCodeDeviceStreamLimitExceededException for service response error code
	// "DeviceStreamLimitExceededException".
	//
	// Not implemented.
	ErrCodeDeviceStreamLimitExceededException = "DeviceStreamLimitExceededException"

	// ErrCodeInvalidArgumentException for service response error code
	// "InvalidArgumentException".
	//
	// The value for this input parameter is invalid.
	ErrCodeInvalidArgumentException = "InvalidArgumentException"

	// ErrCodeInvalidDeviceException for service response error code
	// "InvalidDeviceException".
	//
	// Not implemented.
	ErrCodeInvalidDeviceException = "InvalidDeviceException"

	// ErrCodeInvalidResourceFormatException for service response error code
	// "InvalidResourceFormatException".
	//
	// The format of the StreamARN is invalid.
	ErrCodeInvalidResourceFormatException = "InvalidResourceFormatException"

	// ErrCodeNotAuthorizedException for service response error code
	// "NotAuthorizedException".
	//
	// The caller is not authorized to perform this operation.
	ErrCodeNotAuthorizedException = "NotAuthorizedException"

	// ErrCodeResourceInUseException for service response error code
	// "ResourceInUseException".
	//
	// The signaling channel is currently not available for this operation.
	ErrCodeResourceInUseException = "ResourceInUseException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// Amazon Kinesis Video Streams can't find the stream that you specified.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeTagsPerResourceExceededLimitException for service response error code
	// "TagsPerResourceExceededLimitException".
	//
	// You have exceeded the limit of tags that you can associate with the resource.
	// Kinesis video streams support up to 50 tags.
	ErrCodeTagsPerResourceExceededLimitException = "TagsPerResourceExceededLimitException"

	// ErrCodeVersionMismatchException for service response error code
	// "VersionMismatchException".
	//
	// The stream version that you specified is not the latest version. To get the
	// latest version, use the DescribeStream (https://docs.aws.amazon.com/kinesisvideostreams/latest/dg/API_DescribeStream.html)
	// API.
	ErrCodeVersionMismatchException = "VersionMismatchException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":                 newErrorAccessDeniedException,
	"AccountChannelLimitExceededException":  newErrorAccountChannelLimitExceededException,
	"AccountStreamLimitExceededException":   newErrorAccountStreamLimitExceededException,
	"ClientLimitExceededException":          newErrorClientLimitExceededException,
	"DeviceStreamLimitExceededException":    newErrorDeviceStreamLimitExceededException,
	"InvalidArgumentException":              newErrorInvalidArgumentException,
	"InvalidDeviceException":                newErrorInvalidDeviceException,
	"InvalidResourceFormatException":        newErrorInvalidResourceFormatException,
	"NotAuthorizedException":                newErrorNotAuthorizedException,
	"ResourceInUseException":                newErrorResourceInUseException,
	"ResourceNotFoundException":             newErrorResourceNotFoundException,
	"TagsPerResourceExceededLimitException": newErrorTagsPerResourceExceededLimitException,
	"VersionMismatchException":              newErrorVersionMismatchException,
}
