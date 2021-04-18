// +build go1.7,integration

package lookoutmetrics

import (
	"testing"

	"github.com/vidu171/dell-ecs-s3-client/awstesting/integration"
)

func TestInteg_ListAnomalyDetectors(t *testing.T) {
	sess := integration.SessionWithDefaultRegion("us-west-2")
	client := New(sess)
	_, err := client.ListAnomalyDetectors(&ListAnomalyDetectorsInput{})
	if err != nil {
		t.Fatalf("expect API call, got %v", err)
	}
}
