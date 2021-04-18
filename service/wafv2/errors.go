// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package wafv2

import (
	"github.com/vidu171/dell-ecs-s3-client/private/protocol"
)

const (

	// ErrCodeWAFAssociatedItemException for service response error code
	// "WAFAssociatedItemException".
	//
	// AWS WAF couldn’t perform the operation because your resource is being used
	// by another resource or it’s associated with another resource.
	ErrCodeWAFAssociatedItemException = "WAFAssociatedItemException"

	// ErrCodeWAFDuplicateItemException for service response error code
	// "WAFDuplicateItemException".
	//
	// AWS WAF couldn’t perform the operation because the resource that you tried
	// to save is a duplicate of an existing one.
	ErrCodeWAFDuplicateItemException = "WAFDuplicateItemException"

	// ErrCodeWAFInternalErrorException for service response error code
	// "WAFInternalErrorException".
	//
	// Your request is valid, but AWS WAF couldn’t perform the operation because
	// of a system problem. Retry your request.
	ErrCodeWAFInternalErrorException = "WAFInternalErrorException"

	// ErrCodeWAFInvalidOperationException for service response error code
	// "WAFInvalidOperationException".
	//
	// The operation isn't valid.
	ErrCodeWAFInvalidOperationException = "WAFInvalidOperationException"

	// ErrCodeWAFInvalidParameterException for service response error code
	// "WAFInvalidParameterException".
	//
	// The operation failed because AWS WAF didn't recognize a parameter in the
	// request. For example:
	//
	//    * You specified an invalid parameter name or value.
	//
	//    * Your nested statement isn't valid. You might have tried to nest a statement
	//    that can’t be nested.
	//
	//    * You tried to update a WebACL with a DefaultAction that isn't among the
	//    types available at DefaultAction.
	//
	//    * Your request references an ARN that is malformed, or corresponds to
	//    a resource with which a Web ACL cannot be associated.
	ErrCodeWAFInvalidParameterException = "WAFInvalidParameterException"

	// ErrCodeWAFInvalidPermissionPolicyException for service response error code
	// "WAFInvalidPermissionPolicyException".
	//
	// The operation failed because the specified policy isn't in the proper format.
	//
	// The policy specifications must conform to the following:
	//
	//    * The policy must be composed using IAM Policy version 2012-10-17 or version
	//    2015-01-01.
	//
	//    * The policy must include specifications for Effect, Action, and Principal.
	//
	//    * Effect must specify Allow.
	//
	//    * Action must specify wafv2:CreateWebACL, wafv2:UpdateWebACL, and wafv2:PutFirewallManagerRuleGroups.
	//    AWS WAF rejects any extra actions or wildcard actions in the policy.
	//
	//    * The policy must not include a Resource parameter.
	//
	// For more information, see IAM Policies (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html).
	ErrCodeWAFInvalidPermissionPolicyException = "WAFInvalidPermissionPolicyException"

	// ErrCodeWAFInvalidResourceException for service response error code
	// "WAFInvalidResourceException".
	//
	// AWS WAF couldn’t perform the operation because the resource that you requested
	// isn’t valid. Check the resource, and try again.
	ErrCodeWAFInvalidResourceException = "WAFInvalidResourceException"

	// ErrCodeWAFLimitsExceededException for service response error code
	// "WAFLimitsExceededException".
	//
	// AWS WAF couldn’t perform the operation because you exceeded your resource
	// limit. For example, the maximum number of WebACL objects that you can create
	// for an AWS account. For more information, see Limits (https://docs.aws.amazon.com/waf/latest/developerguide/limits.html)
	// in the AWS WAF Developer Guide.
	ErrCodeWAFLimitsExceededException = "WAFLimitsExceededException"

	// ErrCodeWAFNonexistentItemException for service response error code
	// "WAFNonexistentItemException".
	//
	// AWS WAF couldn’t perform the operation because your resource doesn’t
	// exist.
	ErrCodeWAFNonexistentItemException = "WAFNonexistentItemException"

	// ErrCodeWAFOptimisticLockException for service response error code
	// "WAFOptimisticLockException".
	//
	// AWS WAF couldn’t save your changes because you tried to update or delete
	// a resource that has changed since you last retrieved it. Get the resource
	// again, make any changes you need to make to the new copy, and retry your
	// operation.
	ErrCodeWAFOptimisticLockException = "WAFOptimisticLockException"

	// ErrCodeWAFServiceLinkedRoleErrorException for service response error code
	// "WAFServiceLinkedRoleErrorException".
	//
	// AWS WAF is not able to access the service linked role. This can be caused
	// by a previous PutLoggingConfiguration request, which can lock the service
	// linked role for about 20 seconds. Please try your request again. The service
	// linked role can also be locked by a previous DeleteServiceLinkedRole request,
	// which can lock the role for 15 minutes or more. If you recently made a call
	// to DeleteServiceLinkedRole, wait at least 15 minutes and try the request
	// again. If you receive this same exception again, you will have to wait additional
	// time until the role is unlocked.
	ErrCodeWAFServiceLinkedRoleErrorException = "WAFServiceLinkedRoleErrorException"

	// ErrCodeWAFSubscriptionNotFoundException for service response error code
	// "WAFSubscriptionNotFoundException".
	//
	// You tried to use a managed rule group that's available by subscription, but
	// you aren't subscribed to it yet.
	ErrCodeWAFSubscriptionNotFoundException = "WAFSubscriptionNotFoundException"

	// ErrCodeWAFTagOperationException for service response error code
	// "WAFTagOperationException".
	//
	// An error occurred during the tagging operation. Retry your request.
	ErrCodeWAFTagOperationException = "WAFTagOperationException"

	// ErrCodeWAFTagOperationInternalErrorException for service response error code
	// "WAFTagOperationInternalErrorException".
	//
	// AWS WAF couldn’t perform your tagging operation because of an internal
	// error. Retry your request.
	ErrCodeWAFTagOperationInternalErrorException = "WAFTagOperationInternalErrorException"

	// ErrCodeWAFUnavailableEntityException for service response error code
	// "WAFUnavailableEntityException".
	//
	// AWS WAF couldn’t retrieve the resource that you requested. Retry your request.
	ErrCodeWAFUnavailableEntityException = "WAFUnavailableEntityException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"WAFAssociatedItemException":            newErrorWAFAssociatedItemException,
	"WAFDuplicateItemException":             newErrorWAFDuplicateItemException,
	"WAFInternalErrorException":             newErrorWAFInternalErrorException,
	"WAFInvalidOperationException":          newErrorWAFInvalidOperationException,
	"WAFInvalidParameterException":          newErrorWAFInvalidParameterException,
	"WAFInvalidPermissionPolicyException":   newErrorWAFInvalidPermissionPolicyException,
	"WAFInvalidResourceException":           newErrorWAFInvalidResourceException,
	"WAFLimitsExceededException":            newErrorWAFLimitsExceededException,
	"WAFNonexistentItemException":           newErrorWAFNonexistentItemException,
	"WAFOptimisticLockException":            newErrorWAFOptimisticLockException,
	"WAFServiceLinkedRoleErrorException":    newErrorWAFServiceLinkedRoleErrorException,
	"WAFSubscriptionNotFoundException":      newErrorWAFSubscriptionNotFoundException,
	"WAFTagOperationException":              newErrorWAFTagOperationException,
	"WAFTagOperationInternalErrorException": newErrorWAFTagOperationInternalErrorException,
	"WAFUnavailableEntityException":         newErrorWAFUnavailableEntityException,
}
