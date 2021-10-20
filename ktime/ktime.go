package ktime

import (
	"fmt"
	"time"
)

func Millisecond() int64 {
	return time.Now().UnixNano() / 1e6
}

func Microsecond() int64 {
	return time.Now().UnixNano() / 1e3
}

func Cost(t time.Time, label ...string) {
	tc := time.Since(t)
	if len(label) > 0 {
		fmt.Printf("label: %s, cost: %v\n", label[0], tc)
	} else {
		fmt.Printf("cost: %v\n", tc)
	}
}
