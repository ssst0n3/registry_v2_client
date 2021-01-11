package registry

import (
	"encoding/json"
	"github.com/docker/distribution/manifest/schema2"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/registry_v2_client/entity"
	"net/http"
)

func (r Registry) GetManifest(repositoryName, reference string) (manifest schema2.Manifest, err error) {
	e := entity.NewManifest(repositoryName, reference, http.MethodGet, schema2.Manifest{})
	_, body, err := r.AutoReadBody(e)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	err = json.Unmarshal(body, &manifest)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func (r Registry) PutManifest(repositoryName, reference string, manifest schema2.Manifest) (err error) {
	e := entity.NewManifest(repositoryName, reference, http.MethodPut, manifest)
	_, err = r.NotCareBody(e)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func (r Registry) DeleteManifest(repositoryName, reference string) (err error) {
	e := entity.NewManifest(repositoryName, reference, http.MethodDelete, schema2.Manifest{})
	_, err = r.NotCareBody(e)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}
