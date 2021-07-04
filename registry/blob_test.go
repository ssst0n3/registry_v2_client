package registry

import (
	"github.com/sirupsen/logrus"
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/ssst0n3/registry_v2_client/test/test_data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegistry_GetBlob(t *testing.T) {
	log.Logger.Level = logrus.DebugLevel
	r := GetRegistryForTest()
	t.Run("exists", func(t *testing.T) {
		exists, err := r.GetBlob("dkdk/hello-world", test_data.BlobHelloWorldDigest)
		assert.NoError(t, err)
		assert.Equal(t, true, exists)
	})
	t.Run("not exists", func(t *testing.T) {
		exists, err := r.GetBlob("dkdk/hello-world", "v1")
		assert.NoError(t, err)
		assert.Equal(t, false, exists)
	})
}

func TestRegistry_FetchBlob(t *testing.T) {
	log.Logger.Level = logrus.DebugLevel
	r := GetRegistryForTest()
	body, err := r.FetchBlob("dkdk/hello-world", test_data.BlobHelloWorldDigest, true)
	assert.NoError(t, err)
	assert.Equal(t, true, len(body) > 0)
}

func TestRegistry_FetchBlobPart(t *testing.T) {
	log.Logger.Level = logrus.DebugLevel
	r := GetRegistryForTest()
	body, err := r.FetchBlobPart("dkdk/hello-world", test_data.BlobHelloWorldDigest, 0, 0, true)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(body))
}

func TestRegistry_DeleteBlob(t *testing.T) {
	log.Logger.Level = logrus.DebugLevel
	repositoryName := "dkdk/hello-world"
	r := GetRegistryForTest()
	err := r.DeleteBlob(repositoryName, test_data.BlobHelloWorldDigest)
	assert.NoError(t, err)
	exists, err := r.GetBlob(repositoryName, test_data.BlobHelloWorldDigest)
	assert.NoError(t, err)
	assert.Equal(t, false, exists)
}
