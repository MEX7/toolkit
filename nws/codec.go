package nws

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
)

func EncodeNodeSocketIO(b []byte) []byte {
	return EncodeNodeSocketIOEmit("message", b)
}

func EncodeNodeSocketIOEmit(typ string, b []byte) []byte {
	return []byte(fmt.Sprintf(`["%s", %s]`, typ, string(b)))
}

var reg = regexp.MustCompile(`[0-9]*?[a-z]*?\s*,?({.*})]?`)

func DecodeNodeSocketIO(in []byte) (content string, typ string, err error) {
	fss := reg.FindStringSubmatch(string(in))
	if len(fss) != 2 {
		return "", "", errors.New("rule mismatch")
	}
	var res CompatMsg
	err = json.Unmarshal([]byte(fss[1]), &res)
	if err != nil {
		return
	}
	return fss[1], res.Type, nil
}
