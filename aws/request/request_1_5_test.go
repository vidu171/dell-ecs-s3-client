// +build !go1.6

package request_test

import (
	"errors"

	"github.com/vidu171/dell-ecs-s3-client/aws/awserr"
)

var errTimeout = awserr.New("foo", "bar", errors.New("net/http: request canceled Timeout"))
