// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package healthlakeiface provides an interface to enable mocking the Amazon HealthLake service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package healthlakeiface

import (
	"github.com/vidu171/dell-ecs-s3-client/aws"
	"github.com/vidu171/dell-ecs-s3-client/aws/request"
	"github.com/vidu171/dell-ecs-s3-client/service/healthlake"
)

// HealthLakeAPI provides an interface to enable mocking the
// healthlake.HealthLake service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon HealthLake.
//    func myFunc(svc healthlakeiface.HealthLakeAPI) bool {
//        // Make svc.CreateFHIRDatastore request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := healthlake.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockHealthLakeClient struct {
//        healthlakeiface.HealthLakeAPI
//    }
//    func (m *mockHealthLakeClient) CreateFHIRDatastore(input *healthlake.CreateFHIRDatastoreInput) (*healthlake.CreateFHIRDatastoreOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockHealthLakeClient{}
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
type HealthLakeAPI interface {
	CreateFHIRDatastore(*healthlake.CreateFHIRDatastoreInput) (*healthlake.CreateFHIRDatastoreOutput, error)
	CreateFHIRDatastoreWithContext(aws.Context, *healthlake.CreateFHIRDatastoreInput, ...request.Option) (*healthlake.CreateFHIRDatastoreOutput, error)
	CreateFHIRDatastoreRequest(*healthlake.CreateFHIRDatastoreInput) (*request.Request, *healthlake.CreateFHIRDatastoreOutput)

	DeleteFHIRDatastore(*healthlake.DeleteFHIRDatastoreInput) (*healthlake.DeleteFHIRDatastoreOutput, error)
	DeleteFHIRDatastoreWithContext(aws.Context, *healthlake.DeleteFHIRDatastoreInput, ...request.Option) (*healthlake.DeleteFHIRDatastoreOutput, error)
	DeleteFHIRDatastoreRequest(*healthlake.DeleteFHIRDatastoreInput) (*request.Request, *healthlake.DeleteFHIRDatastoreOutput)

	DescribeFHIRDatastore(*healthlake.DescribeFHIRDatastoreInput) (*healthlake.DescribeFHIRDatastoreOutput, error)
	DescribeFHIRDatastoreWithContext(aws.Context, *healthlake.DescribeFHIRDatastoreInput, ...request.Option) (*healthlake.DescribeFHIRDatastoreOutput, error)
	DescribeFHIRDatastoreRequest(*healthlake.DescribeFHIRDatastoreInput) (*request.Request, *healthlake.DescribeFHIRDatastoreOutput)

	DescribeFHIRExportJob(*healthlake.DescribeFHIRExportJobInput) (*healthlake.DescribeFHIRExportJobOutput, error)
	DescribeFHIRExportJobWithContext(aws.Context, *healthlake.DescribeFHIRExportJobInput, ...request.Option) (*healthlake.DescribeFHIRExportJobOutput, error)
	DescribeFHIRExportJobRequest(*healthlake.DescribeFHIRExportJobInput) (*request.Request, *healthlake.DescribeFHIRExportJobOutput)

	DescribeFHIRImportJob(*healthlake.DescribeFHIRImportJobInput) (*healthlake.DescribeFHIRImportJobOutput, error)
	DescribeFHIRImportJobWithContext(aws.Context, *healthlake.DescribeFHIRImportJobInput, ...request.Option) (*healthlake.DescribeFHIRImportJobOutput, error)
	DescribeFHIRImportJobRequest(*healthlake.DescribeFHIRImportJobInput) (*request.Request, *healthlake.DescribeFHIRImportJobOutput)

	ListFHIRDatastores(*healthlake.ListFHIRDatastoresInput) (*healthlake.ListFHIRDatastoresOutput, error)
	ListFHIRDatastoresWithContext(aws.Context, *healthlake.ListFHIRDatastoresInput, ...request.Option) (*healthlake.ListFHIRDatastoresOutput, error)
	ListFHIRDatastoresRequest(*healthlake.ListFHIRDatastoresInput) (*request.Request, *healthlake.ListFHIRDatastoresOutput)

	ListFHIRDatastoresPages(*healthlake.ListFHIRDatastoresInput, func(*healthlake.ListFHIRDatastoresOutput, bool) bool) error
	ListFHIRDatastoresPagesWithContext(aws.Context, *healthlake.ListFHIRDatastoresInput, func(*healthlake.ListFHIRDatastoresOutput, bool) bool, ...request.Option) error

	StartFHIRExportJob(*healthlake.StartFHIRExportJobInput) (*healthlake.StartFHIRExportJobOutput, error)
	StartFHIRExportJobWithContext(aws.Context, *healthlake.StartFHIRExportJobInput, ...request.Option) (*healthlake.StartFHIRExportJobOutput, error)
	StartFHIRExportJobRequest(*healthlake.StartFHIRExportJobInput) (*request.Request, *healthlake.StartFHIRExportJobOutput)

	StartFHIRImportJob(*healthlake.StartFHIRImportJobInput) (*healthlake.StartFHIRImportJobOutput, error)
	StartFHIRImportJobWithContext(aws.Context, *healthlake.StartFHIRImportJobInput, ...request.Option) (*healthlake.StartFHIRImportJobOutput, error)
	StartFHIRImportJobRequest(*healthlake.StartFHIRImportJobInput) (*request.Request, *healthlake.StartFHIRImportJobOutput)
}

var _ HealthLakeAPI = (*healthlake.HealthLake)(nil)
