// Package unit performs initialization and validation for unit tests
package unit

import (
	"time"

	"github.com/vidu171/dell-ecs-s3-client/aws"
	"github.com/vidu171/dell-ecs-s3-client/aws/credentials"
	"github.com/vidu171/dell-ecs-s3-client/aws/session"
)

// Session is a shared session for unit tests to use.
var Session = session.Must(session.NewSession(&aws.Config{
	Credentials: credentials.NewStaticCredentials("AKID", "SECRET", "SESSION"),
	Region:      aws.String("mock-region"),
	SleepDelay:  func(time.Duration) {},
}))
