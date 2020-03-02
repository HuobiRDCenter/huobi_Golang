package getrequest

import (
	"net/url"
)

type GetRequest struct {
	urls url.Values
}

func (p *GetRequest) Init() *GetRequest {
	p.urls = url.Values{}
	return p
}

func (p *GetRequest) InitFrom(reqParams *GetRequest) *GetRequest {
	if reqParams != nil {
		p.urls = reqParams.urls
	} else {
		p.urls = url.Values{}
	}
	return p
}

func (p *GetRequest) AddParam(property string, value string) *GetRequest {
	if property != "" && value != "" {
		p.urls.Add(property, value)
	}
	return p
}

func (p *GetRequest) BuildParams() string {
	return p.urls.Encode()
}