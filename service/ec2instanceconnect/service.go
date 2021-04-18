// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2instanceconnect

import (
	"github.com/vidu171/dell-ecs-s3-client/aws"
	"github.com/vidu171/dell-ecs-s3-client/aws/client"
	"github.com/vidu171/dell-ecs-s3-client/aws/client/metadata"
	"github.com/vidu171/dell-ecs-s3-client/aws/request"
	"github.com/vidu171/dell-ecs-s3-client/aws/signer/v4"
	"github.com/vidu171/dell-ecs-s3-client/private/protocol"
	"github.com/vidu171/dell-ecs-s3-client/private/protocol/jsonrpc"
)

// EC2InstanceConnect provides the API operation methods for making requests to
// AWS EC2 Instance Connect. See this package's package overview docs
// for details on the service.
//
// EC2InstanceConnect methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type EC2InstanceConnect struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "EC2 Instance Connect" // Name of service.
	EndpointsID = "ec2-instance-connect" // ID to lookup a service endpoint with.
	ServiceID   = "EC2 Instance Connect" // ServiceID is a unique identifier of a specific service.
)

// New creates a new instance of the EC2InstanceConnect client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     mySession := session.Must(session.NewSession())
//
//     // Create a EC2InstanceConnect client from just a session.
//     svc := ec2instanceconnect.New(mySession)
//
//     // Create a EC2InstanceConnect client with additional configuration
//     svc := ec2instanceconnect.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *EC2InstanceConnect {
	c := p.ClientConfig(EndpointsID, cfgs...)
	return newClient(*c.Config, c.Handlers, c.PartitionID, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, partitionID, endpoint, signingRegion, signingName string) *EC2InstanceConnect {
	svc := &EC2InstanceConnect{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				ServiceID:     ServiceID,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				PartitionID:   partitionID,
				Endpoint:      endpoint,
				APIVersion:    "2018-04-02",
				JSONVersion:   "1.1",
				TargetPrefix:  "AWSEC2InstanceConnectService",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(jsonrpc.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(jsonrpc.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(jsonrpc.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(
		protocol.NewUnmarshalErrorHandler(jsonrpc.NewUnmarshalTypedError(exceptionFromCode)).NamedHandler(),
	)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a EC2InstanceConnect operation and runs any
// custom request initialization.
func (c *EC2InstanceConnect) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
