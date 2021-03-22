package algo

import (
	"fmt"
	"reflect"
	"unsafe"
)

func MindApi() {
	str := "我佛了"
	fmt.Println(stringToBytes(str))
	fmt.Println(stringToBytes2(str))
}

func stringToBytes(str string) []byte {
	var buf []byte
	*(*string)(unsafe.Pointer(&buf)) = str
	(*reflect.SliceHeader)(unsafe.Pointer(&buf)).Cap = len(str)
	return buf
}

func stringToBytes2(str string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&str))
	h := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&h))
}
