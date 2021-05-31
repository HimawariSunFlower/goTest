package main

import (
	"github.com/HimawariSunFlower/goTest/algo"
	_ "github.com/HimawariSunFlower/goTest/docs"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//todo 用gin拉一个网站,swagger提供api,网站传递参数去测试
//万能方法 每个模块实现一个接口,网页传参?
//swagger提供api 一个模块对应一个入口,可变参数,具体api判断参数长度
//排序,queue这种数据结构的测试暂时不动 test里的能拆的拆一下

// @title goTest API
// @version V114514
// @description 自用go语言测试
// @in header
func main() {
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//test mathem
	r.GET("/test/testModulo/:Modulo", testModulo)
	r.GET("/test/testBitOperation/:BitOperation/:Type", testBitOperation)

	r.Run() // listen and serve on 0.0.0.0:8080
}

func testModulo(c *gin.Context) {
	algo.MathApi(c, 1)
}

func testBitOperation(c *gin.Context) {
	algo.MathApi(c, 2)
}
