package aws

import "github.com/vidu171/dell-ecs-s3-client/aws/awserr"

var (
	// ErrMissingRegion is an error that is returned if region configuration is
	// not found.
	ErrMissingRegion = awserr.New("MissingRegion", "could not find region configuration", nil)

	// ErrMissingEndpoint is an error that is returned if an endpoint cannot be
	// resolved for a service.
	ErrMissingEndpoint = awserr.New("MissingEndpoint", "'Endpoint' configuration is required for this service", nil)
)
