package xtime

import (
	"time"
)

func Millisecond() int64{
	return time.Now().UnixNano() / 1e6
}