package requestbuilder

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

type Signer struct {
	key []byte
}

func (p *Signer) Init(key string) *Signer {
	p.key = []byte(key)
	return p
}

func (p *Signer) Sign(method string, host string, path string, parameters string) string {
	if method == "" || host == "" || path == "" || parameters == "" {
		return ""
	}

	var sb strings.Builder
	sb.WriteString(method)
	sb.WriteString("\n")
	sb.WriteString(host)
	sb.WriteString("\n")
	sb.WriteString(path)
	sb.WriteString("\n")
	sb.WriteString(parameters)

	return p.sign(sb.String())
}

func (p *Signer) sign(payload string) string {
	hash := hmac.New(sha256.New, p.key)
	hash.Write([]byte(payload))
	result := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return result
}
