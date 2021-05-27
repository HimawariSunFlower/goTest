package main

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

//todo 用gin拉一个网站,swagger提供api,网站传递参数去测试
func main() {
    r := gin.Default()
    r.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    })
    r.Run() // listen and serve on 0.0.0.0:8080
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
