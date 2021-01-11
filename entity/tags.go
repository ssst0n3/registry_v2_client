package entity

import (
	"github.com/ssst0n3/awesome_libs"
	"io"
	"net/http"
)

type TagsPaginatedQuery struct {
	Number int `json:"n,string"`
	Last   int `json:"last,string"`
}

type Tags struct {
	RepositoryName string             `json:"repository_name"`
	Method         string             `json:"method"`
	Query          TagsPaginatedQuery `json:"query"`
}

func NewTags(repositoryName string, method string, query TagsPaginatedQuery) Tags {
	return Tags{
		RepositoryName: repositoryName,
		Method:         method,
		Query:          query,
	}
}

func (e Tags) Uri() (uri string) {
	return awesome_libs.Format("/v2/{.name}/tags/list", awesome_libs.Dict{
		"name": e.RepositoryName,
	})
}

func (e Tags) GetMethod() (method string) {
	return e.Method
}

func (e Tags) GetQuery() (query interface{}) {
	return e.Query
}

func (e Tags) GetBody() (body io.Reader, err error) {
	return
}

func (e Tags) GetHeader() (header map[string]string) {
	return
}

func (e Tags) CheckRedirect() (f func(req *http.Request, via []*http.Request) (err error)) {
	return
}