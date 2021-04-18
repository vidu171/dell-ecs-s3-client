// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// +build go1.10,integration

package costandusagereportservice_test

import (
	"context"
	"testing"
	"time"

	"github.com/vidu171/dell-ecs-s3-client/aws"
	"github.com/vidu171/dell-ecs-s3-client/aws/awserr"
	"github.com/vidu171/dell-ecs-s3-client/aws/request"
	"github.com/vidu171/dell-ecs-s3-client/awstesting/integration"
	"github.com/vidu171/dell-ecs-s3-client/service/costandusagereportservice"
)

var _ aws.Config
var _ awserr.Error
var _ request.Request

func TestInteg_00_DescribeReportDefinitions(t *testing.T) {
	ctx, cancelFn := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFn()

	sess := integration.SessionWithDefaultRegion("us-east-1")
	svc := costandusagereportservice.New(sess)
	params := &costandusagereportservice.DescribeReportDefinitionsInput{}
	_, err := svc.DescribeReportDefinitionsWithContext(ctx, params, func(r *request.Request) {
		r.Handlers.Validate.RemoveByName("core.ValidateParametersHandler")
	})
	if err != nil {
		t.Errorf("expect no error, got %v", err)
	}
}
