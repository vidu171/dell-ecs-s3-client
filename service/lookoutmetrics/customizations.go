package lookoutmetrics

import (
	"strings"

	"github.com/vidu171/dell-ecs-s3-client/aws/client"
	"github.com/vidu171/dell-ecs-s3-client/aws/request"
)

func init() {
	initClient = func(c *client.Client) {
		c.Handlers.Build.PushBack(func(r *request.Request) {
			if strings.EqualFold(r.HTTPRequest.Header.Get("Content-Type"), "application/json") {
				r.HTTPRequest.Header.Set("Content-Type", "application/x-amz-json-1.1")
			}
		})
	}
}
