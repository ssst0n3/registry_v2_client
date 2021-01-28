package entity

import (
	"bytes"
	"encoding/json"
	"github.com/containerd/containerd/images"
	"github.com/docker/distribution/manifest/schema2"
	"github.com/ssst0n3/awesome_libs"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"io"
	"net/http"
)

type Manifest struct {
	Name          string           `json:"name"`
	Reference     string           `json:"reference"`
	Method        string           `json:"method"`
	Manifest      schema2.Manifest `json:"manifest"`
	Tampering     bool             `json:"tampering"`
	TamperingBody io.Reader
}

func NewManifest(repositoryName, reference, method string, manifest schema2.Manifest) Manifest {
	return Manifest{
		Name:      repositoryName,
		Reference: reference,
		Method:    method,
		Manifest:  manifest,
	}
}

func NewManifestForTampering(repositoryName, reference string, tamperingBody io.Reader) Manifest {
	return Manifest{
		Name:          repositoryName,
		Reference:     reference,
		Method:        http.MethodPut,
		Tampering:     true,
		TamperingBody: tamperingBody,
	}
}

func (e Manifest) Uri() (uri string) {
	return awesome_libs.Format("/v2/{.name}/manifests/{.reference}", awesome_libs.Dict{
		"name":      e.Name,
		"reference": e.Reference,
	})
}

func (e Manifest) GetMethod() (method string) {
	return e.Method
}

func (e Manifest) GetQuery() (query interface{}) {
	return
}

func (e Manifest) GetBody() (body io.Reader, err error) {
	if e.Tampering {
		body = e.TamperingBody
		return
	}
	switch e.Method {
	case http.MethodPut:
		m, err := json.Marshal(e.Manifest)
		if err != nil {
			awesome_error.CheckErr(err)
			return nil, err
		}
		body = bytes.NewReader(m)
	}
	return
}

func (e Manifest) GetHeader() (header map[string]string) {
	switch e.Method {
	case http.MethodGet:
		header = map[string]string{
			// TODO: as choice
			"Accept": "application/vnd.docker.distribution.manifest.v2+json, application/vnd.docker.distribution.manifest.list.v2+json, application/vnd.oci.image.manifest.v1+json, application/vnd.oci.image.index.v1+json, */*",
		}
	case http.MethodPut:
		header = map[string]string{
			"Content-Type": images.MediaTypeDockerSchema2Manifest,
		}
	}
	return
}

func (e Manifest) CheckRedirect() (f func(req *http.Request, via []*http.Request) (err error)) {
	return
}
