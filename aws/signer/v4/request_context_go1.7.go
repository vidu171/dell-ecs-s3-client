// +build go1.7

package v4

import (
	"net/http"

	"github.com/vidu171/dell-ecs-s3-client/aws"
)

func requestContext(r *http.Request) aws.Context {
	return r.Context()
}
