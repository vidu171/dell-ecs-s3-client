// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ses

import (
	"github.com/vidu171/dell-ecs-s3-client/aws"
	"github.com/vidu171/dell-ecs-s3-client/aws/client"
	"github.com/vidu171/dell-ecs-s3-client/aws/client/metadata"
	"github.com/vidu171/dell-ecs-s3-client/aws/request"
	"github.com/vidu171/dell-ecs-s3-client/aws/signer/v4"
	"github.com/vidu171/dell-ecs-s3-client/private/protocol/query"
)

// SES provides the API operation methods for making requests to
// Amazon Simple Email Service. See this package's package overview docs
// for details on the service.
//
// SES methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type SES struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "email"     // Name of service.
	EndpointsID = ServiceName // ID to lookup a service endpoint with.
	ServiceID   = "SES"       // ServiceID is a unique identifier of a specific service.
)

// New creates a new instance of the SES client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     mySession := session.Must(session.NewSession())
//
//     // Create a SES client from just a session.
//     svc := ses.New(mySession)
//
//     // Create a SES client with additional configuration
//     svc := ses.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *SES {
	c := p.ClientConfig(EndpointsID, cfgs...)
	if c.SigningNameDerived || len(c.SigningName) == 0 {
		c.SigningName = "ses"
	}
	return newClient(*c.Config, c.Handlers, c.PartitionID, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, partitionID, endpoint, signingRegion, signingName string) *SES {
	svc := &SES{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				ServiceID:     ServiceID,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				PartitionID:   partitionID,
				Endpoint:      endpoint,
				APIVersion:    "2010-12-01",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(query.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(query.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(query.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(query.UnmarshalErrorHandler)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a SES operation and runs any
// custom request initialization.
func (c *SES) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}
