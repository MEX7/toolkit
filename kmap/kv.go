package kmap

import (
	"fmt"
	"strconv"
	"strings"
)

type Kv struct {
	Key   interface{} `json:"key"`
	Value interface{} `json:"value"`
}

func Kv2MapStr(in []Kv) (out map[string]string) {
	out = make(map[string]string)
	for _, item := range in {
		switch item.Key.(type) {
		case string:
			out[item.Key.(string)] = item.Value.(string)
		}
	}
	return
}

func Kv2MapInt(in []Kv) (out map[int]int) {
	out = make(map[int]int)
	for _, item := range in {
		switch item.Key.(type) {
		case string:
			valInt, _ := strconv.Atoi(item.Value.(string))
			keyInt, _ := strconv.Atoi(item.Key.(string))
			out[keyInt] = valInt
		case int:
			out[item.Key.(int)] = item.Value.(int)
		}
	}
	return
}

// KvFormat TCP:9001,TCP:9003
func KvFormat(in []Kv) string {
	var out string
	for _, item := range in {
		var k, v string
		switch item.Key.(type) {
		case string:
			k = item.Key.(string)
			v = item.Value.(string)
		}
		if out == "" {
			out = fmt.Sprintf("%s,%s", k, v)
		} else {
			out = fmt.Sprintf("%s:%s:%s", out, k, v)
		}
	}
	return out
}

// KvUnFormat TCP:9001,TCP:9003
func KvUnFormat(in string) []Kv {
	var (
		res = make([]Kv, 0)
	)
	for _, item := range strings.Split(in, ",") {
		schemePort := strings.Split(item, ":")
		if len(schemePort) != 2 {
			continue
		}
		res = append(res, Kv{
			Key:   schemePort[0],
			Value: schemePort[1],
		})

	}
	return res
}

func MapInt2Kv(in map[int]int) []Kv {
	res := make([]Kv, 0)
	for k, v := range in {
		res = append(res, Kv{
			Key:   k,
			Value: v,
		})
	}
	return res
}

func MapString2Kv(in map[string]string) []Kv {
	res := make([]Kv, 0)
	for k, v := range in {
		res = append(res, Kv{
			Key:   k,
			Value: v,
		})
	}
	return res
}
