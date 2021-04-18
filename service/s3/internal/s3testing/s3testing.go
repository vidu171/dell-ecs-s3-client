package s3testing

import (
	"fmt"
	"math/rand"

	"github.com/vidu171/dell-ecs-s3-client/internal/sdkio"
	"github.com/vidu171/dell-ecs-s3-client/internal/sdkrand"
)

var randBytes = func() []byte {
	rr := rand.New(rand.NewSource(0))
	b := make([]byte, 10*sdkio.MebiByte)

	if _, err := sdkrand.Read(rr, b); err != nil {
		panic(fmt.Sprintf("failed to read random bytes, %v", err))
	}
	return b
}()

// GetTestBytes returns a pseudo-random []byte of length size
func GetTestBytes(size int) []byte {
	if len(randBytes) >= size {
		return randBytes[:size]
	}

	b := append(randBytes, GetTestBytes(size-len(randBytes))...)
	return b
}
