package registry

import (
	"encoding/json"
	"errors"
	"github.com/docker/distribution/manifest/schema2"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/ssst0n3/registry_v2_client/entity"
	"io"
	"net/http"
)

func (r *Registry) GetManifest(repositoryName, ref string) (manifest schema2.Manifest, err error) {
	e := entity.NewManifest(repositoryName, ref, http.MethodGet, schema2.Manifest{})
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

func (r *Registry) TamperingManifest(repositoryName, reference string, body io.Reader) (err error) {
	e := entity.NewManifestForTampering(repositoryName, reference, body)
	_, err = r.NotCareBody(e)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}

func (r *Registry) PutManifest(repositoryName, reference string, manifest schema2.Manifest) (err error) {
	e := entity.NewManifest(repositoryName, reference, http.MethodPut, manifest)
	resp, body, err := r.AutoReadBody(e)
	//_, err = r.NotCareBody(e)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	log.Logger.Debug(string(body))
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		err = errors.New(string(body))
		return
	}
	return
}

func (r *Registry) DeleteManifest(repositoryName, reference string) (err error) {
	e := entity.NewManifest(repositoryName, reference, http.MethodDelete, schema2.Manifest{})
	_, err = r.NotCareBody(e)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}
