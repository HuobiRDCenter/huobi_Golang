package model

import "encoding/json"

type WebSocketV2AuthenticationResponse struct {
	Action string `json:"action"`
	Ch     string `json:"ch"`
	Code   int    `json:"code"`
	Data   interface{}
}

func (p *WebSocketV2AuthenticationResponse) IsAuth() bool {
	return p.Action == "req" && p.Ch == "auth" && p.Code == 200
}

func ParseWSV2AuthResp(message string) *WebSocketV2AuthenticationResponse {
	result := &WebSocketV2AuthenticationResponse{}
	err := json.Unmarshal([]byte(message), result)
	if err != nil {
		return nil
	}

	return result
}
