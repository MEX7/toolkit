package nws

import (
	"errors"
	"fmt"
	"regexp"
)

func EncodeNodeSocketIO(code int, b []byte) []byte {
	return EncodeNodeSocketIOEmit("message", code, b)
}

func EncodeNodeSocketIOEmit(typ string, code int, b []byte) []byte {
	out := fmt.Sprintf(`["%s",%s]`, typ, string(b))
	return []byte(fmt.Sprintf("%d%s", code, out))
}

var reg = regexp.MustCompile(`[0-9]*?[a-z]*?\s*,?({.*})]?`)

func DecodeNodeSocketIO(in []byte) (content string, typ string, err error) {
	fss := reg.FindStringSubmatch(string(in))
	if len(fss) != 2 {
		return "", "", errors.New("rule mismatch")
	}
	var res CompatMsg
	err = res.UnmarshalJSON([]byte(fss[1]))
	// err = json.Unmarshal([]byte(fss[1]), &res)
	if err != nil {
		return
	}
	return fss[1], res.Type, nil
}
