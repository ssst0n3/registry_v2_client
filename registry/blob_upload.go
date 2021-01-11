package registry

import (
	"github.com/ssst0n3/registry_v2_client/entity"
	"net/http"
)

func (r Registry) GetBlobUpload(repositoryName, uuid string) (uploadRange string, err error) {
	e := entity.NewBlobUpload(http.MethodGet, repositoryName, uuid, nil, false, 0, 0, entity.BlobUploadQuery{})
	resp, err := r.NotCareBody(e)
	if err != nil {
		return
	}
	uploadRange = resp.Header.Get("Range")
	return
}

func (r Registry) PatchBlobUpload(repositoryName, uuid string, binary []byte, chunked bool, start, end int) (location, uploadRange string) {
	e := entity.NewBlobUpload(http.MethodPatch, repositoryName, uuid, binary, chunked, start, end, entity.BlobUploadQuery{})
	resp, err := r.NotCareBody(e)
	if err != nil {
		return
	}
	location = resp.Header.Get("Location")
	uploadRange = resp.Header.Get("Range")
	return
}

func (r Registry) PutBlobUpload(repositoryName, uuid, digest string, binary []byte) (location, contentRange string) {
	e := entity.NewBlobUpload(http.MethodPut, repositoryName, uuid, binary, false, 0, 0, entity.BlobUploadQuery{Digest: digest})
	resp, err := r.NotCareBody(e)
	if err != nil {
		return
	}
	location = resp.Header.Get("Location")
	contentRange = resp.Header.Get("Content-Range")
	return
}

func (r Registry) DeleteBlobUpload(repositoryName, uuid string) (err error) {
	e := entity.NewBlobUpload(http.MethodDelete, repositoryName, uuid, nil, false, 0, 0, entity.BlobUploadQuery{})
	_, err = r.NotCareBody(e)
	return
}
