package entity

import (
	"bytes"
	"github.com/ssst0n3/awesome_libs"
	"io"
	"net/http"
	"strconv"
)

const (
	HeaderDockerUploadUUID = "Docker-Upload-UUID"
)

type InitiateBlobUploadQuery struct {
	Digest string `json:"digest"`
}

type MountBlobQuery struct {
	Mount string `json:"mount"`
	From  string `json:"from"`
}

type InitiateBlobUpload struct {
	RepositoryName          string                  `json:"repository_name"`
	InitiateBlobUploadQuery InitiateBlobUploadQuery `json:"query"`
	MountBlobQuery          MountBlobQuery          `json:"mount_blob_query"`
	Binary                  []byte
}

func NewInitiateBlobUpload(repositoryName string, binary []byte, initiateBlobUploadQuery InitiateBlobUploadQuery, mountBlobQuery MountBlobQuery) InitiateBlobUpload {
	return InitiateBlobUpload{
		RepositoryName:          repositoryName,
		InitiateBlobUploadQuery: initiateBlobUploadQuery,
		MountBlobQuery:          mountBlobQuery,
		Binary:                  binary,
	}
}

func (e InitiateBlobUpload) Uri() (uri string) {
	return awesome_libs.Format("/v2/{.name}/blobs/uploads/", awesome_libs.Dict{
		"name": e.RepositoryName,
	})
}

func (e InitiateBlobUpload) GetMethod() (method string) {
	return http.MethodPost
}

func (e InitiateBlobUpload) GetQuery() (query interface{}) {
	if len(e.InitiateBlobUploadQuery.Digest) > 0 {
		query = e.InitiateBlobUploadQuery
	}
	return
}

func (e InitiateBlobUpload) GetBody() (body io.Reader, err error) {
	body = bytes.NewReader(e.Binary)
	return
}

func (e InitiateBlobUpload) GetHeader() (header map[string]string) {
	header = map[string]string{
		"Content-Length": strconv.Itoa(len(e.Binary)),
		"Content-Type":   "application/octect-stream",
	}
	return
}

func (e InitiateBlobUpload) CheckRedirect() (f func(req *http.Request, via []*http.Request) (err error)) {
	return
}
