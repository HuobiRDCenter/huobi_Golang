package requestbuilder

import (
	"sync"
	"testing"
	// "crypto/ed25519"
	// "encoding/base64"
	// "fmt"
)

func TestEd25519Signer_Sign_FourString_Success(t *testing.T) {
	// 使用一个有效的 Ed25519 私钥进行测试，确保它是 Base64 编码的
	privateKeyBase64 := "nJanXRB7WIiSgr6eQzgNaQjp+xCkieLj0RQ2NxAZefqYiJajRf6PVrd0kE8kdHG3XuiXLlM6w5xipkdTNw5Ucw==" // 请替换为实际的 Base64 编码私钥

	// publicKey, privateKey, err := ed25519.GenerateKey(nil)
	// if err != nil {
	//     fmt.Println("failed to generate key:", err)
	//     return
	// }

	// // 编码私钥为 Base64
	// privateKeyBase64 := base64.StdEncoding.EncodeToString(privateKey)
	// publicKeyBase64 := base64.StdEncoding.EncodeToString(publicKey)

	// fmt.Println("Private Key (Base64):", privateKeyBase64)
	// fmt.Println("Public Key (Base64):", publicKeyBase64)
	signer := new(Ed25519Signer)

	_, err := signer.Init(privateKeyBase64)
	if err != nil {
		t.Fatalf("failed to initialize signer: %v", err)
	}

	result, err := signer.Sign("GET", "api.huobi.pro", "/v1/account/history", "account-id=1&currency=btcusdt")
	if err != nil {
		t.Fatalf("failed to sign: %v", err)
	}

	// 请替换为您预期的签名结果
	expected := "KnEhhsa511CrWu++7BMSdVrLtFhvxARINGsGpq8S699qMgghekPwW90RV14A4sniE3RS1J/v0hAo70iA9Ln4Aw==" // 将此替换为根据您的私钥生成的实际期望签名
	if result != expected {
		t.Errorf("expected: %s, actual: %s", expected, result)
	}
}

func TestEd25519Signer_Sign_RunTwice_GetSameResult(t *testing.T) {
	privateKeyBase64 := "nJanXRB7WIiSgr6eQzgNaQjp+xCkieLj0RQ2NxAZefqYiJajRf6PVrd0kE8kdHG3XuiXLlM6w5xipkdTNw5Ucw=="
	signer := new(Ed25519Signer)

	_, err := signer.Init(privateKeyBase64)
	if err != nil {
		t.Fatalf("failed to initialize signer: %v", err)
	}

	result1, err := signer.Sign("GET", "api.huobi.pro", "/v1/account/history", "account-id=1&currency=btcusdt")
	if err != nil {
		t.Fatalf("failed to sign: %v", err)
	}
	result2, err := signer.Sign("GET", "api.huobi.pro", "/v1/account/history", "account-id=1&currency=btcusdt")
	if err != nil {
		t.Fatalf("failed to sign: %v", err)
	}

	if result1 != result2 {
		t.Errorf("expected: %s, actual: %s", result1, result2)
	}
}

func TestEd25519Signer_Sign_OneEmptyString_ReturnEmpty(t *testing.T) {
	privateKeyBase64 := "nJanXRB7WIiSgr6eQzgNaQjp+xCkieLj0RQ2NxAZefqYiJajRf6PVrd0kE8kdHG3XuiXLlM6w5xipkdTNw5Ucw=="
	signer := new(Ed25519Signer)

	_, err := signer.Init(privateKeyBase64)
	if err != nil {
		t.Fatalf("failed to initialize signer: %v", err)
	}

	result, err := signer.Sign("GET", "api.huobi.pro", "", "account-id=1&currency=btcusdt")
	if err == nil {
		t.Errorf("expected an error for empty path, actual result: %s", result)
	} else {
		// Optional: Print the error for debugging
		t.Logf("Received expected error: %v", err)
	}
}

func TestEd25519Signer_Sign_RaceCondition(t *testing.T) {
	privateKeyBase64 := "nJanXRB7WIiSgr6eQzgNaQjp+xCkieLj0RQ2NxAZefqYiJajRf6PVrd0kE8kdHG3XuiXLlM6w5xipkdTNw5Ucw=="
	signer := new(Ed25519Signer)

	_, err := signer.Init(privateKeyBase64)
	if err != nil {
		t.Fatalf("failed to initialize signer: %v", err)
	}

	var r = 100
	wg := sync.WaitGroup{}
	wg.Add(r)
	for i := 0; i < r; i++ {
		go func() {
			defer func() {
				if r := recover(); r != nil {
					t.Errorf("race condition: panic %s", r)
				}
				wg.Done()
			}()
			_, err := signer.Sign("GET", "api.huobi.pro", "/v1/account/history", "account-id=1&currency=btcusdt")
			if err != nil {
				t.Errorf("failed to sign: %v", err)
			}
		}()
	}

	wg.Wait()
}
