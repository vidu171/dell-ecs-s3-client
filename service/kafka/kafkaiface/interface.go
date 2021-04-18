// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package kafkaiface provides an interface to enable mocking the Managed Streaming for Kafka service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package kafkaiface

import (
	"github.com/vidu171/dell-ecs-s3-client/aws"
	"github.com/vidu171/dell-ecs-s3-client/aws/request"
	"github.com/vidu171/dell-ecs-s3-client/service/kafka"
)

// KafkaAPI provides an interface to enable mocking the
// kafka.Kafka service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Managed Streaming for Kafka.
//    func myFunc(svc kafkaiface.KafkaAPI) bool {
//        // Make svc.BatchAssociateScramSecret request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := kafka.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockKafkaClient struct {
//        kafkaiface.KafkaAPI
//    }
//    func (m *mockKafkaClient) BatchAssociateScramSecret(input *kafka.BatchAssociateScramSecretInput) (*kafka.BatchAssociateScramSecretOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockKafkaClient{}
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
type KafkaAPI interface {
	BatchAssociateScramSecret(*kafka.BatchAssociateScramSecretInput) (*kafka.BatchAssociateScramSecretOutput, error)
	BatchAssociateScramSecretWithContext(aws.Context, *kafka.BatchAssociateScramSecretInput, ...request.Option) (*kafka.BatchAssociateScramSecretOutput, error)
	BatchAssociateScramSecretRequest(*kafka.BatchAssociateScramSecretInput) (*request.Request, *kafka.BatchAssociateScramSecretOutput)

	BatchDisassociateScramSecret(*kafka.BatchDisassociateScramSecretInput) (*kafka.BatchDisassociateScramSecretOutput, error)
	BatchDisassociateScramSecretWithContext(aws.Context, *kafka.BatchDisassociateScramSecretInput, ...request.Option) (*kafka.BatchDisassociateScramSecretOutput, error)
	BatchDisassociateScramSecretRequest(*kafka.BatchDisassociateScramSecretInput) (*request.Request, *kafka.BatchDisassociateScramSecretOutput)

	CreateCluster(*kafka.CreateClusterInput) (*kafka.CreateClusterOutput, error)
	CreateClusterWithContext(aws.Context, *kafka.CreateClusterInput, ...request.Option) (*kafka.CreateClusterOutput, error)
	CreateClusterRequest(*kafka.CreateClusterInput) (*request.Request, *kafka.CreateClusterOutput)

	CreateConfiguration(*kafka.CreateConfigurationInput) (*kafka.CreateConfigurationOutput, error)
	CreateConfigurationWithContext(aws.Context, *kafka.CreateConfigurationInput, ...request.Option) (*kafka.CreateConfigurationOutput, error)
	CreateConfigurationRequest(*kafka.CreateConfigurationInput) (*request.Request, *kafka.CreateConfigurationOutput)

	DeleteCluster(*kafka.DeleteClusterInput) (*kafka.DeleteClusterOutput, error)
	DeleteClusterWithContext(aws.Context, *kafka.DeleteClusterInput, ...request.Option) (*kafka.DeleteClusterOutput, error)
	DeleteClusterRequest(*kafka.DeleteClusterInput) (*request.Request, *kafka.DeleteClusterOutput)

	DeleteConfiguration(*kafka.DeleteConfigurationInput) (*kafka.DeleteConfigurationOutput, error)
	DeleteConfigurationWithContext(aws.Context, *kafka.DeleteConfigurationInput, ...request.Option) (*kafka.DeleteConfigurationOutput, error)
	DeleteConfigurationRequest(*kafka.DeleteConfigurationInput) (*request.Request, *kafka.DeleteConfigurationOutput)

	DescribeCluster(*kafka.DescribeClusterInput) (*kafka.DescribeClusterOutput, error)
	DescribeClusterWithContext(aws.Context, *kafka.DescribeClusterInput, ...request.Option) (*kafka.DescribeClusterOutput, error)
	DescribeClusterRequest(*kafka.DescribeClusterInput) (*request.Request, *kafka.DescribeClusterOutput)

	DescribeClusterOperation(*kafka.DescribeClusterOperationInput) (*kafka.DescribeClusterOperationOutput, error)
	DescribeClusterOperationWithContext(aws.Context, *kafka.DescribeClusterOperationInput, ...request.Option) (*kafka.DescribeClusterOperationOutput, error)
	DescribeClusterOperationRequest(*kafka.DescribeClusterOperationInput) (*request.Request, *kafka.DescribeClusterOperationOutput)

	DescribeConfiguration(*kafka.DescribeConfigurationInput) (*kafka.DescribeConfigurationOutput, error)
	DescribeConfigurationWithContext(aws.Context, *kafka.DescribeConfigurationInput, ...request.Option) (*kafka.DescribeConfigurationOutput, error)
	DescribeConfigurationRequest(*kafka.DescribeConfigurationInput) (*request.Request, *kafka.DescribeConfigurationOutput)

	DescribeConfigurationRevision(*kafka.DescribeConfigurationRevisionInput) (*kafka.DescribeConfigurationRevisionOutput, error)
	DescribeConfigurationRevisionWithContext(aws.Context, *kafka.DescribeConfigurationRevisionInput, ...request.Option) (*kafka.DescribeConfigurationRevisionOutput, error)
	DescribeConfigurationRevisionRequest(*kafka.DescribeConfigurationRevisionInput) (*request.Request, *kafka.DescribeConfigurationRevisionOutput)

	GetBootstrapBrokers(*kafka.GetBootstrapBrokersInput) (*kafka.GetBootstrapBrokersOutput, error)
	GetBootstrapBrokersWithContext(aws.Context, *kafka.GetBootstrapBrokersInput, ...request.Option) (*kafka.GetBootstrapBrokersOutput, error)
	GetBootstrapBrokersRequest(*kafka.GetBootstrapBrokersInput) (*request.Request, *kafka.GetBootstrapBrokersOutput)

	GetCompatibleKafkaVersions(*kafka.GetCompatibleKafkaVersionsInput) (*kafka.GetCompatibleKafkaVersionsOutput, error)
	GetCompatibleKafkaVersionsWithContext(aws.Context, *kafka.GetCompatibleKafkaVersionsInput, ...request.Option) (*kafka.GetCompatibleKafkaVersionsOutput, error)
	GetCompatibleKafkaVersionsRequest(*kafka.GetCompatibleKafkaVersionsInput) (*request.Request, *kafka.GetCompatibleKafkaVersionsOutput)

	ListClusterOperations(*kafka.ListClusterOperationsInput) (*kafka.ListClusterOperationsOutput, error)
	ListClusterOperationsWithContext(aws.Context, *kafka.ListClusterOperationsInput, ...request.Option) (*kafka.ListClusterOperationsOutput, error)
	ListClusterOperationsRequest(*kafka.ListClusterOperationsInput) (*request.Request, *kafka.ListClusterOperationsOutput)

	ListClusterOperationsPages(*kafka.ListClusterOperationsInput, func(*kafka.ListClusterOperationsOutput, bool) bool) error
	ListClusterOperationsPagesWithContext(aws.Context, *kafka.ListClusterOperationsInput, func(*kafka.ListClusterOperationsOutput, bool) bool, ...request.Option) error

	ListClusters(*kafka.ListClustersInput) (*kafka.ListClustersOutput, error)
	ListClustersWithContext(aws.Context, *kafka.ListClustersInput, ...request.Option) (*kafka.ListClustersOutput, error)
	ListClustersRequest(*kafka.ListClustersInput) (*request.Request, *kafka.ListClustersOutput)

	ListClustersPages(*kafka.ListClustersInput, func(*kafka.ListClustersOutput, bool) bool) error
	ListClustersPagesWithContext(aws.Context, *kafka.ListClustersInput, func(*kafka.ListClustersOutput, bool) bool, ...request.Option) error

	ListConfigurationRevisions(*kafka.ListConfigurationRevisionsInput) (*kafka.ListConfigurationRevisionsOutput, error)
	ListConfigurationRevisionsWithContext(aws.Context, *kafka.ListConfigurationRevisionsInput, ...request.Option) (*kafka.ListConfigurationRevisionsOutput, error)
	ListConfigurationRevisionsRequest(*kafka.ListConfigurationRevisionsInput) (*request.Request, *kafka.ListConfigurationRevisionsOutput)

	ListConfigurationRevisionsPages(*kafka.ListConfigurationRevisionsInput, func(*kafka.ListConfigurationRevisionsOutput, bool) bool) error
	ListConfigurationRevisionsPagesWithContext(aws.Context, *kafka.ListConfigurationRevisionsInput, func(*kafka.ListConfigurationRevisionsOutput, bool) bool, ...request.Option) error

	ListConfigurations(*kafka.ListConfigurationsInput) (*kafka.ListConfigurationsOutput, error)
	ListConfigurationsWithContext(aws.Context, *kafka.ListConfigurationsInput, ...request.Option) (*kafka.ListConfigurationsOutput, error)
	ListConfigurationsRequest(*kafka.ListConfigurationsInput) (*request.Request, *kafka.ListConfigurationsOutput)

	ListConfigurationsPages(*kafka.ListConfigurationsInput, func(*kafka.ListConfigurationsOutput, bool) bool) error
	ListConfigurationsPagesWithContext(aws.Context, *kafka.ListConfigurationsInput, func(*kafka.ListConfigurationsOutput, bool) bool, ...request.Option) error

	ListKafkaVersions(*kafka.ListKafkaVersionsInput) (*kafka.ListKafkaVersionsOutput, error)
	ListKafkaVersionsWithContext(aws.Context, *kafka.ListKafkaVersionsInput, ...request.Option) (*kafka.ListKafkaVersionsOutput, error)
	ListKafkaVersionsRequest(*kafka.ListKafkaVersionsInput) (*request.Request, *kafka.ListKafkaVersionsOutput)

	ListKafkaVersionsPages(*kafka.ListKafkaVersionsInput, func(*kafka.ListKafkaVersionsOutput, bool) bool) error
	ListKafkaVersionsPagesWithContext(aws.Context, *kafka.ListKafkaVersionsInput, func(*kafka.ListKafkaVersionsOutput, bool) bool, ...request.Option) error

	ListNodes(*kafka.ListNodesInput) (*kafka.ListNodesOutput, error)
	ListNodesWithContext(aws.Context, *kafka.ListNodesInput, ...request.Option) (*kafka.ListNodesOutput, error)
	ListNodesRequest(*kafka.ListNodesInput) (*request.Request, *kafka.ListNodesOutput)

	ListNodesPages(*kafka.ListNodesInput, func(*kafka.ListNodesOutput, bool) bool) error
	ListNodesPagesWithContext(aws.Context, *kafka.ListNodesInput, func(*kafka.ListNodesOutput, bool) bool, ...request.Option) error

	ListScramSecrets(*kafka.ListScramSecretsInput) (*kafka.ListScramSecretsOutput, error)
	ListScramSecretsWithContext(aws.Context, *kafka.ListScramSecretsInput, ...request.Option) (*kafka.ListScramSecretsOutput, error)
	ListScramSecretsRequest(*kafka.ListScramSecretsInput) (*request.Request, *kafka.ListScramSecretsOutput)

	ListScramSecretsPages(*kafka.ListScramSecretsInput, func(*kafka.ListScramSecretsOutput, bool) bool) error
	ListScramSecretsPagesWithContext(aws.Context, *kafka.ListScramSecretsInput, func(*kafka.ListScramSecretsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*kafka.ListTagsForResourceInput) (*kafka.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *kafka.ListTagsForResourceInput, ...request.Option) (*kafka.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*kafka.ListTagsForResourceInput) (*request.Request, *kafka.ListTagsForResourceOutput)

	RebootBroker(*kafka.RebootBrokerInput) (*kafka.RebootBrokerOutput, error)
	RebootBrokerWithContext(aws.Context, *kafka.RebootBrokerInput, ...request.Option) (*kafka.RebootBrokerOutput, error)
	RebootBrokerRequest(*kafka.RebootBrokerInput) (*request.Request, *kafka.RebootBrokerOutput)

	TagResource(*kafka.TagResourceInput) (*kafka.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *kafka.TagResourceInput, ...request.Option) (*kafka.TagResourceOutput, error)
	TagResourceRequest(*kafka.TagResourceInput) (*request.Request, *kafka.TagResourceOutput)

	UntagResource(*kafka.UntagResourceInput) (*kafka.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *kafka.UntagResourceInput, ...request.Option) (*kafka.UntagResourceOutput, error)
	UntagResourceRequest(*kafka.UntagResourceInput) (*request.Request, *kafka.UntagResourceOutput)

	UpdateBrokerCount(*kafka.UpdateBrokerCountInput) (*kafka.UpdateBrokerCountOutput, error)
	UpdateBrokerCountWithContext(aws.Context, *kafka.UpdateBrokerCountInput, ...request.Option) (*kafka.UpdateBrokerCountOutput, error)
	UpdateBrokerCountRequest(*kafka.UpdateBrokerCountInput) (*request.Request, *kafka.UpdateBrokerCountOutput)

	UpdateBrokerStorage(*kafka.UpdateBrokerStorageInput) (*kafka.UpdateBrokerStorageOutput, error)
	UpdateBrokerStorageWithContext(aws.Context, *kafka.UpdateBrokerStorageInput, ...request.Option) (*kafka.UpdateBrokerStorageOutput, error)
	UpdateBrokerStorageRequest(*kafka.UpdateBrokerStorageInput) (*request.Request, *kafka.UpdateBrokerStorageOutput)

	UpdateBrokerType(*kafka.UpdateBrokerTypeInput) (*kafka.UpdateBrokerTypeOutput, error)
	UpdateBrokerTypeWithContext(aws.Context, *kafka.UpdateBrokerTypeInput, ...request.Option) (*kafka.UpdateBrokerTypeOutput, error)
	UpdateBrokerTypeRequest(*kafka.UpdateBrokerTypeInput) (*request.Request, *kafka.UpdateBrokerTypeOutput)

	UpdateClusterConfiguration(*kafka.UpdateClusterConfigurationInput) (*kafka.UpdateClusterConfigurationOutput, error)
	UpdateClusterConfigurationWithContext(aws.Context, *kafka.UpdateClusterConfigurationInput, ...request.Option) (*kafka.UpdateClusterConfigurationOutput, error)
	UpdateClusterConfigurationRequest(*kafka.UpdateClusterConfigurationInput) (*request.Request, *kafka.UpdateClusterConfigurationOutput)

	UpdateClusterKafkaVersion(*kafka.UpdateClusterKafkaVersionInput) (*kafka.UpdateClusterKafkaVersionOutput, error)
	UpdateClusterKafkaVersionWithContext(aws.Context, *kafka.UpdateClusterKafkaVersionInput, ...request.Option) (*kafka.UpdateClusterKafkaVersionOutput, error)
	UpdateClusterKafkaVersionRequest(*kafka.UpdateClusterKafkaVersionInput) (*request.Request, *kafka.UpdateClusterKafkaVersionOutput)

	UpdateConfiguration(*kafka.UpdateConfigurationInput) (*kafka.UpdateConfigurationOutput, error)
	UpdateConfigurationWithContext(aws.Context, *kafka.UpdateConfigurationInput, ...request.Option) (*kafka.UpdateConfigurationOutput, error)
	UpdateConfigurationRequest(*kafka.UpdateConfigurationInput) (*request.Request, *kafka.UpdateConfigurationOutput)

	UpdateMonitoring(*kafka.UpdateMonitoringInput) (*kafka.UpdateMonitoringOutput, error)
	UpdateMonitoringWithContext(aws.Context, *kafka.UpdateMonitoringInput, ...request.Option) (*kafka.UpdateMonitoringOutput, error)
	UpdateMonitoringRequest(*kafka.UpdateMonitoringInput) (*request.Request, *kafka.UpdateMonitoringOutput)
}

var _ KafkaAPI = (*kafka.Kafka)(nil)
