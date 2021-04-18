package session

import (
	"os"

	"github.com/vidu171/dell-ecs-s3-client/internal/sdktesting"
)

func initSessionTestEnv() (oldEnv func()) {
	oldEnv = sdktesting.StashEnv()
	os.Setenv("AWS_CONFIG_FILE", "file_not_exists")
	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", "file_not_exists")

	return oldEnv
}
