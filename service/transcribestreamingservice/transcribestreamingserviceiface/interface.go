// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package transcribestreamingserviceiface provides an interface to enable mocking the Amazon Transcribe Streaming Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package transcribestreamingserviceiface

import (
	"github.com/vidu171/dell-ecs-s3-client/aws"
	"github.com/vidu171/dell-ecs-s3-client/aws/request"
	"github.com/vidu171/dell-ecs-s3-client/service/transcribestreamingservice"
)

// TranscribeStreamingServiceAPI provides an interface to enable mocking the
// transcribestreamingservice.TranscribeStreamingService service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Transcribe Streaming Service.
//    func myFunc(svc transcribestreamingserviceiface.TranscribeStreamingServiceAPI) bool {
//        // Make svc.StartMedicalStreamTranscription request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := transcribestreamingservice.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockTranscribeStreamingServiceClient struct {
//        transcribestreamingserviceiface.TranscribeStreamingServiceAPI
//    }
//    func (m *mockTranscribeStreamingServiceClient) StartMedicalStreamTranscription(input *transcribestreamingservice.StartMedicalStreamTranscriptionInput) (*transcribestreamingservice.StartMedicalStreamTranscriptionOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockTranscribeStreamingServiceClient{}
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
type TranscribeStreamingServiceAPI interface {
	StartMedicalStreamTranscription(*transcribestreamingservice.StartMedicalStreamTranscriptionInput) (*transcribestreamingservice.StartMedicalStreamTranscriptionOutput, error)
	StartMedicalStreamTranscriptionWithContext(aws.Context, *transcribestreamingservice.StartMedicalStreamTranscriptionInput, ...request.Option) (*transcribestreamingservice.StartMedicalStreamTranscriptionOutput, error)
	StartMedicalStreamTranscriptionRequest(*transcribestreamingservice.StartMedicalStreamTranscriptionInput) (*request.Request, *transcribestreamingservice.StartMedicalStreamTranscriptionOutput)

	StartStreamTranscription(*transcribestreamingservice.StartStreamTranscriptionInput) (*transcribestreamingservice.StartStreamTranscriptionOutput, error)
	StartStreamTranscriptionWithContext(aws.Context, *transcribestreamingservice.StartStreamTranscriptionInput, ...request.Option) (*transcribestreamingservice.StartStreamTranscriptionOutput, error)
	StartStreamTranscriptionRequest(*transcribestreamingservice.StartStreamTranscriptionInput) (*request.Request, *transcribestreamingservice.StartStreamTranscriptionOutput)
}

var _ TranscribeStreamingServiceAPI = (*transcribestreamingservice.TranscribeStreamingService)(nil)
