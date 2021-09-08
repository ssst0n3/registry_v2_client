package http

import (
	"github.com/genuinetools/reg/registry"
	"net/http"
)

func NewClient(username, password string) *http.Client {
	tokenTransport := &registry.TokenTransport{
		Transport: http.DefaultTransport,
		Username:  username,
		Password:  password,
	}
	basicAuthTransport := &registry.BasicTransport{
		Transport: tokenTransport,
		Username:  username,
		Password:  password,
	}
	errorTransport := &registry.ErrorTransport{
		Transport: basicAuthTransport,
	}
	customTransport := &registry.CustomTransport{
		Transport: errorTransport,
	}
	return &http.Client{
		Transport: customTransport,
	}
}
