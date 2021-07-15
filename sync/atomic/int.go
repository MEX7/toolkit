package atomic

import (
	"sync/atomic"
)

type Int64 struct {
	val int64
}

func (i *Int64) Add(input int64) int64 {
	return atomic.AddInt64(&i.val, input)
}

func (i *Int64) Swap(input int64) int64 {
	return atomic.SwapInt64(&i.val, input)
}

func (i *Int64) Get() int64 {
	return atomic.LoadInt64(&i.val)
}

func (i *Int64) CompareAndSwap(old, new int64) bool {
	return atomic.CompareAndSwapInt64(&i.val, old, new)
}
