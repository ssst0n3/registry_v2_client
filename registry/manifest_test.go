package registry

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/docker/distribution/manifest/schema2"
	"github.com/sirupsen/logrus"
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/ssst0n3/registry_v2_client/test/test_data"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDockerHub(t *testing.T) {
	log.Logger.Level = logrus.DebugLevel
	r := NewRegistry("index.docker.io", "", "", false)
	repositoryName := "library/hello-world"
	m, err := r.GetManifest(repositoryName, "latest")
	assert.NoError(t, err)
	spew.Dump(m)
}

func TestRegistryForTest(t *testing.T) {
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
		InitRegistryForTest(t)
		reference := "v2"
		err := r.PutManifest(repositoryName, reference, test_data.ManifestHelloWorld)
		assert.NoError(t, err)
		m, err := r.GetManifest(repositoryName, reference)
		assert.NoError(t, err)
		assert.Equal(t, test_data.ManifestHelloWorld, m)
		InitRegistryForTest(t)
	})

	t.Run("delete", func(t *testing.T) {
		reference := test_data.ManifestHelloWorldDigest
		assert.NoError(t, r.DeleteManifest(repositoryName, reference))
		m, err := r.GetManifest(repositoryName, "v1")
		assert.NoError(t, err)
		assert.Equal(t, schema2.Manifest{}, m)
		InitRegistryForTest(t)
	})
}
