package convert

import (
	"strconv"
)

func BytesToInt(bys []byte) int {
	data, _ := strconv.Atoi(string(bys))
	return int(data)
}
