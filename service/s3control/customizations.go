package s3control

import (
	"github.com/vidu171/dell-ecs-s3-client/aws/client"
	"github.com/vidu171/dell-ecs-s3-client/aws/request"
	"github.com/vidu171/dell-ecs-s3-client/internal/s3shared/arn"
	"github.com/vidu171/dell-ecs-s3-client/internal/s3shared/s3err"
)

func init() {
	initClient = defaultInitClientFn
}

func defaultInitClientFn(c *client.Client) {
	// Support building custom endpoints based on config
	c.Handlers.Build.PushFrontNamed(request.NamedHandler{
		Name: "s3ControlEndpointHandler",
		Fn:   endpointHandler,
	})

	// S3 uses custom error unmarshaling logic
	c.Handlers.UnmarshalError.PushBackNamed(s3err.RequestFailureWrapperHandler())
}

// endpointARNGetter is an accessor interface to grab the
// the field corresponding to an endpoint ARN input.
type endpointARNGetter interface {
	getEndpointARN() (arn.Resource, error)
	hasEndpointARN() bool
	updateArnableField(string) (interface{}, error)
}

// endpointOutpostIDGetter is an accessor interface to grab the
// the field corresponding to an outpost ID input.
type endpointOutpostIDGetter interface {
	getOutpostID() (string, error)
	hasOutpostID() bool
}

// accountIDValidator is an accessor interface to validate the
// account id member value and account id present in endpoint ARN.
type accountIDValidator interface {
	updateAccountID(string) (interface{}, error)
}
