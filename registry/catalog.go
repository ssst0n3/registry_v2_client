package registry

import (

	"encoding/json"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/registry_v2_client/entity"
	"github.com/ssst0n3/registry_v2_client/response"
)

func (r Registry) GetCatalog(query entity.CatalogQuery) (catalogResponse response.Catalog, err error) {
	e := entity.NewCatalog(query)
	_, body, err := r.AutoReadBody(e)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &catalogResponse)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}
