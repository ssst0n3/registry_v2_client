package registry

import (
	"github.com/ssst0n3/awesome_libs"
	"github.com/ssst0n3/awesome_libs/awesome_error"
	"github.com/ssst0n3/awesome_libs/awesome_structs"
	"github.com/ssst0n3/awesome_libs/log"
	"github.com/ssst0n3/registry_v2_client/entity"
	http2 "github.com/ssst0n3/registry_v2_client/http"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
)

type Registry struct {
	ServiceAddress string            `json:"service_address"`
	Insecure       bool              `json:"insecure"`
	Username       string            `json:"username"`
	Password       string            `json:"password"`
	TokenHeader    map[string]string `json:"token_header"`
	Client         *http.Client
}

func NewRegistry(serviceAddress, username, password string, insecure bool) Registry {
	return Registry{
		ServiceAddress: serviceAddress,
		Insecure:       insecure,
		Username:       username,
		Password:       password,
		Client:         http2.NewClient(username, password),
	}
}

func (r *Registry) Url(e entity.Entity) string {
	protocol := "https"
	if r.Insecure {
		protocol = "http"
	}
	return awesome_libs.Format("{.protocol}://{.registry}{.uri}", awesome_libs.Dict{
		"protocol": protocol,
		"registry": r.ServiceAddress,
		"uri":      e.Uri(),
	})
}

func (r *Registry) GetBody(e entity.Entity) (reader io.Reader, err error) {
	reader, err = e.GetBody()
	if err != nil {
		return
	}
	return
}

func (r *Registry) Do(e entity.Entity) (resp *http.Response, err error) {
	url := r.Url(e)
	log.Logger.Debugf("url: %s", url)

	reader, err := r.GetBody(e)
	if err != nil {
		return
	}
	req, err := http.NewRequest(e.GetMethod(), url, reader)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	query := req.URL.Query()
	q, err := awesome_structs.StringMap(e.GetQuery())
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	for k, v := range q {
		query.Add(k, v)
	}
	req.URL.RawQuery = query.Encode()
	//token := fmt.Sprintf("Basic %s", base64.StdEncoding.EncodeToString([]byte(r.Username+":"+r.Password)))
	//req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	for k, v := range e.GetHeader() {
		req.Header.Add(k, v)
	}
	if len(r.TokenHeader) == 0 {
		//req.SetBasicAuth(r.Username, r.Password)
		//_, err := r.GetTokenHeader(req.URL.String())
		//if err != nil {
		//	return resp, err
		//}
	}
	for k, v := range r.TokenHeader {
		req.Header.Add(k, v)
	}

	requestDump, err := httputil.DumpRequest(req, true)
	log.Logger.Debugf("\n%s\n\n", requestDump)

	if e.CheckRedirect() != nil {
		r.Client.CheckRedirect = e.CheckRedirect()
	}
	resp, err = r.Client.Do(req)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}

	responseDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		awesome_error.CheckErr(err)
		return
	}
	log.Logger.Debugf("\n%s\n\n", responseDump)

	return
}

func (r *Registry) NotCareBody(e entity.Entity) (resp *http.Response, err error) {
	resp, err = r.Do(e)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	return
}

func (r *Registry) AutoReadBody(e entity.Entity) (resp *http.Response, body []byte, err error) {
	resp, err = r.Do(e)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	return
}

// TODO: auto parse response status_code
