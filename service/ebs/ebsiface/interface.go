// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package ebsiface provides an interface to enable mocking the Amazon Elastic Block Store service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package ebsiface

import (
	"github.com/vidu171/dell-ecs-s3-client/aws"
	"github.com/vidu171/dell-ecs-s3-client/aws/request"
	"github.com/vidu171/dell-ecs-s3-client/service/ebs"
)

// EBSAPI provides an interface to enable mocking the
// ebs.EBS service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Elastic Block Store.
//    func myFunc(svc ebsiface.EBSAPI) bool {
//        // Make svc.CompleteSnapshot request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := ebs.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockEBSClient struct {
//        ebsiface.EBSAPI
//    }
//    func (m *mockEBSClient) CompleteSnapshot(input *ebs.CompleteSnapshotInput) (*ebs.CompleteSnapshotOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockEBSClient{}
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
type EBSAPI interface {
	CompleteSnapshot(*ebs.CompleteSnapshotInput) (*ebs.CompleteSnapshotOutput, error)
	CompleteSnapshotWithContext(aws.Context, *ebs.CompleteSnapshotInput, ...request.Option) (*ebs.CompleteSnapshotOutput, error)
	CompleteSnapshotRequest(*ebs.CompleteSnapshotInput) (*request.Request, *ebs.CompleteSnapshotOutput)

	GetSnapshotBlock(*ebs.GetSnapshotBlockInput) (*ebs.GetSnapshotBlockOutput, error)
	GetSnapshotBlockWithContext(aws.Context, *ebs.GetSnapshotBlockInput, ...request.Option) (*ebs.GetSnapshotBlockOutput, error)
	GetSnapshotBlockRequest(*ebs.GetSnapshotBlockInput) (*request.Request, *ebs.GetSnapshotBlockOutput)

	ListChangedBlocks(*ebs.ListChangedBlocksInput) (*ebs.ListChangedBlocksOutput, error)
	ListChangedBlocksWithContext(aws.Context, *ebs.ListChangedBlocksInput, ...request.Option) (*ebs.ListChangedBlocksOutput, error)
	ListChangedBlocksRequest(*ebs.ListChangedBlocksInput) (*request.Request, *ebs.ListChangedBlocksOutput)

	ListChangedBlocksPages(*ebs.ListChangedBlocksInput, func(*ebs.ListChangedBlocksOutput, bool) bool) error
	ListChangedBlocksPagesWithContext(aws.Context, *ebs.ListChangedBlocksInput, func(*ebs.ListChangedBlocksOutput, bool) bool, ...request.Option) error

	ListSnapshotBlocks(*ebs.ListSnapshotBlocksInput) (*ebs.ListSnapshotBlocksOutput, error)
	ListSnapshotBlocksWithContext(aws.Context, *ebs.ListSnapshotBlocksInput, ...request.Option) (*ebs.ListSnapshotBlocksOutput, error)
	ListSnapshotBlocksRequest(*ebs.ListSnapshotBlocksInput) (*request.Request, *ebs.ListSnapshotBlocksOutput)

	ListSnapshotBlocksPages(*ebs.ListSnapshotBlocksInput, func(*ebs.ListSnapshotBlocksOutput, bool) bool) error
	ListSnapshotBlocksPagesWithContext(aws.Context, *ebs.ListSnapshotBlocksInput, func(*ebs.ListSnapshotBlocksOutput, bool) bool, ...request.Option) error

	PutSnapshotBlock(*ebs.PutSnapshotBlockInput) (*ebs.PutSnapshotBlockOutput, error)
	PutSnapshotBlockWithContext(aws.Context, *ebs.PutSnapshotBlockInput, ...request.Option) (*ebs.PutSnapshotBlockOutput, error)
	PutSnapshotBlockRequest(*ebs.PutSnapshotBlockInput) (*request.Request, *ebs.PutSnapshotBlockOutput)

	StartSnapshot(*ebs.StartSnapshotInput) (*ebs.StartSnapshotOutput, error)
	StartSnapshotWithContext(aws.Context, *ebs.StartSnapshotInput, ...request.Option) (*ebs.StartSnapshotOutput, error)
	StartSnapshotRequest(*ebs.StartSnapshotInput) (*request.Request, *ebs.StartSnapshotOutput)
}

var _ EBSAPI = (*ebs.EBS)(nil)
