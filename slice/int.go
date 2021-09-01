package slice

import (
	"fmt"
)

func Ints2String(ints []int) string {
	var res = ""
	for k, v := range ints {
		if k == 0 {
			res = fmt.Sprintf("%d", v)
		} else {
			res = fmt.Sprintf("%s,%d", res, v)
		}
	}
	return res
}
