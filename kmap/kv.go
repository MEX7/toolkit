package kmap

import (
	"strconv"
)

type Kv struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func Kv2MapStr(in []Kv) (out map[string]string) {
	out = make(map[string]string)
	for _, item := range in {
		out[item.Key] = item.Value
	}
	return
}

func Kv2MapInt(in []Kv) (out map[int]int) {
	out = make(map[int]int)
	for _, item := range in {
		valInt, _ := strconv.Atoi(item.Value)
		keyInt, _ := strconv.Atoi(item.Key)
		out[keyInt] = valInt
	}
	return
}
