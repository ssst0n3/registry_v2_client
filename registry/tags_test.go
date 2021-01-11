package registry

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegistry_GetTags(t *testing.T) {
	r := GetRegistryForTest()
	tags, err := r.GetTags("dkdk/hello-world")
	assert.NoError(t, err)
	assert.Equal(t, true, len(tags) > 0)
}
