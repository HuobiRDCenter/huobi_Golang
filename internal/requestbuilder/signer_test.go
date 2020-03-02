package requestbuilder

import (
	"testing"
)

func TestSigner_Sign_FourString_Success(t *testing.T) {
	signer := new(Signer).Init("secret")

	result := signer.Sign("GET", "api.huobi.pro", "/v1/account/history", "account-id=1&currency=btcusdt")

	expected := "HUP3n78npIuTzVKyjEOrPictRKEUTRoYs7Ld5y38hmA="
	if result != expected {
		t.Errorf("expected: %s, actual: %s", expected, result)
	}
}

func TestSigner_Sign_OneEmptyString_Success(t *testing.T) {
	signer := new(Signer).Init("secret")

	result := signer.Sign("GET", "api.huobi.pro", "", "account-id=1&currency=btcusdt")

	expected := ""
	if result != expected {
		t.Errorf("expected: %s, actual: %s", expected, result)
	}
}