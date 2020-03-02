package requestbuilder

import (
	"../../pkg/getrequest"
	"fmt"
)

type PublicUrlBuilder struct {
	host string
}

func (p *PublicUrlBuilder) Init(host string) *PublicUrlBuilder {
	p.host = host
	return p
}

func (p *PublicUrlBuilder) Build(path string, request *getrequest.GetRequest) string {
	if request != nil {
		result := fmt.Sprintf("https://%s%s?%s", p.host, path, request.BuildParams())
		return result
	} else {
		result := fmt.Sprintf("https://%s%s", p.host, path)
		return result
	}
}