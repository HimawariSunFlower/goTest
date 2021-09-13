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

var bits = 64

type BitSet int64

func (p BitSet) Has(key int) bool {
	return int64(p)&(1<<(key%bits)) != 0
}

func (p *BitSet) Set(key int) {
	*p = BitSet(int64(*p) | 1<<(key%bits))
}

func (p *BitSet) Unset(key int) {
	*p = BitSet(int64(*p) ^ (1 << (key % bits)))
}

func (p BitSet) Check(key int) int32 {
	if p.Has(key) {
		return 1
	}
	return 0
}
