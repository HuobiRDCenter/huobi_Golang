package requestbuilder

import (
	"sync"
	"testing"
)

func TestSigner_Sign_FourString_Success(t *testing.T) {
	signer := new(Signer).Init("secret")

	result, err := signer.Sign("GET", "api.huobi.pro", "/v1/account/history", "account-id=1&currency=btcusdt")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "HUP3n78npIuTzVKyjEOrPictRKEUTRoYs7Ld5y38hmA="
	if result != expected {
		t.Errorf("expected: %s, actual: %s", expected, result)
	}
}

func TestSigner_Sign_RunTwice_GetSameResult(t *testing.T) {
	signer := new(Signer).Init("secret")

	result1, err1 := signer.Sign("GET", "api.huobi.pro", "/v1/account/history", "account-id=1&currency=btcusdt")
	if err1 != nil {
		t.Fatalf("unexpected error: %v", err1)
	}

	result2, err2 := signer.Sign("GET", "api.huobi.pro", "/v1/account/history", "account-id=1&currency=btcusdt")
	if err2 != nil {
		t.Fatalf("unexpected error: %v", err2)
	}

	if result1 != result2 {
		t.Errorf("expected: %s, actual: %s", result1, result2)
	}
}

func TestSigner_Sign_OneEmptyString_ReturnEmpty(t *testing.T) {
	signer := new(Signer)

	result, err := signer.Sign("GET", "api.huobi.pro", "", "account-id=1&currency=btcusdt")
	if err != nil {
		t.Errorf("expected an error for empty path, but got none")
	}

	expected := ""
	if result != expected {
		t.Errorf("expected: %s, actual: %s", expected, result)
	}
}

func TestSigner_Sign_RaceCondition(t *testing.T) {
	signer := new(Signer).Init("secret")

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
			signer.Sign("GET", "api.huobi.pro", "/v1/account/history", "account-id=1&currency=btcusdt")
		}()
	}

	wg.Wait()
}
