package registry

import (
	"github.com/sirupsen/logrus"
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/ssst0n3/registry_v2_client/test/test_data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegistry(t *testing.T) {
	log.Logger.Level = logrus.DebugLevel
	r := GetRegistryForTest()
	repositoryName := "dkdk/hello-world"

	t.Run("get manifest", func(t *testing.T) {
		reference := "v1"
		m, err := r.GetManifest(repositoryName, reference)
		assert.NoError(t, err)
		assert.Equal(t, test_data.ManifestHelloWorld, m)
	})

	t.Run("put", func(t *testing.T) {
		reference := "v2"
		err := r.PutManifest(repositoryName, reference, test_data.ManifestHelloWorld)
		assert.NoError(t, err)
		m, err := r.GetManifest(repositoryName, reference)
		assert.NoError(t, err)
		assert.Equal(t, test_data.ManifestHelloWorld, m)
	})

	t.Run("delete", func(t *testing.T) {
		reference := "sha256:deea9d5fd7a5915c08d0b988e5b5bfed0d5ea9a632ddf70c9a79d08764b104fa"
		assert.NoError(t, r.DeleteManifest(repositoryName, reference))
		_, err := r.GetManifest(repositoryName, "v2")
		assert.Error(t, err)
	})
}
