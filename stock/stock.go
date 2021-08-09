package stock

import (
	"fmt"
	"strconv"
)

func Reduce(original float64, coe float64, n int) (res float64) {
	res = original
	for i := 0; i < n; i++ {
		res = res * (1 - coe)
	}
	return res
}

func Increase(original float64, coe float64, n int) (res float64) {
	res = original
	for i := 0; i < n; i++ {
		res = res * (1 + coe)
	}
	return
}

func ReDivide(original float64, coe float64, n int) (res float64) {
	res = original
	for i := 0; i < n; i++ {
		res = res / (1 - coe)
	}
	return res
}

func ReMult(original float64, coe float64, n int) (res float64) {
	res = original
	for i := 0; i < n; i++ {
		res = res / (1 + coe)
	}
	return res
}

const MIN = 0.02

// MIN 为用户自定义的比较精度
func IsEqual(f1, f2 float64) bool {
	if f1 > f2 {
		return f1-f2 < MIN
	} else {
		return f2-f1 < MIN
	}
}

func LogX(bef float64, end float64) (n int) {
	n = 1
	cal := bef
	for {
		calDim, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", cal), 64)
		if IsEqual(calDim, end) {
			return n
		}
		cal = cal * bef
		n++
	}
}
