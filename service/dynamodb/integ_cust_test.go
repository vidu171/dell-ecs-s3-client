// +build go1.10,integration

package dynamodb

import (
	"testing"
	"time"

	"github.com/vidu171/dell-ecs-s3-client/aws"
	"github.com/vidu171/dell-ecs-s3-client/awstesting/integration"
)

func TestEndpointDiscovery(t *testing.T) {
	sess := integration.SessionWithDefaultRegion("us-west-2")
	svc := New(sess, &aws.Config{EnableEndpointDiscovery: aws.Bool(true)})

	req, _ := svc.ListTablesRequest(nil)
	err := req.Send()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	timeLimit := time.NewTimer(time.Second * 5)
	defer timeLimit.Stop()
	for !svc.endpointCache.Has(req.Operation.Name) {
		select {
		case <-timeLimit.C:
			t.Fatalf("timed out waiting for endpoint")
		default:
		}
	}

	req, _ = svc.ListTablesRequest(nil)
	err = req.Send()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
