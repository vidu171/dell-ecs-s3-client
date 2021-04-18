// +build windows

package shareddefaults_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/vidu171/dell-ecs-s3-client/internal/sdktesting"
	"github.com/vidu171/dell-ecs-s3-client/internal/shareddefaults"
)

func TestSharedCredsFilename(t *testing.T) {
	restoreEnvFn := sdktesting.StashEnv()
	defer restoreEnvFn()

	os.Setenv("HOME", "home_dir")
	os.Setenv("USERPROFILE", "profile_dir")

	expect := filepath.Join("profile_dir", ".aws", "credentials")

	name := shareddefaults.SharedCredentialsFilename()
	if e, a := expect, name; e != a {
		t.Errorf("expect %q shared creds filename, got %q", e, a)
	}
}

func TestSharedConfigFilename(t *testing.T) {
	restoreEnvFn := sdktesting.StashEnv()
	defer restoreEnvFn()

	os.Setenv("HOME", "home_dir")
	os.Setenv("USERPROFILE", "profile_dir")

	expect := filepath.Join("profile_dir", ".aws", "config")

	name := shareddefaults.SharedConfigFilename()
	if e, a := expect, name; e != a {
		t.Errorf("expect %q shared config filename, got %q", e, a)
	}
}
