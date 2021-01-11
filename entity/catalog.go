package entity

import (
	"io"
	"net/http"
)

type CatalogQuery struct {
	Number int `json:"number"`
	Last   int `json:"last"`
}

type Catalog struct {
	Query CatalogQuery `json:"query"`
}

func NewCatalog(query CatalogQuery) Catalog {
	return Catalog{
		Query: query,
	}
}

func (e Catalog) Uri() (uri string) {
	return "/v2/_catalog"
}

func (e Catalog) GetMethod() (method string) {
	return http.MethodGet
}

func (e Catalog) GetQuery() (query interface{}) {
	if e.Query.Number > 0 {
		query = e.Query
	}
	return
}

func (e Catalog) GetBody() (body io.Reader, err error) {
	return
}

func (e Catalog) GetHeader() (header map[string]string) {
	return
}

func (e Catalog) CheckRedirect() (f func(req *http.Request, via []*http.Request) (err error)) {
	return
}
