package awstesting

import (
	"github.com/vidu171/dell-ecs-s3-client/aws"
	"github.com/vidu171/dell-ecs-s3-client/aws/client"
	"github.com/vidu171/dell-ecs-s3-client/aws/client/metadata"
	"github.com/vidu171/dell-ecs-s3-client/aws/defaults"
)

// NewClient creates and initializes a generic service client for testing.
func NewClient(cfgs ...*aws.Config) *client.Client {
	info := metadata.ClientInfo{
		Endpoint:    "http://endpoint",
		SigningName: "",
	}
	def := defaults.Get()
	def.Config.MergeIn(cfgs...)

	if v := aws.StringValue(def.Config.Endpoint); len(v) > 0 {
		info.Endpoint = v
	}

	return client.New(*def.Config, info, def.Handlers)
}
