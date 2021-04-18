// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package applicationautoscalingiface provides an interface to enable mocking the Application Auto Scaling service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package applicationautoscalingiface

import (
	"github.com/vidu171/dell-ecs-s3-client/aws"
	"github.com/vidu171/dell-ecs-s3-client/aws/request"
	"github.com/vidu171/dell-ecs-s3-client/service/applicationautoscaling"
)

// ApplicationAutoScalingAPI provides an interface to enable mocking the
// applicationautoscaling.ApplicationAutoScaling service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Application Auto Scaling.
//    func myFunc(svc applicationautoscalingiface.ApplicationAutoScalingAPI) bool {
//        // Make svc.DeleteScalingPolicy request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := applicationautoscaling.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockApplicationAutoScalingClient struct {
//        applicationautoscalingiface.ApplicationAutoScalingAPI
//    }
//    func (m *mockApplicationAutoScalingClient) DeleteScalingPolicy(input *applicationautoscaling.DeleteScalingPolicyInput) (*applicationautoscaling.DeleteScalingPolicyOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockApplicationAutoScalingClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type ApplicationAutoScalingAPI interface {
	DeleteScalingPolicy(*applicationautoscaling.DeleteScalingPolicyInput) (*applicationautoscaling.DeleteScalingPolicyOutput, error)
	DeleteScalingPolicyWithContext(aws.Context, *applicationautoscaling.DeleteScalingPolicyInput, ...request.Option) (*applicationautoscaling.DeleteScalingPolicyOutput, error)
	DeleteScalingPolicyRequest(*applicationautoscaling.DeleteScalingPolicyInput) (*request.Request, *applicationautoscaling.DeleteScalingPolicyOutput)

	DeleteScheduledAction(*applicationautoscaling.DeleteScheduledActionInput) (*applicationautoscaling.DeleteScheduledActionOutput, error)
	DeleteScheduledActionWithContext(aws.Context, *applicationautoscaling.DeleteScheduledActionInput, ...request.Option) (*applicationautoscaling.DeleteScheduledActionOutput, error)
	DeleteScheduledActionRequest(*applicationautoscaling.DeleteScheduledActionInput) (*request.Request, *applicationautoscaling.DeleteScheduledActionOutput)

	DeregisterScalableTarget(*applicationautoscaling.DeregisterScalableTargetInput) (*applicationautoscaling.DeregisterScalableTargetOutput, error)
	DeregisterScalableTargetWithContext(aws.Context, *applicationautoscaling.DeregisterScalableTargetInput, ...request.Option) (*applicationautoscaling.DeregisterScalableTargetOutput, error)
	DeregisterScalableTargetRequest(*applicationautoscaling.DeregisterScalableTargetInput) (*request.Request, *applicationautoscaling.DeregisterScalableTargetOutput)

	DescribeScalableTargets(*applicationautoscaling.DescribeScalableTargetsInput) (*applicationautoscaling.DescribeScalableTargetsOutput, error)
	DescribeScalableTargetsWithContext(aws.Context, *applicationautoscaling.DescribeScalableTargetsInput, ...request.Option) (*applicationautoscaling.DescribeScalableTargetsOutput, error)
	DescribeScalableTargetsRequest(*applicationautoscaling.DescribeScalableTargetsInput) (*request.Request, *applicationautoscaling.DescribeScalableTargetsOutput)

	DescribeScalableTargetsPages(*applicationautoscaling.DescribeScalableTargetsInput, func(*applicationautoscaling.DescribeScalableTargetsOutput, bool) bool) error
	DescribeScalableTargetsPagesWithContext(aws.Context, *applicationautoscaling.DescribeScalableTargetsInput, func(*applicationautoscaling.DescribeScalableTargetsOutput, bool) bool, ...request.Option) error

	DescribeScalingActivities(*applicationautoscaling.DescribeScalingActivitiesInput) (*applicationautoscaling.DescribeScalingActivitiesOutput, error)
	DescribeScalingActivitiesWithContext(aws.Context, *applicationautoscaling.DescribeScalingActivitiesInput, ...request.Option) (*applicationautoscaling.DescribeScalingActivitiesOutput, error)
	DescribeScalingActivitiesRequest(*applicationautoscaling.DescribeScalingActivitiesInput) (*request.Request, *applicationautoscaling.DescribeScalingActivitiesOutput)

	DescribeScalingActivitiesPages(*applicationautoscaling.DescribeScalingActivitiesInput, func(*applicationautoscaling.DescribeScalingActivitiesOutput, bool) bool) error
	DescribeScalingActivitiesPagesWithContext(aws.Context, *applicationautoscaling.DescribeScalingActivitiesInput, func(*applicationautoscaling.DescribeScalingActivitiesOutput, bool) bool, ...request.Option) error

	DescribeScalingPolicies(*applicationautoscaling.DescribeScalingPoliciesInput) (*applicationautoscaling.DescribeScalingPoliciesOutput, error)
	DescribeScalingPoliciesWithContext(aws.Context, *applicationautoscaling.DescribeScalingPoliciesInput, ...request.Option) (*applicationautoscaling.DescribeScalingPoliciesOutput, error)
	DescribeScalingPoliciesRequest(*applicationautoscaling.DescribeScalingPoliciesInput) (*request.Request, *applicationautoscaling.DescribeScalingPoliciesOutput)

	DescribeScalingPoliciesPages(*applicationautoscaling.DescribeScalingPoliciesInput, func(*applicationautoscaling.DescribeScalingPoliciesOutput, bool) bool) error
	DescribeScalingPoliciesPagesWithContext(aws.Context, *applicationautoscaling.DescribeScalingPoliciesInput, func(*applicationautoscaling.DescribeScalingPoliciesOutput, bool) bool, ...request.Option) error

	DescribeScheduledActions(*applicationautoscaling.DescribeScheduledActionsInput) (*applicationautoscaling.DescribeScheduledActionsOutput, error)
	DescribeScheduledActionsWithContext(aws.Context, *applicationautoscaling.DescribeScheduledActionsInput, ...request.Option) (*applicationautoscaling.DescribeScheduledActionsOutput, error)
	DescribeScheduledActionsRequest(*applicationautoscaling.DescribeScheduledActionsInput) (*request.Request, *applicationautoscaling.DescribeScheduledActionsOutput)

	DescribeScheduledActionsPages(*applicationautoscaling.DescribeScheduledActionsInput, func(*applicationautoscaling.DescribeScheduledActionsOutput, bool) bool) error
	DescribeScheduledActionsPagesWithContext(aws.Context, *applicationautoscaling.DescribeScheduledActionsInput, func(*applicationautoscaling.DescribeScheduledActionsOutput, bool) bool, ...request.Option) error

	PutScalingPolicy(*applicationautoscaling.PutScalingPolicyInput) (*applicationautoscaling.PutScalingPolicyOutput, error)
	PutScalingPolicyWithContext(aws.Context, *applicationautoscaling.PutScalingPolicyInput, ...request.Option) (*applicationautoscaling.PutScalingPolicyOutput, error)
	PutScalingPolicyRequest(*applicationautoscaling.PutScalingPolicyInput) (*request.Request, *applicationautoscaling.PutScalingPolicyOutput)

	PutScheduledAction(*applicationautoscaling.PutScheduledActionInput) (*applicationautoscaling.PutScheduledActionOutput, error)
	PutScheduledActionWithContext(aws.Context, *applicationautoscaling.PutScheduledActionInput, ...request.Option) (*applicationautoscaling.PutScheduledActionOutput, error)
	PutScheduledActionRequest(*applicationautoscaling.PutScheduledActionInput) (*request.Request, *applicationautoscaling.PutScheduledActionOutput)

	RegisterScalableTarget(*applicationautoscaling.RegisterScalableTargetInput) (*applicationautoscaling.RegisterScalableTargetOutput, error)
	RegisterScalableTargetWithContext(aws.Context, *applicationautoscaling.RegisterScalableTargetInput, ...request.Option) (*applicationautoscaling.RegisterScalableTargetOutput, error)
	RegisterScalableTargetRequest(*applicationautoscaling.RegisterScalableTargetInput) (*request.Request, *applicationautoscaling.RegisterScalableTargetOutput)
}

var _ ApplicationAutoScalingAPI = (*applicationautoscaling.ApplicationAutoScaling)(nil)
