package request_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/vidu171/dell-ecs-s3-client/aws/corehandlers"
	"github.com/vidu171/dell-ecs-s3-client/aws/request"
	"github.com/vidu171/dell-ecs-s3-client/awstesting"
)

func TestRequest_SetContext(t *testing.T) {
	svc := awstesting.NewClient()
	svc.Handlers.Clear()
	svc.Handlers.Send.PushBackNamed(corehandlers.SendHandler)

	r := svc.NewRequest(&request.Operation{Name: "Operation"}, nil, nil)
	ctx := &awstesting.FakeContext{DoneCh: make(chan struct{})}
	r.SetContext(ctx)

	ctx.Error = fmt.Errorf("context canceled")
	close(ctx.DoneCh)

	err := r.Send()
	if err == nil {
		t.Fatalf("expected error, got none")
	}

	// Only check against canceled because go 1.6 will not use the context's
	// Err().
	if e, a := "canceled", err.Error(); !strings.Contains(a, e) {
		t.Errorf("expect %q to be in %q, but was not", e, a)
	}
}

func TestRequest_SetContextPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Fatalf("expect SetContext to panic, did not")
		}
	}()
	r := &request.Request{}

	r.SetContext(nil)
}
