package registry

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/registry_v2_client/entity"
	"io/ioutil"
	"net/http"
)

func (r *Registry) GetBlob(repositoryName, dgs string) (exists bool, err error) {
	e := entity.NewBlob(http.MethodHead, repositoryName, dgs, false, 0, 0, true)
	resp, err := r.NotCareBody(e)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	if resp.StatusCode <= 399 {
		exists = true
	}
	return
}

func (r *Registry) fetchBlobCommon(e entity.Entity) (content []byte, err error) {
	resp, err := r.Do(e)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 302 {
		content = []byte(resp.Header.Get("Location"))
	} else if resp.StatusCode < 300 {
		content, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			awesome_error.CheckErr(err)
			return
		}
	}
	return
}

func (r *Registry) FetchBlob(repositoryName, dgs string, followRedirects bool) (content []byte, err error) {
	e := entity.NewBlob(http.MethodGet, repositoryName, dgs, false, 0, 0, followRedirects)
	return r.fetchBlobCommon(e)
}

func (r *Registry) FetchBlobPart(repositoryName, dgs string, start, end int, followRedirects bool) (content []byte, err error) {
	e := entity.NewBlob(http.MethodGet, repositoryName, dgs, true, start, end, followRedirects)
	return r.fetchBlobCommon(e)
}

func (r *Registry) DeleteBlob(repositoryName, dgs string) (err error) {
	e := entity.NewBlob(http.MethodDelete, repositoryName, dgs, false, 0, 0, true)
	_, err = r.NotCareBody(e)
	if err != nil {
		return
	}
	return
}
