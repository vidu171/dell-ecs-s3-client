// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package schemas

import (
	"time"

	"github.com/vidu171/dell-ecs-s3-client/aws"
	"github.com/vidu171/dell-ecs-s3-client/aws/request"
)

// WaitUntilCodeBindingExists uses the Schemas API operation
// DescribeCodeBinding to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *Schemas) WaitUntilCodeBindingExists(input *DescribeCodeBindingInput) error {
	return c.WaitUntilCodeBindingExistsWithContext(aws.BackgroundContext(), input)
}

// WaitUntilCodeBindingExistsWithContext is an extended version of WaitUntilCodeBindingExists.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *Schemas) WaitUntilCodeBindingExistsWithContext(ctx aws.Context, input *DescribeCodeBindingInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilCodeBindingExists",
		MaxAttempts: 30,
		Delay:       request.ConstantWaiterDelay(2 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "Status",
				Expected: "CREATE_COMPLETE",
			},
			{
				State:   request.RetryWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "Status",
				Expected: "CREATE_IN_PROGRESS",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathWaiterMatch, Argument: "Status",
				Expected: "CREATE_FAILED",
			},
			{
				State:    request.FailureWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "NotFoundException",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeCodeBindingInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeCodeBindingRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}