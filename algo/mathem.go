package algo

import (
	"fmt"
	"time"
)

func MathApi() {
	testBitOperation()
	_Modulo(32)
	_Modulo(64)

	fmt.Println("----------------")
	//fmt.Println((int(time.Now().Unix()) - 1619712000) | _leflef(99999-1000, 29)) //2021 5 17 18 :
	r:=time.Now().Unix()-1619712000
	v:=int64(99999-1000)<<29
	fmt.Printf("%F\n",float64(r | v))
	fmt.Printf("%F\n",float64(r | int64(99999-1200)<<29))
	fmt.Printf("%F\n",float64(r | int64(99999-1001)<<29))
	fmt.Printf("%F\n",float64((r-20000) | int64(99999-1000)<<29))
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

func _Modulo(num int) {
	for i := 0; i < 100; i++ {
		fmt.Println(i, " ", i%num)
	}
}
