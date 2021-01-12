package entity

import (
	"fmt"
	"github.com/opencontainers/go-digest"
	"github.com/ssst0n3/awesome_libs"
	"io"
	"net/http"
)

type Blob struct {
	Method          string        `json:"method"`
	RepositoryName  string        `json:"repository_name"`
	Digest          digest.Digest `json:"digest"`
	Part            bool          `json:"part"`
	Start           int           `json:"start"`
	End             int           `json:"end"`
	FollowRedirects bool          `json:"follow_redirects"`
}

func NewBlob(method, repositoryName, dgs string, part bool, start, end int, followRedirects bool) Blob {
	return Blob{
		Method:          method,
		RepositoryName:  repositoryName,
		Digest:          digest.Digest(dgs),
		Part:            part,
		Start:           start,
		End:             end,
		FollowRedirects: followRedirects,
	}
}

func (e Blob) Uri() (uri string) {
	return awesome_libs.Format("/v2/{.name}/blobs/{.digest}", awesome_libs.Dict{
		"name":   e.RepositoryName,
		"digest": e.Digest,
	})
}

func (e Blob) GetMethod() (method string) {
	return e.Method
}

func (e Blob) GetQuery() (query interface{}) {
	return
}
func (e Blob) GetBody() (body io.Reader, err error) {
	return
}

func (e Blob) GetHeader() (header map[string]string) {
	if e.Part {
		header = map[string]string{
			"Range": fmt.Sprintf("bytes=%d-%d", e.Start, e.End),
		}
	}
	return
}

func (e Blob) CheckRedirect() (f func(req *http.Request, via []*http.Request) (err error)) {
	if e.FollowRedirects {
		return
	}
	return func(req *http.Request, via []*http.Request) (err error) {
		return http.ErrUseLastResponse
	}
}
