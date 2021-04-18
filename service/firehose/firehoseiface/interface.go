// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package firehoseiface provides an interface to enable mocking the Amazon Kinesis Firehose service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package firehoseiface

import (
	"github.com/vidu171/dell-ecs-s3-client/aws"
	"github.com/vidu171/dell-ecs-s3-client/aws/request"
	"github.com/vidu171/dell-ecs-s3-client/service/firehose"
)

// FirehoseAPI provides an interface to enable mocking the
// firehose.Firehose service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Kinesis Firehose.
//    func myFunc(svc firehoseiface.FirehoseAPI) bool {
//        // Make svc.CreateDeliveryStream request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := firehose.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockFirehoseClient struct {
//        firehoseiface.FirehoseAPI
//    }
//    func (m *mockFirehoseClient) CreateDeliveryStream(input *firehose.CreateDeliveryStreamInput) (*firehose.CreateDeliveryStreamOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockFirehoseClient{}
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
type FirehoseAPI interface {
	CreateDeliveryStream(*firehose.CreateDeliveryStreamInput) (*firehose.CreateDeliveryStreamOutput, error)
	CreateDeliveryStreamWithContext(aws.Context, *firehose.CreateDeliveryStreamInput, ...request.Option) (*firehose.CreateDeliveryStreamOutput, error)
	CreateDeliveryStreamRequest(*firehose.CreateDeliveryStreamInput) (*request.Request, *firehose.CreateDeliveryStreamOutput)

	DeleteDeliveryStream(*firehose.DeleteDeliveryStreamInput) (*firehose.DeleteDeliveryStreamOutput, error)
	DeleteDeliveryStreamWithContext(aws.Context, *firehose.DeleteDeliveryStreamInput, ...request.Option) (*firehose.DeleteDeliveryStreamOutput, error)
	DeleteDeliveryStreamRequest(*firehose.DeleteDeliveryStreamInput) (*request.Request, *firehose.DeleteDeliveryStreamOutput)

	DescribeDeliveryStream(*firehose.DescribeDeliveryStreamInput) (*firehose.DescribeDeliveryStreamOutput, error)
	DescribeDeliveryStreamWithContext(aws.Context, *firehose.DescribeDeliveryStreamInput, ...request.Option) (*firehose.DescribeDeliveryStreamOutput, error)
	DescribeDeliveryStreamRequest(*firehose.DescribeDeliveryStreamInput) (*request.Request, *firehose.DescribeDeliveryStreamOutput)

	ListDeliveryStreams(*firehose.ListDeliveryStreamsInput) (*firehose.ListDeliveryStreamsOutput, error)
	ListDeliveryStreamsWithContext(aws.Context, *firehose.ListDeliveryStreamsInput, ...request.Option) (*firehose.ListDeliveryStreamsOutput, error)
	ListDeliveryStreamsRequest(*firehose.ListDeliveryStreamsInput) (*request.Request, *firehose.ListDeliveryStreamsOutput)

	ListTagsForDeliveryStream(*firehose.ListTagsForDeliveryStreamInput) (*firehose.ListTagsForDeliveryStreamOutput, error)
	ListTagsForDeliveryStreamWithContext(aws.Context, *firehose.ListTagsForDeliveryStreamInput, ...request.Option) (*firehose.ListTagsForDeliveryStreamOutput, error)
	ListTagsForDeliveryStreamRequest(*firehose.ListTagsForDeliveryStreamInput) (*request.Request, *firehose.ListTagsForDeliveryStreamOutput)

	PutRecord(*firehose.PutRecordInput) (*firehose.PutRecordOutput, error)
	PutRecordWithContext(aws.Context, *firehose.PutRecordInput, ...request.Option) (*firehose.PutRecordOutput, error)
	PutRecordRequest(*firehose.PutRecordInput) (*request.Request, *firehose.PutRecordOutput)

	PutRecordBatch(*firehose.PutRecordBatchInput) (*firehose.PutRecordBatchOutput, error)
	PutRecordBatchWithContext(aws.Context, *firehose.PutRecordBatchInput, ...request.Option) (*firehose.PutRecordBatchOutput, error)
	PutRecordBatchRequest(*firehose.PutRecordBatchInput) (*request.Request, *firehose.PutRecordBatchOutput)

	StartDeliveryStreamEncryption(*firehose.StartDeliveryStreamEncryptionInput) (*firehose.StartDeliveryStreamEncryptionOutput, error)
	StartDeliveryStreamEncryptionWithContext(aws.Context, *firehose.StartDeliveryStreamEncryptionInput, ...request.Option) (*firehose.StartDeliveryStreamEncryptionOutput, error)
	StartDeliveryStreamEncryptionRequest(*firehose.StartDeliveryStreamEncryptionInput) (*request.Request, *firehose.StartDeliveryStreamEncryptionOutput)

	StopDeliveryStreamEncryption(*firehose.StopDeliveryStreamEncryptionInput) (*firehose.StopDeliveryStreamEncryptionOutput, error)
	StopDeliveryStreamEncryptionWithContext(aws.Context, *firehose.StopDeliveryStreamEncryptionInput, ...request.Option) (*firehose.StopDeliveryStreamEncryptionOutput, error)
	StopDeliveryStreamEncryptionRequest(*firehose.StopDeliveryStreamEncryptionInput) (*request.Request, *firehose.StopDeliveryStreamEncryptionOutput)

	TagDeliveryStream(*firehose.TagDeliveryStreamInput) (*firehose.TagDeliveryStreamOutput, error)
	TagDeliveryStreamWithContext(aws.Context, *firehose.TagDeliveryStreamInput, ...request.Option) (*firehose.TagDeliveryStreamOutput, error)
	TagDeliveryStreamRequest(*firehose.TagDeliveryStreamInput) (*request.Request, *firehose.TagDeliveryStreamOutput)

	UntagDeliveryStream(*firehose.UntagDeliveryStreamInput) (*firehose.UntagDeliveryStreamOutput, error)
	UntagDeliveryStreamWithContext(aws.Context, *firehose.UntagDeliveryStreamInput, ...request.Option) (*firehose.UntagDeliveryStreamOutput, error)
	UntagDeliveryStreamRequest(*firehose.UntagDeliveryStreamInput) (*request.Request, *firehose.UntagDeliveryStreamOutput)

	UpdateDestination(*firehose.UpdateDestinationInput) (*firehose.UpdateDestinationOutput, error)
	UpdateDestinationWithContext(aws.Context, *firehose.UpdateDestinationInput, ...request.Option) (*firehose.UpdateDestinationOutput, error)
	UpdateDestinationRequest(*firehose.UpdateDestinationInput) (*request.Request, *firehose.UpdateDestinationOutput)
}

var _ FirehoseAPI = (*firehose.Firehose)(nil)
