package apigateway

import (
	"github.com/vidu171/dell-ecs-s3-client/aws/client"
	"github.com/vidu171/dell-ecs-s3-client/aws/request"
)

func init() {
	initClient = func(c *client.Client) {
		c.Handlers.Build.PushBack(func(r *request.Request) {
			r.HTTPRequest.Header.Add("Accept", "application/json")
		})
	}
}
