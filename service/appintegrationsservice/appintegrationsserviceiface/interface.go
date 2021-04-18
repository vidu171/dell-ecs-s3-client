// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package appintegrationsserviceiface provides an interface to enable mocking the Amazon AppIntegrations Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package appintegrationsserviceiface

import (
	"github.com/vidu171/dell-ecs-s3-client/aws"
	"github.com/vidu171/dell-ecs-s3-client/aws/request"
	"github.com/vidu171/dell-ecs-s3-client/service/appintegrationsservice"
)

// AppIntegrationsServiceAPI provides an interface to enable mocking the
// appintegrationsservice.AppIntegrationsService service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon AppIntegrations Service.
//    func myFunc(svc appintegrationsserviceiface.AppIntegrationsServiceAPI) bool {
//        // Make svc.CreateEventIntegration request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := appintegrationsservice.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockAppIntegrationsServiceClient struct {
//        appintegrationsserviceiface.AppIntegrationsServiceAPI
//    }
//    func (m *mockAppIntegrationsServiceClient) CreateEventIntegration(input *appintegrationsservice.CreateEventIntegrationInput) (*appintegrationsservice.CreateEventIntegrationOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockAppIntegrationsServiceClient{}
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
type AppIntegrationsServiceAPI interface {
	CreateEventIntegration(*appintegrationsservice.CreateEventIntegrationInput) (*appintegrationsservice.CreateEventIntegrationOutput, error)
	CreateEventIntegrationWithContext(aws.Context, *appintegrationsservice.CreateEventIntegrationInput, ...request.Option) (*appintegrationsservice.CreateEventIntegrationOutput, error)
	CreateEventIntegrationRequest(*appintegrationsservice.CreateEventIntegrationInput) (*request.Request, *appintegrationsservice.CreateEventIntegrationOutput)

	DeleteEventIntegration(*appintegrationsservice.DeleteEventIntegrationInput) (*appintegrationsservice.DeleteEventIntegrationOutput, error)
	DeleteEventIntegrationWithContext(aws.Context, *appintegrationsservice.DeleteEventIntegrationInput, ...request.Option) (*appintegrationsservice.DeleteEventIntegrationOutput, error)
	DeleteEventIntegrationRequest(*appintegrationsservice.DeleteEventIntegrationInput) (*request.Request, *appintegrationsservice.DeleteEventIntegrationOutput)

	GetEventIntegration(*appintegrationsservice.GetEventIntegrationInput) (*appintegrationsservice.GetEventIntegrationOutput, error)
	GetEventIntegrationWithContext(aws.Context, *appintegrationsservice.GetEventIntegrationInput, ...request.Option) (*appintegrationsservice.GetEventIntegrationOutput, error)
	GetEventIntegrationRequest(*appintegrationsservice.GetEventIntegrationInput) (*request.Request, *appintegrationsservice.GetEventIntegrationOutput)

	ListEventIntegrationAssociations(*appintegrationsservice.ListEventIntegrationAssociationsInput) (*appintegrationsservice.ListEventIntegrationAssociationsOutput, error)
	ListEventIntegrationAssociationsWithContext(aws.Context, *appintegrationsservice.ListEventIntegrationAssociationsInput, ...request.Option) (*appintegrationsservice.ListEventIntegrationAssociationsOutput, error)
	ListEventIntegrationAssociationsRequest(*appintegrationsservice.ListEventIntegrationAssociationsInput) (*request.Request, *appintegrationsservice.ListEventIntegrationAssociationsOutput)

	ListEventIntegrations(*appintegrationsservice.ListEventIntegrationsInput) (*appintegrationsservice.ListEventIntegrationsOutput, error)
	ListEventIntegrationsWithContext(aws.Context, *appintegrationsservice.ListEventIntegrationsInput, ...request.Option) (*appintegrationsservice.ListEventIntegrationsOutput, error)
	ListEventIntegrationsRequest(*appintegrationsservice.ListEventIntegrationsInput) (*request.Request, *appintegrationsservice.ListEventIntegrationsOutput)

	ListTagsForResource(*appintegrationsservice.ListTagsForResourceInput) (*appintegrationsservice.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *appintegrationsservice.ListTagsForResourceInput, ...request.Option) (*appintegrationsservice.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*appintegrationsservice.ListTagsForResourceInput) (*request.Request, *appintegrationsservice.ListTagsForResourceOutput)

	TagResource(*appintegrationsservice.TagResourceInput) (*appintegrationsservice.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *appintegrationsservice.TagResourceInput, ...request.Option) (*appintegrationsservice.TagResourceOutput, error)
	TagResourceRequest(*appintegrationsservice.TagResourceInput) (*request.Request, *appintegrationsservice.TagResourceOutput)

	UntagResource(*appintegrationsservice.UntagResourceInput) (*appintegrationsservice.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *appintegrationsservice.UntagResourceInput, ...request.Option) (*appintegrationsservice.UntagResourceOutput, error)
	UntagResourceRequest(*appintegrationsservice.UntagResourceInput) (*request.Request, *appintegrationsservice.UntagResourceOutput)

	UpdateEventIntegration(*appintegrationsservice.UpdateEventIntegrationInput) (*appintegrationsservice.UpdateEventIntegrationOutput, error)
	UpdateEventIntegrationWithContext(aws.Context, *appintegrationsservice.UpdateEventIntegrationInput, ...request.Option) (*appintegrationsservice.UpdateEventIntegrationOutput, error)
	UpdateEventIntegrationRequest(*appintegrationsservice.UpdateEventIntegrationInput) (*request.Request, *appintegrationsservice.UpdateEventIntegrationOutput)
}

var _ AppIntegrationsServiceAPI = (*appintegrationsservice.AppIntegrationsService)(nil)
