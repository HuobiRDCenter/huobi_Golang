package requestbuilder

import (
	"time"

	"github.com/huobirdcenter/huobi_golang/internal/model"
	model2 "github.com/huobirdcenter/huobi_golang/pkg/model"
)

type WebSocketV1RequestBuilder struct {
	akKey   string
	akValue string
	smKey   string
	smValue string
	svKey   string
	svValue string
	tKey    string
	tValue  string

	host string
	path string
	sign string

	signer SignerInterface
}

func (p *WebSocketV1RequestBuilder) Init(accessKey string, secretKey string, host string, path string, sign string) *WebSocketV1RequestBuilder {
	p.akKey = "AccessKeyId"
	p.akValue = accessKey
	p.smKey = "SignatureMethod"
	p.smValue = "HmacSHA256"
	p.svKey = "SignatureVersion"
	p.svValue = "2"
	p.tKey = "Timestamp"

	p.host = host
	p.path = path
	p.sign = sign

	if sign == "256" {
		p.signer = new(Signer).Init(secretKey) // Signer 实现了接口
	} else {
		// 使用 Ed25519 签名
		edSigner := new(Ed25519Signer)
		var err error
		edSigner, err = edSigner.Init(secretKey)
		if err != nil {
			// 处理错误
			return nil // 假设这是在一个返回 error 的函数中
		}
		p.signer = edSigner // Ed25519Signer 也实现了接口
	}

	return p
}

func (p *WebSocketV1RequestBuilder) Build() (string, error) {
	time := time.Now().UTC()
	return p.build(time)
}

func (p *WebSocketV1RequestBuilder) build(utcDate time.Time) (string, error) {
	time := utcDate.Format("2006-01-02T15:04:05")

	req := new(model2.GetRequest).Init()
	req.AddParam(p.akKey, p.akValue)
	req.AddParam(p.smKey, p.smValue)
	req.AddParam(p.svKey, p.svValue)
	req.AddParam(p.tKey, time)

	signature, err := p.signer.Sign("GET", p.host, p.path, req.BuildParams())

	auth := new(model.WebSocketV1AuthenticationRequest).Init()
	auth.AccessKeyId = p.akValue
	auth.Timestamp = time
	auth.Signature = signature

	result, err := model2.ToJson(auth)
	return result, err
}
