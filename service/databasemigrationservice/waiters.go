// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package databasemigrationservice

import (
	"time"

	"github.com/vidu171/dell-ecs-s3-client/aws"
	"github.com/vidu171/dell-ecs-s3-client/aws/request"
)

// WaitUntilEndpointDeleted uses the AWS Database Migration Service API operation
// DescribeEndpoints to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *DatabaseMigrationService) WaitUntilEndpointDeleted(input *DescribeEndpointsInput) error {
	return c.WaitUntilEndpointDeletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilEndpointDeletedWithContext is an extended version of WaitUntilEndpointDeleted.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DatabaseMigrationService) WaitUntilEndpointDeletedWithContext(ctx aws.Context, input *DescribeEndpointsInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilEndpointDeleted",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "ResourceNotFoundFault",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Endpoints[].Status",
				Expected: "active",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Endpoints[].Status",
				Expected: "creating",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeEndpointsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeEndpointsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilReplicationInstanceAvailable uses the AWS Database Migration Service API operation
// DescribeReplicationInstances to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *DatabaseMigrationService) WaitUntilReplicationInstanceAvailable(input *DescribeReplicationInstancesInput) error {
	return c.WaitUntilReplicationInstanceAvailableWithContext(aws.BackgroundContext(), input)
}

// WaitUntilReplicationInstanceAvailableWithContext is an extended version of WaitUntilReplicationInstanceAvailable.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DatabaseMigrationService) WaitUntilReplicationInstanceAvailableWithContext(ctx aws.Context, input *DescribeReplicationInstancesInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilReplicationInstanceAvailable",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(60 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "ReplicationInstances[].ReplicationInstanceStatus",
				Expected: "available",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationInstances[].ReplicationInstanceStatus",
				Expected: "deleting",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationInstances[].ReplicationInstanceStatus",
				Expected: "incompatible-credentials",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationInstances[].ReplicationInstanceStatus",
				Expected: "incompatible-network",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationInstances[].ReplicationInstanceStatus",
				Expected: "inaccessible-encryption-credentials",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeReplicationInstancesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeReplicationInstancesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilReplicationInstanceDeleted uses the AWS Database Migration Service API operation
// DescribeReplicationInstances to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *DatabaseMigrationService) WaitUntilReplicationInstanceDeleted(input *DescribeReplicationInstancesInput) error {
	return c.WaitUntilReplicationInstanceDeletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilReplicationInstanceDeletedWithContext is an extended version of WaitUntilReplicationInstanceDeleted.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DatabaseMigrationService) WaitUntilReplicationInstanceDeletedWithContext(ctx aws.Context, input *DescribeReplicationInstancesInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilReplicationInstanceDeleted",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationInstances[].ReplicationInstanceStatus",
				Expected: "available",
			},
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "ResourceNotFoundFault",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeReplicationInstancesInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeReplicationInstancesRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilReplicationTaskDeleted uses the AWS Database Migration Service API operation
// DescribeReplicationTasks to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *DatabaseMigrationService) WaitUntilReplicationTaskDeleted(input *DescribeReplicationTasksInput) error {
	return c.WaitUntilReplicationTaskDeletedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilReplicationTaskDeletedWithContext is an extended version of WaitUntilReplicationTaskDeleted.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DatabaseMigrationService) WaitUntilReplicationTaskDeletedWithContext(ctx aws.Context, input *DescribeReplicationTasksInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilReplicationTaskDeleted",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "ready",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "creating",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "stopped",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "running",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "failed",
			},
			{
				State:    request.SuccessWaiterState,
				Matcher:  request.ErrorWaiterMatch,
				Expected: "ResourceNotFoundFault",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeReplicationTasksInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeReplicationTasksRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilReplicationTaskReady uses the AWS Database Migration Service API operation
// DescribeReplicationTasks to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *DatabaseMigrationService) WaitUntilReplicationTaskReady(input *DescribeReplicationTasksInput) error {
	return c.WaitUntilReplicationTaskReadyWithContext(aws.BackgroundContext(), input)
}

