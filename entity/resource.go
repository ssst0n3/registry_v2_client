package entity

import (
	"io"
	"net/http"
)

type Entity interface {
	Uri() (url string)
	GetMethod() (method string)
	GetQuery() (query interface{})
	GetBody() (body io.Reader, err error)
	GetHeader() (header map[string]string)
	CheckRedirect() (f func(req *http.Request, via []*http.Request) (err error))
}
