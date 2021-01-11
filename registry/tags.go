package registry

import (
	"encoding/json"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/registry_v2_client/entity"
	"github.com/ssst0n3/registry_v2_client/response"
	"net/http"
)

func (r Registry) getTagsBase(repositoryName string, query entity.TagsPaginatedQuery) (tags []string, err error) {
	e := entity.NewTags(repositoryName, http.MethodGet, query)
	_, body, err := r.AutoReadBody(e)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	var resp response.TagsList
	err = json.Unmarshal(body, &resp)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	tags = resp.Tags
	return

}

func (r Registry) GetTags(repositoryName string) (tags []string, err error) {
	return r.getTagsBase(repositoryName, entity.TagsPaginatedQuery{})
}

func (r Registry) TagsPaginated(repositoryName string, number, last int) (tags []string, err error) {
	return r.getTagsBase(repositoryName, entity.TagsPaginatedQuery{
		Number: number,
		Last:   last,
	})
}
