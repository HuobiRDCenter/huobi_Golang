package auth

import "testing"

func TestIsAuth_OpIsAuthCodeIs0_ReturnTrue(t *testing.T) {
	response := WebSocketV1AuthenticationResponse{}
	response.Op = "auth"
	response.ErrorCode = 0

	result := response.IsAuth()

	expected := true
	if result != expected {
		t.Errorf("expected: %v, actual: %v", expected, result)
	}
}

func TestIsAuth_OpIsAuthCodeIs4002_ReturnTrue(t *testing.T) {
	response := WebSocketV1AuthenticationResponse{}
	response.Op = "auth"
	response.ErrorCode = 4002

	result := response.IsAuth()

	expected := true
	if result != expected {
		t.Errorf("expected: %v, actual: %v", expected, result)
	}
}

func TestIsAuth_OpIsSub_ReturnFalse(t *testing.T) {
	response := WebSocketV1AuthenticationResponse{}
	response.Op = "sub"

	result := response.IsAuth()

	expected := false
	if result != expected {
		t.Errorf("expected: %v, actual: %v", expected, result)
	}
}