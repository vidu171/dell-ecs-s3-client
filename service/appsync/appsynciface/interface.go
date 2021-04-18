// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package appsynciface provides an interface to enable mocking the AWS AppSync service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package appsynciface

import (
	"github.com/vidu171/dell-ecs-s3-client/aws"
	"github.com/vidu171/dell-ecs-s3-client/aws/request"
	"github.com/vidu171/dell-ecs-s3-client/service/appsync"
)

// AppSyncAPI provides an interface to enable mocking the
// appsync.AppSync service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS AppSync.
//    func myFunc(svc appsynciface.AppSyncAPI) bool {
//        // Make svc.CreateApiCache request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := appsync.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockAppSyncClient struct {
//        appsynciface.AppSyncAPI
//    }
//    func (m *mockAppSyncClient) CreateApiCache(input *appsync.CreateApiCacheInput) (*appsync.CreateApiCacheOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockAppSyncClient{}
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
type AppSyncAPI interface {
	CreateApiCache(*appsync.CreateApiCacheInput) (*appsync.CreateApiCacheOutput, error)
	CreateApiCacheWithContext(aws.Context, *appsync.CreateApiCacheInput, ...request.Option) (*appsync.CreateApiCacheOutput, error)
	CreateApiCacheRequest(*appsync.CreateApiCacheInput) (*request.Request, *appsync.CreateApiCacheOutput)

	CreateApiKey(*appsync.CreateApiKeyInput) (*appsync.CreateApiKeyOutput, error)
	CreateApiKeyWithContext(aws.Context, *appsync.CreateApiKeyInput, ...request.Option) (*appsync.CreateApiKeyOutput, error)
	CreateApiKeyRequest(*appsync.CreateApiKeyInput) (*request.Request, *appsync.CreateApiKeyOutput)

	CreateDataSource(*appsync.CreateDataSourceInput) (*appsync.CreateDataSourceOutput, error)
	CreateDataSourceWithContext(aws.Context, *appsync.CreateDataSourceInput, ...request.Option) (*appsync.CreateDataSourceOutput, error)
	CreateDataSourceRequest(*appsync.CreateDataSourceInput) (*request.Request, *appsync.CreateDataSourceOutput)

	CreateFunction(*appsync.CreateFunctionInput) (*appsync.CreateFunctionOutput, error)
	CreateFunctionWithContext(aws.Context, *appsync.CreateFunctionInput, ...request.Option) (*appsync.CreateFunctionOutput, error)
	CreateFunctionRequest(*appsync.CreateFunctionInput) (*request.Request, *appsync.CreateFunctionOutput)

	CreateGraphqlApi(*appsync.CreateGraphqlApiInput) (*appsync.CreateGraphqlApiOutput, error)
	CreateGraphqlApiWithContext(aws.Context, *appsync.CreateGraphqlApiInput, ...request.Option) (*appsync.CreateGraphqlApiOutput, error)
	CreateGraphqlApiRequest(*appsync.CreateGraphqlApiInput) (*request.Request, *appsync.CreateGraphqlApiOutput)

	CreateResolver(*appsync.CreateResolverInput) (*appsync.CreateResolverOutput, error)
	CreateResolverWithContext(aws.Context, *appsync.CreateResolverInput, ...request.Option) (*appsync.CreateResolverOutput, error)
	CreateResolverRequest(*appsync.CreateResolverInput) (*request.Request, *appsync.CreateResolverOutput)

	CreateType(*appsync.CreateTypeInput) (*appsync.CreateTypeOutput, error)
	CreateTypeWithContext(aws.Context, *appsync.CreateTypeInput, ...request.Option) (*appsync.CreateTypeOutput, error)
	CreateTypeRequest(*appsync.CreateTypeInput) (*request.Request, *appsync.CreateTypeOutput)

	DeleteApiCache(*appsync.DeleteApiCacheInput) (*appsync.DeleteApiCacheOutput, error)
	DeleteApiCacheWithContext(aws.Context, *appsync.DeleteApiCacheInput, ...request.Option) (*appsync.DeleteApiCacheOutput, error)
	DeleteApiCacheRequest(*appsync.DeleteApiCacheInput) (*request.Request, *appsync.DeleteApiCacheOutput)

	DeleteApiKey(*appsync.DeleteApiKeyInput) (*appsync.DeleteApiKeyOutput, error)
	DeleteApiKeyWithContext(aws.Context, *appsync.DeleteApiKeyInput, ...request.Option) (*appsync.DeleteApiKeyOutput, error)
	DeleteApiKeyRequest(*appsync.DeleteApiKeyInput) (*request.Request, *appsync.DeleteApiKeyOutput)

	DeleteDataSource(*appsync.DeleteDataSourceInput) (*appsync.DeleteDataSourceOutput, error)
	DeleteDataSourceWithContext(aws.Context, *appsync.DeleteDataSourceInput, ...request.Option) (*appsync.DeleteDataSourceOutput, error)
	DeleteDataSourceRequest(*appsync.DeleteDataSourceInput) (*request.Request, *appsync.DeleteDataSourceOutput)

	DeleteFunction(*appsync.DeleteFunctionInput) (*appsync.DeleteFunctionOutput, error)
	DeleteFunctionWithContext(aws.Context, *appsync.DeleteFunctionInput, ...request.Option) (*appsync.DeleteFunctionOutput, error)
	DeleteFunctionRequest(*appsync.DeleteFunctionInput) (*request.Request, *appsync.DeleteFunctionOutput)

	DeleteGraphqlApi(*appsync.DeleteGraphqlApiInput) (*appsync.DeleteGraphqlApiOutput, error)
	DeleteGraphqlApiWithContext(aws.Context, *appsync.DeleteGraphqlApiInput, ...request.Option) (*appsync.DeleteGraphqlApiOutput, error)
	DeleteGraphqlApiRequest(*appsync.DeleteGraphqlApiInput) (*request.Request, *appsync.DeleteGraphqlApiOutput)

	DeleteResolver(*appsync.DeleteResolverInput) (*appsync.DeleteResolverOutput, error)
	DeleteResolverWithContext(aws.Context, *appsync.DeleteResolverInput, ...request.Option) (*appsync.DeleteResolverOutput, error)
	DeleteResolverRequest(*appsync.DeleteResolverInput) (*request.Request, *appsync.DeleteResolverOutput)

	DeleteType(*appsync.DeleteTypeInput) (*appsync.DeleteTypeOutput, error)
	DeleteTypeWithContext(aws.Context, *appsync.DeleteTypeInput, ...request.Option) (*appsync.DeleteTypeOutput, error)
	DeleteTypeRequest(*appsync.DeleteTypeInput) (*request.Request, *appsync.DeleteTypeOutput)

	FlushApiCache(*appsync.FlushApiCacheInput) (*appsync.FlushApiCacheOutput, error)
	FlushApiCacheWithContext(aws.Context, *appsync.FlushApiCacheInput, ...request.Option) (*appsync.FlushApiCacheOutput, error)
	FlushApiCacheRequest(*appsync.FlushApiCacheInput) (*request.Request, *appsync.FlushApiCacheOutput)

	GetApiCache(*appsync.GetApiCacheInput) (*appsync.GetApiCacheOutput, error)
	GetApiCacheWithContext(aws.Context, *appsync.GetApiCacheInput, ...request.Option) (*appsync.GetApiCacheOutput, error)
	GetApiCacheRequest(*appsync.GetApiCacheInput) (*request.Request, *appsync.GetApiCacheOutput)

	GetDataSource(*appsync.GetDataSourceInput) (*appsync.GetDataSourceOutput, error)
	GetDataSourceWithContext(aws.Context, *appsync.GetDataSourceInput, ...request.Option) (*appsync.GetDataSourceOutput, error)
	GetDataSourceRequest(*appsync.GetDataSourceInput) (*request.Request, *appsync.GetDataSourceOutput)

	GetFunction(*appsync.GetFunctionInput) (*appsync.GetFunctionOutput, error)
	GetFunctionWithContext(aws.Context, *appsync.GetFunctionInput, ...request.Option) (*appsync.GetFunctionOutput, error)
	GetFunctionRequest(*appsync.GetFunctionInput) (*request.Request, *appsync.GetFunctionOutput)

	GetGraphqlApi(*appsync.GetGraphqlApiInput) (*appsync.GetGraphqlApiOutput, error)
	GetGraphqlApiWithContext(aws.Context, *appsync.GetGraphqlApiInput, ...request.Option) (*appsync.GetGraphqlApiOutput, error)
	GetGraphqlApiRequest(*appsync.GetGraphqlApiInput) (*request.Request, *appsync.GetGraphqlApiOutput)

	GetIntrospectionSchema(*appsync.GetIntrospectionSchemaInput) (*appsync.GetIntrospectionSchemaOutput, error)
	GetIntrospectionSchemaWithContext(aws.Context, *appsync.GetIntrospectionSchemaInput, ...request.Option) (*appsync.GetIntrospectionSchemaOutput, error)
	GetIntrospectionSchemaRequest(*appsync.GetIntrospectionSchemaInput) (*request.Request, *appsync.GetIntrospectionSchemaOutput)

	GetResolver(*appsync.GetResolverInput) (*appsync.GetResolverOutput, error)
	GetResolverWithContext(aws.Context, *appsync.GetResolverInput, ...request.Option) (*appsync.GetResolverOutput, error)
	GetResolverRequest(*appsync.GetResolverInput) (*request.Request, *appsync.GetResolverOutput)

	GetSchemaCreationStatus(*appsync.GetSchemaCreationStatusInput) (*appsync.GetSchemaCreationStatusOutput, error)
	GetSchemaCreationStatusWithContext(aws.Context, *appsync.GetSchemaCreationStatusInput, ...request.Option) (*appsync.GetSchemaCreationStatusOutput, error)
	GetSchemaCreationStatusRequest(*appsync.GetSchemaCreationStatusInput) (*request.Request, *appsync.GetSchemaCreationStatusOutput)

	GetType(*appsync.GetTypeInput) (*appsync.GetTypeOutput, error)
	GetTypeWithContext(aws.Context, *appsync.GetTypeInput, ...request.Option) (*appsync.GetTypeOutput, error)
	GetTypeRequest(*appsync.GetTypeInput) (*request.Request, *appsync.GetTypeOutput)

	ListApiKeys(*appsync.ListApiKeysInput) (*appsync.ListApiKeysOutput, error)
	ListApiKeysWithContext(aws.Context, *appsync.ListApiKeysInput, ...request.Option) (*appsync.ListApiKeysOutput, error)
	ListApiKeysRequest(*appsync.ListApiKeysInput) (*request.Request, *appsync.ListApiKeysOutput)

	ListDataSources(*appsync.ListDataSourcesInput) (*appsync.ListDataSourcesOutput, error)
	ListDataSourcesWithContext(aws.Context, *appsync.ListDataSourcesInput, ...request.Option) (*appsync.ListDataSourcesOutput, error)
	ListDataSourcesRequest(*appsync.ListDataSourcesInput) (*request.Request, *appsync.ListDataSourcesOutput)

	ListFunctions(*appsync.ListFunctionsInput) (*appsync.ListFunctionsOutput, error)
	ListFunctionsWithContext(aws.Context, *appsync.ListFunctionsInput, ...request.Option) (*appsync.ListFunctionsOutput, error)
	ListFunctionsRequest(*appsync.ListFunctionsInput) (*request.Request, *appsync.ListFunctionsOutput)

	ListGraphqlApis(*appsync.ListGraphqlApisInput) (*appsync.ListGraphqlApisOutput, error)
	ListGraphqlApisWithContext(aws.Context, *appsync.ListGraphqlApisInput, ...request.Option) (*appsync.ListGraphqlApisOutput, error)
	ListGraphqlApisRequest(*appsync.ListGraphqlApisInput) (*request.Request, *appsync.ListGraphqlApisOutput)

	ListResolvers(*appsync.ListResolversInput) (*appsync.ListResolversOutput, error)
	ListResolversWithContext(aws.Context, *appsync.ListResolversInput, ...request.Option) (*appsync.ListResolversOutput, error)
	ListResolversRequest(*appsync.ListResolversInput) (*request.Request, *appsync.ListResolversOutput)

	ListResolversByFunction(*appsync.ListResolversByFunctionInput) (*appsync.ListResolversByFunctionOutput, error)
	ListResolversByFunctionWithContext(aws.Context, *appsync.ListResolversByFunctionInput, ...request.Option) (*appsync.ListResolversByFunctionOutput, error)
	ListResolversByFunctionRequest(*appsync.ListResolversByFunctionInput) (*request.Request, *appsync.ListResolversByFunctionOutput)

	ListTagsForResource(*appsync.ListTagsForResourceInput) (*appsync.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *appsync.ListTagsForResourceInput, ...request.Option) (*appsync.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*appsync.ListTagsForResourceInput) (*request.Request, *appsync.ListTagsForResourceOutput)

	ListTypes(*appsync.ListTypesInput) (*appsync.ListTypesOutput, error)
	ListTypesWithContext(aws.Context, *appsync.ListTypesInput, ...request.Option) (*appsync.ListTypesOutput, error)
	ListTypesRequest(*appsync.ListTypesInput) (*request.Request, *appsync.ListTypesOutput)

	StartSchemaCreation(*appsync.StartSchemaCreationInput) (*appsync.StartSchemaCreationOutput, error)
	StartSchemaCreationWithContext(aws.Context, *appsync.StartSchemaCreationInput, ...request.Option) (*appsync.StartSchemaCreationOutput, error)
	StartSchemaCreationRequest(*appsync.StartSchemaCreationInput) (*request.Request, *appsync.StartSchemaCreationOutput)

	TagResource(*appsync.TagResourceInput) (*appsync.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *appsync.TagResourceInput, ...request.Option) (*appsync.TagResourceOutput, error)
	TagResourceRequest(*appsync.TagResourceInput) (*request.Request, *appsync.TagResourceOutput)

	UntagResource(*appsync.UntagResourceInput) (*appsync.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *appsync.UntagResourceInput, ...request.Option) (*appsync.UntagResourceOutput, error)
	UntagResourceRequest(*appsync.UntagResourceInput) (*request.Request, *appsync.UntagResourceOutput)

	UpdateApiCache(*appsync.UpdateApiCacheInput) (*appsync.UpdateApiCacheOutput, error)
	UpdateApiCacheWithContext(aws.Context, *appsync.UpdateApiCacheInput, ...request.Option) (*appsync.UpdateApiCacheOutput, error)
	UpdateApiCacheRequest(*appsync.UpdateApiCacheInput) (*request.Request, *appsync.UpdateApiCacheOutput)

	UpdateApiKey(*appsync.UpdateApiKeyInput) (*appsync.UpdateApiKeyOutput, error)
	UpdateApiKeyWithContext(aws.Context, *appsync.UpdateApiKeyInput, ...request.Option) (*appsync.UpdateApiKeyOutput, error)
	UpdateApiKeyRequest(*appsync.UpdateApiKeyInput) (*request.Request, *appsync.UpdateApiKeyOutput)

	UpdateDataSource(*appsync.UpdateDataSourceInput) (*appsync.UpdateDataSourceOutput, error)
	UpdateDataSourceWithContext(aws.Context, *appsync.UpdateDataSourceInput, ...request.Option) (*appsync.UpdateDataSourceOutput, error)
	UpdateDataSourceRequest(*appsync.UpdateDataSourceInput) (*request.Request, *appsync.UpdateDataSourceOutput)

	UpdateFunction(*appsync.UpdateFunctionInput) (*appsync.UpdateFunctionOutput, error)
	UpdateFunctionWithContext(aws.Context, *appsync.UpdateFunctionInput, ...request.Option) (*appsync.UpdateFunctionOutput, error)
	UpdateFunctionRequest(*appsync.UpdateFunctionInput) (*request.Request, *appsync.UpdateFunctionOutput)

	UpdateGraphqlApi(*appsync.UpdateGraphqlApiInput) (*appsync.UpdateGraphqlApiOutput, error)
	UpdateGraphqlApiWithContext(aws.Context, *appsync.UpdateGraphqlApiInput, ...request.Option) (*appsync.UpdateGraphqlApiOutput, error)
	UpdateGraphqlApiRequest(*appsync.UpdateGraphqlApiInput) (*request.Request, *appsync.UpdateGraphqlApiOutput)

	UpdateResolver(*appsync.UpdateResolverInput) (*appsync.UpdateResolverOutput, error)
	UpdateResolverWithContext(aws.Context, *appsync.UpdateResolverInput, ...request.Option) (*appsync.UpdateResolverOutput, error)
	UpdateResolverRequest(*appsync.UpdateResolverInput) (*request.Request, *appsync.UpdateResolverOutput)

	UpdateType(*appsync.UpdateTypeInput) (*appsync.UpdateTypeOutput, error)
	UpdateTypeWithContext(aws.Context, *appsync.UpdateTypeInput, ...request.Option) (*appsync.UpdateTypeOutput, error)
	UpdateTypeRequest(*appsync.UpdateTypeInput) (*request.Request, *appsync.UpdateTypeOutput)
}

var _ AppSyncAPI = (*appsync.AppSync)(nil)
