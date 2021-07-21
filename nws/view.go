package nws

import (
	"encoding/json"
)

type CompatMsg struct {
	Type string      `json:"type"`
	Data interface{} `json:"data,omitempty" `
}

func (c CompatMsg) JSON() []byte {
	data, _ := json.Marshal(c)
	return data
}