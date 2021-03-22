package algo

import (
	"testing"
)

//➜  hello go test -bench=. -benchtime=3s -run=none
// Benchmark 名字 - CPU     循环次数          平均每次执行时间
// BenchmarkSprintf-8      50000000               109 ns/op
// PASS
//   哪个目录下执行go test         累计耗时
// ok      flysnow.org/hello       5.628s
//----------------------------------------

//API server listening at: 127.0.0.1:25032
// goos: windows
// goarch: amd64
// pkg: github.com/HimawariSunFlower/goTest/algo
// Benchmark1
// Benchmark1-12    	 4456108	       269 ns/op
// PASS
func Benchmark1(b *testing.B) {
	ll := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Quicksort(ll, len(ll))
	}
}

//API server listening at: 127.0.0.1:12027
// goos: windows
// goarch: amd64
// pkg: github.com/HimawariSunFlower/goTest/algo
// Benchmark2
// Benchmark2-12    	  802208	      1288 ns/op
// PASS
func Benchmark2(b *testing.B) {
	ll := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		merge_sort(ll, len(ll))
	}
}
