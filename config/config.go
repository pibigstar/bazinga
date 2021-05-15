package config

import (
	_ "embed"
	"encoding/json"
)

//go:embed zh-CN.json
var codeMsgBs []byte

var CodeMsg map[string]string

func init() {
	err := json.Unmarshal(codeMsgBs, &CodeMsg)
	if err != nil {
		panic(err)
	}
}
