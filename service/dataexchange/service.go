// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dataexchange

import (
	"github.com/vidu171/dell-ecs-s3-client/aws"
	"github.com/vidu171/dell-ecs-s3-client/aws/client"
	"github.com/vidu171/dell-ecs-s3-client/aws/client/metadata"
	"github.com/vidu171/dell-ecs-s3-client/aws/request"
	"github.com/vidu171/dell-ecs-s3-client/aws/signer/v4"
	"github.com/vidu171/dell-ecs-s3-client/private/protocol"
	"github.com/vidu171/dell-ecs-s3-client/private/protocol/restjson"
)

// DataExchange provides the API operation methods for making requests to
// AWS Data Exchange. See this package's package overview docs
// for details on the service.
//
// DataExchange methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type DataExchange struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "DataExchange" // Name of service.
	EndpointsID = "dataexchange" // ID to lookup a service endpoint with.
	ServiceID   = "DataExchange" // ServiceID is a unique identifier of a specific service.
)

// New creates a new instance of the DataExchange client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//     mySession := session.Must(session.NewSession())
//
//     // Create a DataExchange client from just a session.
//     svc := dataexchange.New(mySession)
//
//     // Create a DataExchange client with additional configuration
//     svc := dataexchange.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *DataExchange {
	c := p.ClientConfig(EndpointsID, cfgs...)
	if c.SigningNameDerived || len(c.SigningName) == 0 {
		c.SigningName = "dataexchange"
	}
	return newClient(*c.Config, c.Handlers, c.PartitionID, c.Endpoint, c.SigningRegion, c.SigningName)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, partitionID, endpoint, signingRegion, signingName string) *DataExchange {
	svc := &DataExchange{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:   ServiceName,
				ServiceID:     ServiceID,
				SigningName:   signingName,
				SigningRegion: signingRegion,
				PartitionID:   partitionID,
				Endpoint:      endpoint,
				APIVersion:    "2017-07-25",
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restjson.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restjson.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restjson.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(
		protocol.NewUnmarshalErrorHandler(restjson.NewUnmarshalTypedError(exceptionFromCode)).NamedHandler(),
	)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a DataExchange operation and runs any
// custom request initialization.
func (c *DataExchange) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}