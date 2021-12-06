package kmap

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func MapKeyString2Int(in map[string]int) (out map[int]int) {
	out = make(map[int]int)
	for key, val := range in {
		keyInt, _ := strconv.Atoi(key)
		out[keyInt] = val
	}
	return
}

// MergeLabels merge label
// the new map will overwrite the old one.
// e.g. new: {"foo": "newbar"} old: {"foo": "bar"} will return {"foo": "newbar"}
func MergeLabels(old map[string]string, new map[string]string) map[string]string {
	if new == nil {
		return old
	}

	if old == nil {
		old = make(map[string]string)
	}

	for key, value := range new {
		old[key] = value
	}
	return old
}

func LabelsToString(labels map[string]string) string {
	result := make([]string, 0)
	for k, v := range labels {
		result = append(result, fmt.Sprintf("%s=%s", k, v))

	}

	return strings.Join(result, ",")
}

func SyncMapLen(m *sync.Map) int {
	length := 0
	m.Range(func(key, value interface{}) bool {
		length++
		return true
	})
	return length
}
