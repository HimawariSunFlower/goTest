package algo

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var result = make(map[string]interface{}, 0)

func MathApi(c *gin.Context, typ int) {
	defer func() {
		result = make(map[string]interface{}, 0)
	}()
	switch typ {
	case 1:
		Modulo(c)
	case 2:
		testBitOperation(c)
	}

	c.JSON(200, result)
}

// 1.负数等于正数取反加一。
// 2.左移一位相当于将这个数扩大两倍，右移两位相当于将这个数缩小两倍
// 3.符号位向右移动后,正数补0,负数补1。
// 4.负数补码最高位是1，正数补码最高位是0

//https://learnku.com/go/t/23460/bit-operation-of-go

// @Summary 位运算
// @description param "xx|yy" type=1 xx<<yy ;type=2 xx>>yy
// @Produce  json
// @Param param path string true "参数"
// @Param type path int true "类型"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /test/testBitOperation/{param}/{type} [get]
func testBitOperation(c *gin.Context) {
	ll := splitToInt(c.Param("BitOperation"), "|")
	if len(ll) != 2 {
		ErrParam(c, "BitOperation")
		return
	}
	num1, num2 := ll[0], ll[1]
	switch c.Param("Type") {
	case "1":
		ans := _leflef(num1, num2)
		PrintMy(num1, ans, "testBitOperation")
	case "2":
		ans := _rgtrgt(num1, num2)
		PrintMy(num1, ans, "testBitOperation")
	default:
		ErrParam(c, "BitOperation-Type")
		return
	}
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

// @Summary 取余
// @description param "xx|yy" xx%yy
// @Produce  json
// @Param param path string true "参数"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /test/testModulo/{param} [get]
func Modulo(c *gin.Context) {
	ll := splitToInt(c.Param("Modulo"), "|")
	if len(ll) != 2 {
		ErrParam(c, "Modulo")
		return
	}
	num := ll[0]
	mid := ll[1]
	ans := num % mid
	PrintMy(num, ans, "Modulo")
}

func PrintMy(bef, aft interface{}, function string) {
	result[function] = fmt.Sprintf("%v => %v", bef, aft)
}

func ErrParam(c *gin.Context, key string) {
	k := "ErrParam" + key
	result[k] = fmt.Sprintf("%v", c.Query("param"))
}

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
