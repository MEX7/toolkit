package kmap

import (
	"strconv"
)

func MapKeyString2Int(in map[string]int) (out map[int]int) {
	out = make(map[int]int)
	for key, val := range in {
		keyInt, _ := strconv.Atoi(key)
		out[keyInt] = val
	}
	return
}
