// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package personalizeeventsiface provides an interface to enable mocking the Amazon Personalize Events service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package personalizeeventsiface

import (
	"github.com/vidu171/dell-ecs-s3-client/aws"
	"github.com/vidu171/dell-ecs-s3-client/aws/request"
	"github.com/vidu171/dell-ecs-s3-client/service/personalizeevents"
)

// PersonalizeEventsAPI provides an interface to enable mocking the
// personalizeevents.PersonalizeEvents service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Personalize Events.
//    func myFunc(svc personalizeeventsiface.PersonalizeEventsAPI) bool {
//        // Make svc.PutEvents request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := personalizeevents.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockPersonalizeEventsClient struct {
//        personalizeeventsiface.PersonalizeEventsAPI
//    }
//    func (m *mockPersonalizeEventsClient) PutEvents(input *personalizeevents.PutEventsInput) (*personalizeevents.PutEventsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockPersonalizeEventsClient{}
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
type PersonalizeEventsAPI interface {
	PutEvents(*personalizeevents.PutEventsInput) (*personalizeevents.PutEventsOutput, error)
	PutEventsWithContext(aws.Context, *personalizeevents.PutEventsInput, ...request.Option) (*personalizeevents.PutEventsOutput, error)
	PutEventsRequest(*personalizeevents.PutEventsInput) (*request.Request, *personalizeevents.PutEventsOutput)

	PutItems(*personalizeevents.PutItemsInput) (*personalizeevents.PutItemsOutput, error)
	PutItemsWithContext(aws.Context, *personalizeevents.PutItemsInput, ...request.Option) (*personalizeevents.PutItemsOutput, error)
	PutItemsRequest(*personalizeevents.PutItemsInput) (*request.Request, *personalizeevents.PutItemsOutput)

	PutUsers(*personalizeevents.PutUsersInput) (*personalizeevents.PutUsersOutput, error)
	PutUsersWithContext(aws.Context, *personalizeevents.PutUsersInput, ...request.Option) (*personalizeevents.PutUsersOutput, error)
	PutUsersRequest(*personalizeevents.PutUsersInput) (*request.Request, *personalizeevents.PutUsersOutput)
}

var _ PersonalizeEventsAPI = (*personalizeevents.PersonalizeEvents)(nil)
