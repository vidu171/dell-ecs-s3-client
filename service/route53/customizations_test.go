package route53_test

import (
	"strings"
	"testing"

	"github.com/vidu171/dell-ecs-s3-client/aws"
	"github.com/vidu171/dell-ecs-s3-client/awstesting/unit"
	"github.com/vidu171/dell-ecs-s3-client/service/route53"
)

func TestBuildCorrectURI(t *testing.T) {
	svc := route53.New(unit.Session)
	svc.Handlers.Validate.Clear()
	req, _ := svc.GetHostedZoneRequest(&route53.GetHostedZoneInput{
		Id: aws.String("/hostedzone/ABCDEFG"),
	})

	expectPath := strings.Replace(req.Operation.HTTPPath, "{Id}", "ABCDEFG", -1)

	req.HTTPRequest.URL.RawQuery = "abc=123"

	req.Build()

	if a, e := req.HTTPRequest.URL.Path, expectPath; a != e {
		t.Errorf("expect path %q, got %q", e, a)
	}

	if a, e := req.HTTPRequest.URL.RawPath, expectPath; a != e {
		t.Errorf("expect raw path %q, got %q", e, a)
	}

	if a, e := req.HTTPRequest.URL.RawQuery, "abc=123"; a != e {
		t.Errorf("expect query to be %q, got %q", e, a)
	}
}
