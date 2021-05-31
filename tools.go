package main

import (
	"strconv"
	"strings"
)

func splitToInt(val string, sep string) []int {
	if len(val) == 0 {
		return []int{}
	}
	vals := strings.Split(val, sep)
	ret := make([]int, len(vals))
	for i, v := range vals {
		vint, _ := strconv.Atoi(v)
		ret[i] = vint
	}
	return ret
}
