package registry

import (
	"github.com/opencontainers/go-digest"
	"github.com/sirupsen/logrus"
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/ssst0n3/registry_v2_client/test/test_data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegistry_InitiateBlobUpload(t *testing.T) {
	log.Logger.Level = logrus.DebugLevel
	r := GetRegistryForTest()
	data := []byte("test123")
	dgs := digest.FromBytes(data)
	t.Run("digest", func(t *testing.T) {
		location, uuid, err := r.InitiateBlobUpload(test_data.RepositoryName, dgs.String(), data)
		assert.NoError(t, err)
		log.Logger.Info(location)
		log.Logger.Info(uuid)
		exists, err := r.GetBlob(test_data.RepositoryName, dgs.String())
		assert.NoError(t, err)
		log.Logger.Info(exists)
	})
}
