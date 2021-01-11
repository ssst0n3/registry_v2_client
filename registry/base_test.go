package registry

import (
	"github.com/ssst0n3/registry_v2_client/test/test_data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func GetRegistryForTest() Registry {
	return NewRegistry(test_data.LOCAL_TEST_SERVER, "", "", true)
}

func TestRegistry_GetBase(t *testing.T) {
	r := GetRegistryForTest()
	assert.NoError(t, r.GetBase())
}