// WaitUntilReplicationTaskReadyWithContext is an extended version of WaitUntilReplicationTaskReady.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DatabaseMigrationService) WaitUntilReplicationTaskReadyWithContext(ctx aws.Context, input *DescribeReplicationTasksInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilReplicationTaskReady",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "ready",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "starting",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "running",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "stopping",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "stopped",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "failed",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "modifying",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "testing",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "deleting",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeReplicationTasksInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeReplicationTasksRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilReplicationTaskRunning uses the AWS Database Migration Service API operation
// DescribeReplicationTasks to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *DatabaseMigrationService) WaitUntilReplicationTaskRunning(input *DescribeReplicationTasksInput) error {
	return c.WaitUntilReplicationTaskRunningWithContext(aws.BackgroundContext(), input)
}

// WaitUntilReplicationTaskRunningWithContext is an extended version of WaitUntilReplicationTaskRunning.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DatabaseMigrationService) WaitUntilReplicationTaskRunningWithContext(ctx aws.Context, input *DescribeReplicationTasksInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilReplicationTaskRunning",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "running",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "ready",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "creating",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "stopping",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "stopped",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "failed",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "modifying",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "testing",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "deleting",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeReplicationTasksInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeReplicationTasksRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilReplicationTaskStopped uses the AWS Database Migration Service API operation
// DescribeReplicationTasks to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *DatabaseMigrationService) WaitUntilReplicationTaskStopped(input *DescribeReplicationTasksInput) error {
	return c.WaitUntilReplicationTaskStoppedWithContext(aws.BackgroundContext(), input)
}

// WaitUntilReplicationTaskStoppedWithContext is an extended version of WaitUntilReplicationTaskStopped.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DatabaseMigrationService) WaitUntilReplicationTaskStoppedWithContext(ctx aws.Context, input *DescribeReplicationTasksInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilReplicationTaskStopped",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(15 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "stopped",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "ready",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "creating",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "starting",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "failed",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "modifying",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "testing",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "ReplicationTasks[].Status",
				Expected: "deleting",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeReplicationTasksInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeReplicationTasksRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}

// WaitUntilTestConnectionSucceeds uses the AWS Database Migration Service API operation
// DescribeConnections to wait for a condition to be met before returning.
// If the condition is not met within the max attempt window, an error will
// be returned.
func (c *DatabaseMigrationService) WaitUntilTestConnectionSucceeds(input *DescribeConnectionsInput) error {
	return c.WaitUntilTestConnectionSucceedsWithContext(aws.BackgroundContext(), input)
}

// WaitUntilTestConnectionSucceedsWithContext is an extended version of WaitUntilTestConnectionSucceeds.
// With the support for passing in a context and options to configure the
// Waiter and the underlying request options.
//
// The context must be non-nil and will be used for request cancellation. If
// the context is nil a panic will occur. In the future the SDK may create
// sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *DatabaseMigrationService) WaitUntilTestConnectionSucceedsWithContext(ctx aws.Context, input *DescribeConnectionsInput, opts ...request.WaiterOption) error {
	w := request.Waiter{
		Name:        "WaitUntilTestConnectionSucceeds",
		MaxAttempts: 60,
		Delay:       request.ConstantWaiterDelay(5 * time.Second),
		Acceptors: []request.WaiterAcceptor{
			{
				State:   request.SuccessWaiterState,
				Matcher: request.PathAllWaiterMatch, Argument: "Connections[].Status",
				Expected: "successful",
			},
			{
				State:   request.FailureWaiterState,
				Matcher: request.PathAnyWaiterMatch, Argument: "Connections[].Status",
				Expected: "failed",
			},
		},
		Logger: c.Config.Logger,
		NewRequest: func(opts []request.Option) (*request.Request, error) {
			var inCpy *DescribeConnectionsInput
			if input != nil {
				tmp := *input
				inCpy = &tmp
			}
			req, _ := c.DescribeConnectionsRequest(inCpy)
			req.SetContext(ctx)
			req.ApplyOptions(opts...)
			return req, nil
		},
	}
	w.ApplyOptions(opts...)

	return w.WaitWithContext(ctx)
}
