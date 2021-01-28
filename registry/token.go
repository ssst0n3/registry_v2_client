package registry

import (
	"github.com/docker/distribution/context"
	"github.com/docker/docker/api/types"
	"github.com/genuinetools/reg/registry"
	"github.com/ssst0n3/awesome_libs/awesome_error"
)

func (r *Registry) GetTokenHeader(url string) (header map[string]string, err error) {
	// use reg for temporary using
	cred := types.AuthConfig{
		Username:      r.Username,
		ServerAddress: r.ServiceAddress,
	}
	if r.Username == "" {
		cred.IdentityToken = r.Password
	} else {
		cred.Password = r.Password
	}
	reg, err := registry.New(context.Background(), cred, registry.Opt{
		Domain:   "",
		Insecure: r.Insecure,
		Debug:    true,
		SkipPing: true,
		NonSSL:   r.Insecure,
		Timeout:  0,
		Headers:  nil,
	})
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	header, err = reg.Headers(context.Background(), url)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	r.TokenHeader = header
	return
}
