package requestbuilder

import (
	"fmt"
	"net/url"
	"time"

	"github.com/huobirdcenter/huobi_golang/pkg/model"
)

type SignerInterface interface {
	Sign(method string, host string, path string, parameters string) (string, error)
}

type PrivateUrlBuilder struct {
	host     string
	akKey    string
	akValue  string
	smKey    string
	smValue  string
	smValue2 string
	svKey    string
	svValue  string
	tKey     string
	sign     string
	signer   interface{}
}

func (p *PrivateUrlBuilder) Init(accessKey string, secretKey string, host string, sign string) *PrivateUrlBuilder {
	p.akKey = "AccessKeyId"
	p.akValue = accessKey
	p.smKey = "SignatureMethod"
	p.smValue = "HmacSHA256"
	// p.smValue2 = "ED25519"
	p.svKey = "SignatureVersion"
	p.svValue = "2"
	p.tKey = "Timestamp"
	p.sign = sign
	p.host = host
	if sign == "256" {
		p.signer = new(Signer).Init(secretKey) // HMAC SHA256 签名
	} else {
		// 使用 Ed25519 签名
		edSigner, err := new(Ed25519Signer).Init(secretKey)
		if err != nil {
			// 处理错误，返回 nil 或其他适当的错误处理
			return nil
		}
		p.smValue = "Ed25519" // 更新签名方法为 ED25519
		p.signer = edSigner   // 使用 Ed25519Signer
	}

	return p
}

func (p *PrivateUrlBuilder) Build(method string, path string, request *model.GetRequest) string {
	time := time.Now().UTC()

	return p.BuildWithTime(method, path, time, request)
}

func (p *PrivateUrlBuilder) BuildWithTime(method string, path string, utcDate time.Time, request *model.GetRequest) string {
	time := utcDate.Format("2006-01-02T15:04:05")

	req := new(model.GetRequest).InitFrom(request)
	req.AddParam(p.akKey, p.akValue)
	req.AddParam(p.tKey, time)
	req.AddParam(p.svKey, p.svValue)
	
	req.AddParam(p.smKey, p.smValue)
	parameters := req.BuildParams()

	// 使用类型断言来调用 Sign 方法
	var signature string
	var err error

	switch signer := p.signer.(type) {
	case *Signer:
		signature, err = signer.Sign(method, p.host, path, parameters)
	case *Ed25519Signer:
		signature, err = signer.Sign(method, p.host, path, parameters)
	default:
		// 处理未知签名器的情况
		return "" // 或者返回适当的错误
	}

	if err != nil {
		// 处理签名错误
		return "" // 或者返回一个适当的错误
	}

	url := fmt.Sprintf("https://%s%s?%s&Signature=%s", p.host, path, parameters, url.QueryEscape(signature))
	return url
}
