package entity

import (
	"bytes"
	"fmt"
	"github.com/ssst0n3/awesome_libs"
	"io"
	"net/http"
	"strconv"
)

type BlobUploadQuery struct {
	Digest string `json:"digest"`
}

type BlobUpload struct {
	Method         string          `json:"method"`
	RepositoryName string          `json:"repository_name"`
	Uuid           string          `json:"uuid"`
	Binary         []byte          `json:"binary"`
	Chunked        bool            `json:"chunked"`
	Start          int             `json:"start"`
	End            int             `json:"end"`
	Query          BlobUploadQuery `json:"query"`
}

func NewBlobUpload(method, repositoryName, uuid string, binary []byte, chunked bool, start, end int, query BlobUploadQuery) BlobUpload {
	return BlobUpload{
		Method:         method,
		RepositoryName: repositoryName,
		Uuid:           uuid,
		Binary:         binary,
		Chunked:        chunked,
		Start:          start,
		End:            end,
		Query:          query,
	}
}

func (e BlobUpload) Uri() (uri string) {
	return awesome_libs.Format("/v2/{.name}/blobs/uploads/{.uuid}", awesome_libs.Dict{
		"name": e.RepositoryName,
		"uuid": e.Uuid,
	})
}

func (e BlobUpload) GetMethod() (method string) {
	return e.Method
}

func (e BlobUpload) GetQuery() (query interface{}) {
	if len(e.Query.Digest) > 0 {
		query = e.Query
	}
	return
}

func (e BlobUpload) GetBody() (body io.Reader, err error) {
	if e.Binary != nil {
		body = bytes.NewReader(e.Binary)
	}
	return
}

func (e BlobUpload) GetHeader() (header map[string]string) {
	switch e.Method {
	case http.MethodPut:
		header = map[string]string{
			"Content-Type":   "application/octect-stream",
			"Content-Length": strconv.Itoa(len(e.Binary)),
		}
	case http.MethodPatch:
		if e.Chunked {
			header = map[string]string{
				"Content-Range":  fmt.Sprintf("%d-%d", e.Start, e.End),
				"Content-Length": strconv.Itoa(len(e.Binary)),
			}
			header["Content-Type"] = "application/octect-stream"
		}
	case http.MethodDelete:
		header = map[string]string{
			"Content-Length": "0",
		}
	}
	return
}

func (e BlobUpload) CheckRedirect() (f func(req *http.Request, via []*http.Request) (err error)) {
	return
}
