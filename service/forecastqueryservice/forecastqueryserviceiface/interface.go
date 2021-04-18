// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package forecastqueryserviceiface provides an interface to enable mocking the Amazon Forecast Query Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package forecastqueryserviceiface

import (
	"github.com/vidu171/dell-ecs-s3-client/aws"
	"github.com/vidu171/dell-ecs-s3-client/aws/request"
	"github.com/vidu171/dell-ecs-s3-client/service/forecastqueryservice"
)

// ForecastQueryServiceAPI provides an interface to enable mocking the
// forecastqueryservice.ForecastQueryService service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Forecast Query Service.
//    func myFunc(svc forecastqueryserviceiface.ForecastQueryServiceAPI) bool {
//        // Make svc.QueryForecast request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := forecastqueryservice.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockForecastQueryServiceClient struct {
//        forecastqueryserviceiface.ForecastQueryServiceAPI
//    }
//    func (m *mockForecastQueryServiceClient) QueryForecast(input *forecastqueryservice.QueryForecastInput) (*forecastqueryservice.QueryForecastOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockForecastQueryServiceClient{}
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
type ForecastQueryServiceAPI interface {
	QueryForecast(*forecastqueryservice.QueryForecastInput) (*forecastqueryservice.QueryForecastOutput, error)
	QueryForecastWithContext(aws.Context, *forecastqueryservice.QueryForecastInput, ...request.Option) (*forecastqueryservice.QueryForecastOutput, error)
	QueryForecastRequest(*forecastqueryservice.QueryForecastInput) (*request.Request, *forecastqueryservice.QueryForecastOutput)
}

var _ ForecastQueryServiceAPI = (*forecastqueryservice.ForecastQueryService)(nil)
