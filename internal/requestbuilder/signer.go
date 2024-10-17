package requestbuilder

import (
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"strings"
)

type Signer struct {
	key []byte
}

func (p *Signer) Init(key string) *Signer {
	p.key = []byte(key)
	return p
}

func (p *Signer) Sign(method string, host string, path string, parameters string) (string, error) {
	if method == "" || host == "" || path == "" || parameters == "" {
		return "", nil
	}

	var sb strings.Builder
	sb.WriteString(method)
	sb.WriteString("\n")
	sb.WriteString(host)
	sb.WriteString("\n")
	sb.WriteString(path)
	sb.WriteString("\n")
	sb.WriteString(parameters)

	return p.sign(sb.String()), nil
}

func (p *Signer) sign(payload string) string {
	hash := hmac.New(sha256.New, p.key)
	hash.Write([]byte(payload))
	result := base64.StdEncoding.EncodeToString(hash.Sum(nil))
	return result
}

// Ed25519Signer 结构体，继承自 Signer
type Ed25519Signer struct {
	Signer     // 嵌入 Signer
	privateKey ed25519.PrivateKey
	publicKey  ed25519.PublicKey
}

// Init initializes the Ed25519 signer with a given private key.
func (s *Ed25519Signer) Init(privateKeyBase64 string) (*Ed25519Signer, error) {
	// 解析 PEM 编码的私钥
	// print(len(privateKeyBase64))
	block, _ := pem.Decode([]byte(privateKeyBase64))
	// print(len(block.Bytes))
	if block == nil || block.Type != "PRIVATE KEY" {
		return nil, errors.New("failed to decode PEM block containing the private key")
	}

	// 解析 x509 格式的私钥
	parsedKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	// 类型断言为 ed25519.PrivateKey
	var ok bool
	s.privateKey, ok = parsedKey.(ed25519.PrivateKey)
	if !ok {
		return nil, errors.New("not an ed25519 private key")
	}

	// 提取公钥
	s.publicKey = s.privateKey.Public().(ed25519.PublicKey)

	return s, nil
}

// 重写 Sign 方法，使用 Ed25519 签名
func (s *Ed25519Signer) Sign(method string, host string, path string, parameters string) (string, error) {
	if method == "" || host == "" || path == "" {
		return "", errors.New("method, host, and path cannot be empty")
	}

	var sb strings.Builder
	sb.WriteString(method)
	sb.WriteString("\n")
	sb.WriteString(host)
	sb.WriteString("\n")
	sb.WriteString(path)
	sb.WriteString("\n")

	// 拼接参数
	sb.WriteString(parameters)

	// 生成签名
	payload := sb.String()
	signature := ed25519.Sign(s.privateKey, []byte(payload))

	// 进行 Base64 编码
	return base64.StdEncoding.EncodeToString(signature), nil
}
