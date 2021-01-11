package registry

import (
	"errors"
	"github.com/PuerkitoBio/goquery"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/registry_v2_client/entity"
	"io/ioutil"
	"net/http"
)

func (r Registry) GetBlob(repositoryName, dgs string) (exists bool, err error) {
	e := entity.NewBlob(http.MethodHead, repositoryName, dgs, false, 0, 0)
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

func (r Registry) fetchBlobCommon(e entity.Entity) (body []byte, link string, err error) {
	resp, err := r.Do(e)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 302 {
		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			awesome_error.CheckErr(err)
			return nil, "", err
		}
		l, exists := doc.Find("a").First().Attr("href")
		if !exists {
			err = errors.New("link not exists")
			awesome_error.CheckErr(err)
			return nil, "", err
		} else {
			link = l
		}
	} else if resp.StatusCode < 300 {
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			awesome_error.CheckErr(err)
			return
		}
	}
	return
}

func (r Registry) FetchBlob(repositoryName, dgs string) (body []byte, link string, err error) {
	e := entity.NewBlob(http.MethodGet, repositoryName, dgs, false, 0, 0)
	return r.fetchBlobCommon(e)
}

func (r *Registry) FetchBlobPart(repositoryName, dgs string, start, end int) (body []byte, link string, err error) {
	e := entity.NewBlob(http.MethodGet, repositoryName, dgs, true, start, end)
	return r.fetchBlobCommon(e)
}

func (r *Registry) DeleteBlob(repositoryName, dgs string) (err error) {
	e := entity.NewBlob(http.MethodDelete, repositoryName, dgs, false, 0, 0)
	_, err = r.NotCareBody(e)
	if err != nil {
		return
	}
	return
}
