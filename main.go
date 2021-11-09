package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"sync"

	"github.com/HimawariSunFlower/goTest/algo"
	_ "github.com/HimawariSunFlower/goTest/docs"
	"github.com/HimawariSunFlower/goTest/test"
	"github.com/HimawariSunFlower/goTest/zaplog"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

//todo 用gin拉一个网站,swagger提供api,网站传递参数去测试
//万能方法 每个模块实现一个接口,网页传参?
//swagger提供api 一个模块对应一个入口,可变参数,具体api判断参数长度
//排序,queue这种数据结构的测试暂时不动 test里的能拆的拆一下

//测试没办法抽象出来,算法,运算符可以,测试加一个button然后返回固定格式的返回就行了

// @title goTest API
// @version V114514
// @description 自用go语言测试
// @in header
func main() {
	zaplog.InitLogger()
	//todo2()
	todo3()
	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//test mathem
	r.GET("/test/testModulo/:Modulo", testModulo)
	r.GET("/test/testBitOperation/:BitOperation/:Type", testBitOperation)

	r.GET("/test/test/", testfunc)
	r.Run() // listen and serve on 0.0.0.0:8080

}

func testModulo(c *gin.Context) {
	algo.MathApi(c, 1)
}

func testBitOperation(c *gin.Context) {
	algo.MathApi(c, 2)
}

// @Summary 测试按钮
// @description print输出
// @Produce  json
// @Router /test/test/ [get]
func testfunc(c *gin.Context) {
	test.Api()
}

// func todo() {//i:9 i:10X10(main.go-60) i:0 i:1...i:8
// 	runtime.GOMAXPROCS(1)
// 	wg := sync.WaitGroup{}
// 	wg.Add(20)
// 	for i := 0; i < 10; i++ {
// 		go func() {
// 			fmt.Println("i: ", i)
// 			wg.Done()
// 		}()
// 	}
// 	for i := 0; i < 10; i++ {
// 		go func(i int) {
// 			fmt.Println("i: ", i)
// 			wg.Done()
// 		}(i)
// 	}
// 	wg.Wait()
// }

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {

}

func live() People {
	var stu *Student
	return stu
}

func todo() {
	if s := live(); s == nil {
		fmt.Println("AAAAAAA")
	} else {

		ss := reflect.TypeOf(s)
		sv := reflect.ValueOf(s)

		fmt.Println(ss, ss.Kind(), ss.Elem().Name(), sv.IsNil())
	}
}

func todo2() {
	out := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	f := func(c chan int) {
		fmt.Println("a")
		c <- rand.Intn(5)
	}
	go func() {
		defer wg.Done()
		for i := 0; i < 5; i++ {
			f(out)
		}
		close(out)
	}()
	go func() {
		defer wg.Done()
		for i := range out {
			fmt.Println("b", i)
			//fmt.Println(i)
		}
	}()
	wg.Wait()
}

func todo3() {
	m := make(map[int]BitSet, 0)
	m[1] = BitSet(0)
	val, _ := m[1]
	val.Set(1)

	if m[1].Has(1) {
		fmt.Println(11)
	}

	x, ok := mp[0]
	fmt.Println(x, ok) //0 false

	ll := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 1, 2, 3, 4, 5}
	ll = ll[len(ll)-20:]
	fmt.Println(ll, len(ll))
}

var mp = make(map[int]int)

func init() { mp[1] = 1; mp[2] = 2 }
