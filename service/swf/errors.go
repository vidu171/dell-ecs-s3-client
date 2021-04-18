// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package swf

import (
	"github.com/vidu171/dell-ecs-s3-client/private/protocol"
)

const (

	// ErrCodeDefaultUndefinedFault for service response error code
	// "DefaultUndefinedFault".
	//
	// The StartWorkflowExecution API action was called without the required parameters
	// set.
	//
	// Some workflow execution parameters, such as the decision taskList, must be
	// set to start the execution. However, these parameters might have been set
	// as defaults when the workflow type was registered. In this case, you can
	// omit these parameters from the StartWorkflowExecution call and Amazon SWF
	// uses the values defined in the workflow type.
	//
	// If these parameters aren't set and no default parameters were defined in
	// the workflow type, this error is displayed.
	ErrCodeDefaultUndefinedFault = "DefaultUndefinedFault"

	// ErrCodeDomainAlreadyExistsFault for service response error code
	// "DomainAlreadyExistsFault".
	//
	// Returned if the domain already exists. You may get this fault if you are
	// registering a domain that is either already registered or deprecated, or
	// if you undeprecate a domain that is currently registered.
	ErrCodeDomainAlreadyExistsFault = "DomainAlreadyExistsFault"

	// ErrCodeDomainDeprecatedFault for service response error code
	// "DomainDeprecatedFault".
	//
	// Returned when the specified domain has been deprecated.
	ErrCodeDomainDeprecatedFault = "DomainDeprecatedFault"

	// ErrCodeLimitExceededFault for service response error code
	// "LimitExceededFault".
	//
	// Returned by any operation if a system imposed limitation has been reached.
	// To address this fault you should either clean up unused resources or increase
	// the limit by contacting AWS.
	ErrCodeLimitExceededFault = "LimitExceededFault"

	// ErrCodeOperationNotPermittedFault for service response error code
	// "OperationNotPermittedFault".
	//
	// Returned when the caller doesn't have sufficient permissions to invoke the
	// action.
	ErrCodeOperationNotPermittedFault = "OperationNotPermittedFault"

	// ErrCodeTooManyTagsFault for service response error code
	// "TooManyTagsFault".
	//
	// You've exceeded the number of tags allowed for a domain.
	ErrCodeTooManyTagsFault = "TooManyTagsFault"

	// ErrCodeTypeAlreadyExistsFault for service response error code
	// "TypeAlreadyExistsFault".
	//
	// Returned if the type already exists in the specified domain. You may get
	// this fault if you are registering a type that is either already registered
	// or deprecated, or if you undeprecate a type that is currently registered.
	ErrCodeTypeAlreadyExistsFault = "TypeAlreadyExistsFault"

	// ErrCodeTypeDeprecatedFault for service response error code
	// "TypeDeprecatedFault".
	//
	// Returned when the specified activity or workflow type was already deprecated.
	ErrCodeTypeDeprecatedFault = "TypeDeprecatedFault"

	// ErrCodeUnknownResourceFault for service response error code
	// "UnknownResourceFault".
	//
	// Returned when the named resource cannot be found with in the scope of this
	// operation (region or domain). This could happen if the named resource was
	// never created or is no longer available for this operation.
	ErrCodeUnknownResourceFault = "UnknownResourceFault"

	// ErrCodeWorkflowExecutionAlreadyStartedFault for service response error code
	// "WorkflowExecutionAlreadyStartedFault".
	//
	// Returned by StartWorkflowExecution when an open execution with the same workflowId
	// is already running in the specified domain.
	ErrCodeWorkflowExecutionAlreadyStartedFault = "WorkflowExecutionAlreadyStartedFault"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"DefaultUndefinedFault":                newErrorDefaultUndefinedFault,
	"DomainAlreadyExistsFault":             newErrorDomainAlreadyExistsFault,
	"DomainDeprecatedFault":                newErrorDomainDeprecatedFault,
	"LimitExceededFault":                   newErrorLimitExceededFault,
	"OperationNotPermittedFault":           newErrorOperationNotPermittedFault,
	"TooManyTagsFault":                     newErrorTooManyTagsFault,
	"TypeAlreadyExistsFault":               newErrorTypeAlreadyExistsFault,
	"TypeDeprecatedFault":                  newErrorTypeDeprecatedFault,
	"UnknownResourceFault":                 newErrorUnknownResourceFault,
	"WorkflowExecutionAlreadyStartedFault": newErrorWorkflowExecutionAlreadyStartedFault,
}
