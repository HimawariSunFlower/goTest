package algo

import "fmt"

func MathApi() {
	testBitOperation()
}

// 1.负数等于正数取反加一。
// 2.左移一位相当于将这个数扩大两倍，右移两位相当于将这个数缩小两倍
// 3.符号位向右移动后,正数补0,负数补1。
// 4.负数补码最高位是1，正数补码最高位是0

//https://learnku.com/go/t/23460/bit-operation-of-go

//位运算
func testBitOperation() {
	i := 20
	i = _leflef(i, 2)
	i = _rgtrgt(i, 1)
}

func _leflef(i, n int) int {
	i <<= n
	fmt.Println(i)
	return i
}

func _rgtrgt(i, n int) int {
	i >>= n
	fmt.Println(i)
	return i
}
