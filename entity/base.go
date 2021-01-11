package entity

import (
	"io"
	"net/http"
)

type Base struct {
}

func NewBase() Base {
	return Base{}
}

func (e Base) Uri() (uri string) {
	return "/v2/"
}

func (e Base) GetMethod() (method string) {
	return http.MethodGet
}

func (e Base) GetQuery() (query interface{}) {
	return
}

func (e Base) GetBody() (body io.Reader, err error) {
	return
}

func (e Base) GetHeader() (header map[string]string) {
	return
}

func (e Base) CheckRedirect() (f func(req *http.Request, via []*http.Request) (err error)) {
	return
}
