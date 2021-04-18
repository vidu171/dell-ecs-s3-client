// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package computeoptimizeriface provides an interface to enable mocking the AWS Compute Optimizer service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package computeoptimizeriface

import (
	"github.com/vidu171/dell-ecs-s3-client/aws"
	"github.com/vidu171/dell-ecs-s3-client/aws/request"
	"github.com/vidu171/dell-ecs-s3-client/service/computeoptimizer"
)

// ComputeOptimizerAPI provides an interface to enable mocking the
// computeoptimizer.ComputeOptimizer service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Compute Optimizer.
//    func myFunc(svc computeoptimizeriface.ComputeOptimizerAPI) bool {
//        // Make svc.DescribeRecommendationExportJobs request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := computeoptimizer.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockComputeOptimizerClient struct {
//        computeoptimizeriface.ComputeOptimizerAPI
//    }
//    func (m *mockComputeOptimizerClient) DescribeRecommendationExportJobs(input *computeoptimizer.DescribeRecommendationExportJobsInput) (*computeoptimizer.DescribeRecommendationExportJobsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockComputeOptimizerClient{}
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
type ComputeOptimizerAPI interface {
	DescribeRecommendationExportJobs(*computeoptimizer.DescribeRecommendationExportJobsInput) (*computeoptimizer.DescribeRecommendationExportJobsOutput, error)
	DescribeRecommendationExportJobsWithContext(aws.Context, *computeoptimizer.DescribeRecommendationExportJobsInput, ...request.Option) (*computeoptimizer.DescribeRecommendationExportJobsOutput, error)
	DescribeRecommendationExportJobsRequest(*computeoptimizer.DescribeRecommendationExportJobsInput) (*request.Request, *computeoptimizer.DescribeRecommendationExportJobsOutput)

	ExportAutoScalingGroupRecommendations(*computeoptimizer.ExportAutoScalingGroupRecommendationsInput) (*computeoptimizer.ExportAutoScalingGroupRecommendationsOutput, error)
	ExportAutoScalingGroupRecommendationsWithContext(aws.Context, *computeoptimizer.ExportAutoScalingGroupRecommendationsInput, ...request.Option) (*computeoptimizer.ExportAutoScalingGroupRecommendationsOutput, error)
	ExportAutoScalingGroupRecommendationsRequest(*computeoptimizer.ExportAutoScalingGroupRecommendationsInput) (*request.Request, *computeoptimizer.ExportAutoScalingGroupRecommendationsOutput)

	ExportEC2InstanceRecommendations(*computeoptimizer.ExportEC2InstanceRecommendationsInput) (*computeoptimizer.ExportEC2InstanceRecommendationsOutput, error)
	ExportEC2InstanceRecommendationsWithContext(aws.Context, *computeoptimizer.ExportEC2InstanceRecommendationsInput, ...request.Option) (*computeoptimizer.ExportEC2InstanceRecommendationsOutput, error)
	ExportEC2InstanceRecommendationsRequest(*computeoptimizer.ExportEC2InstanceRecommendationsInput) (*request.Request, *computeoptimizer.ExportEC2InstanceRecommendationsOutput)

	GetAutoScalingGroupRecommendations(*computeoptimizer.GetAutoScalingGroupRecommendationsInput) (*computeoptimizer.GetAutoScalingGroupRecommendationsOutput, error)
	GetAutoScalingGroupRecommendationsWithContext(aws.Context, *computeoptimizer.GetAutoScalingGroupRecommendationsInput, ...request.Option) (*computeoptimizer.GetAutoScalingGroupRecommendationsOutput, error)
	GetAutoScalingGroupRecommendationsRequest(*computeoptimizer.GetAutoScalingGroupRecommendationsInput) (*request.Request, *computeoptimizer.GetAutoScalingGroupRecommendationsOutput)

	GetEBSVolumeRecommendations(*computeoptimizer.GetEBSVolumeRecommendationsInput) (*computeoptimizer.GetEBSVolumeRecommendationsOutput, error)
	GetEBSVolumeRecommendationsWithContext(aws.Context, *computeoptimizer.GetEBSVolumeRecommendationsInput, ...request.Option) (*computeoptimizer.GetEBSVolumeRecommendationsOutput, error)
	GetEBSVolumeRecommendationsRequest(*computeoptimizer.GetEBSVolumeRecommendationsInput) (*request.Request, *computeoptimizer.GetEBSVolumeRecommendationsOutput)

	GetEC2InstanceRecommendations(*computeoptimizer.GetEC2InstanceRecommendationsInput) (*computeoptimizer.GetEC2InstanceRecommendationsOutput, error)
	GetEC2InstanceRecommendationsWithContext(aws.Context, *computeoptimizer.GetEC2InstanceRecommendationsInput, ...request.Option) (*computeoptimizer.GetEC2InstanceRecommendationsOutput, error)
	GetEC2InstanceRecommendationsRequest(*computeoptimizer.GetEC2InstanceRecommendationsInput) (*request.Request, *computeoptimizer.GetEC2InstanceRecommendationsOutput)

	GetEC2RecommendationProjectedMetrics(*computeoptimizer.GetEC2RecommendationProjectedMetricsInput) (*computeoptimizer.GetEC2RecommendationProjectedMetricsOutput, error)
	GetEC2RecommendationProjectedMetricsWithContext(aws.Context, *computeoptimizer.GetEC2RecommendationProjectedMetricsInput, ...request.Option) (*computeoptimizer.GetEC2RecommendationProjectedMetricsOutput, error)
	GetEC2RecommendationProjectedMetricsRequest(*computeoptimizer.GetEC2RecommendationProjectedMetricsInput) (*request.Request, *computeoptimizer.GetEC2RecommendationProjectedMetricsOutput)

	GetEnrollmentStatus(*computeoptimizer.GetEnrollmentStatusInput) (*computeoptimizer.GetEnrollmentStatusOutput, error)
	GetEnrollmentStatusWithContext(aws.Context, *computeoptimizer.GetEnrollmentStatusInput, ...request.Option) (*computeoptimizer.GetEnrollmentStatusOutput, error)
	GetEnrollmentStatusRequest(*computeoptimizer.GetEnrollmentStatusInput) (*request.Request, *computeoptimizer.GetEnrollmentStatusOutput)

	GetLambdaFunctionRecommendations(*computeoptimizer.GetLambdaFunctionRecommendationsInput) (*computeoptimizer.GetLambdaFunctionRecommendationsOutput, error)
	GetLambdaFunctionRecommendationsWithContext(aws.Context, *computeoptimizer.GetLambdaFunctionRecommendationsInput, ...request.Option) (*computeoptimizer.GetLambdaFunctionRecommendationsOutput, error)
	GetLambdaFunctionRecommendationsRequest(*computeoptimizer.GetLambdaFunctionRecommendationsInput) (*request.Request, *computeoptimizer.GetLambdaFunctionRecommendationsOutput)

	GetRecommendationSummaries(*computeoptimizer.GetRecommendationSummariesInput) (*computeoptimizer.GetRecommendationSummariesOutput, error)
	GetRecommendationSummariesWithContext(aws.Context, *computeoptimizer.GetRecommendationSummariesInput, ...request.Option) (*computeoptimizer.GetRecommendationSummariesOutput, error)
	GetRecommendationSummariesRequest(*computeoptimizer.GetRecommendationSummariesInput) (*request.Request, *computeoptimizer.GetRecommendationSummariesOutput)

	UpdateEnrollmentStatus(*computeoptimizer.UpdateEnrollmentStatusInput) (*computeoptimizer.UpdateEnrollmentStatusOutput, error)
	UpdateEnrollmentStatusWithContext(aws.Context, *computeoptimizer.UpdateEnrollmentStatusInput, ...request.Option) (*computeoptimizer.UpdateEnrollmentStatusOutput, error)
	UpdateEnrollmentStatusRequest(*computeoptimizer.UpdateEnrollmentStatusInput) (*request.Request, *computeoptimizer.UpdateEnrollmentStatusOutput)
}

var _ ComputeOptimizerAPI = (*computeoptimizer.ComputeOptimizer)(nil)
