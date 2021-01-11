package registry

import (
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/registry_v2_client/entity"
)

func (r Registry) GetBase() (err error) {
	base := entity.NewBase()
	_, _, err = r.AutoReadBody(base)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	return
}
