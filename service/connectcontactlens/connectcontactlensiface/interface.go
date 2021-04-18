// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package connectcontactlensiface provides an interface to enable mocking the Amazon Connect Contact Lens service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package connectcontactlensiface

import (
	"github.com/vidu171/dell-ecs-s3-client/aws"
	"github.com/vidu171/dell-ecs-s3-client/aws/request"
	"github.com/vidu171/dell-ecs-s3-client/service/connectcontactlens"
)

// ConnectContactLensAPI provides an interface to enable mocking the
// connectcontactlens.ConnectContactLens service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Connect Contact Lens.
//    func myFunc(svc connectcontactlensiface.ConnectContactLensAPI) bool {
//        // Make svc.ListRealtimeContactAnalysisSegments request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := connectcontactlens.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockConnectContactLensClient struct {
//        connectcontactlensiface.ConnectContactLensAPI
//    }
//    func (m *mockConnectContactLensClient) ListRealtimeContactAnalysisSegments(input *connectcontactlens.ListRealtimeContactAnalysisSegmentsInput) (*connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockConnectContactLensClient{}
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
type ConnectContactLensAPI interface {
	ListRealtimeContactAnalysisSegments(*connectcontactlens.ListRealtimeContactAnalysisSegmentsInput) (*connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput, error)
	ListRealtimeContactAnalysisSegmentsWithContext(aws.Context, *connectcontactlens.ListRealtimeContactAnalysisSegmentsInput, ...request.Option) (*connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput, error)
	ListRealtimeContactAnalysisSegmentsRequest(*connectcontactlens.ListRealtimeContactAnalysisSegmentsInput) (*request.Request, *connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput)

	ListRealtimeContactAnalysisSegmentsPages(*connectcontactlens.ListRealtimeContactAnalysisSegmentsInput, func(*connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput, bool) bool) error
	ListRealtimeContactAnalysisSegmentsPagesWithContext(aws.Context, *connectcontactlens.ListRealtimeContactAnalysisSegmentsInput, func(*connectcontactlens.ListRealtimeContactAnalysisSegmentsOutput, bool) bool, ...request.Option) error
}

var _ ConnectContactLensAPI = (*connectcontactlens.ConnectContactLens)(nil)
