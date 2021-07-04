package registry

import (
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/stretchr/testify/assert"
	"os/exec"
	"testing"
)

func InitRegistryForTest(t *testing.T) {
	output, err := exec.Command("/tmp/registry/init.sh").CombinedOutput()
	assert.NoError(t, err)
	log.Logger.Info(string(output))
}
