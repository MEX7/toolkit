package nws

import (
	"errors"
	"fmt"
	"regexp"
)

func EncodeNodeSocketIO(fc, sc int, b []byte) []byte {
	return EncodeNodeSocketIOEmit("message", fc, sc, b)
}

func EncodeNodeSocketIOEmit(typ string, fc, sc int, b []byte) []byte {
	out := fmt.Sprintf(`["%s",%s]`, typ, string(b))
	var code string
	if sc != -1 {
		code = fmt.Sprintf("%d%d", fc, sc)
	} else {
		code = fmt.Sprintf("%d", fc)
	}
	return []byte(fmt.Sprintf("%s%s", code, out))
}

var reg = regexp.MustCompile(`[0-9]*?[a-z]*?\s*,?({.*})]?`)

func DecodeNodeSocketIO(in []byte) (content string, typ string, err error) {
	fss := reg.FindStringSubmatch(string(in))
	if len(fss) != 2 {
		return "", "", errors.New("rule mismatch")
	}
	var res CompatMsg
	err = res.UnmarshalJSON([]byte(fss[1]))
	if err != nil {
		return
	}
	return fss[1], res.Type, nil
}
