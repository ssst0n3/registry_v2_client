package registry

import (
	"github.com/ssst0n3/registry_v2_client/entity"
)

func (r Registry) initiateBlobUploadCommon(e entity.Entity) (location, dockerUploadUUID string, err error) {
	resp, err := r.NotCareBody(e)
	if err != nil {
		return
	}
	location = resp.Header.Get("Location")
	dockerUploadUUID = resp.Header.Get(entity.HeaderDockerUploadUUID)
	return
}

func (r Registry) InitiateBlobUpload(repositoryName, dgs string, binary []byte) (location, dockerUploadUUID string, err error) {
	e := entity.NewInitiateBlobUpload(repositoryName, binary, entity.InitiateBlobUploadQuery{Digest: dgs}, entity.MountBlobQuery{})
	return r.initiateBlobUploadCommon(e)
}

func (r Registry) InitiateResumableBlobUpload(repositoryName string) (location, dockerUploadUUID string, err error) {
	e := entity.NewInitiateBlobUpload(repositoryName, []byte(""), entity.InitiateBlobUploadQuery{}, entity.MountBlobQuery{})
	return r.initiateBlobUploadCommon(e)
}

func (r Registry) MountBlob(repositoryName, mount, from string) (location, dockerUploadUUID string, err error) {
	e := entity.NewInitiateBlobUpload(repositoryName, nil, entity.InitiateBlobUploadQuery{}, entity.MountBlobQuery{
		Mount: mount,
		From:  from,
	})
	return r.initiateBlobUploadCommon(e)
}
