package test

import (
	"fmt"
	"math/rand"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func GetFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

//todo 格式化
func Api() {
	printMy := func(f func()) {
		fmt.Println("----------------" + reflect.TypeOf(f).Name() + "------------------")
		f()
		fmt.Println("----------------  END ----------------------")
	}
	printMys := func(f func(string), data string) {
		fmt.Println("----------------" + GetFunctionName(f) + "------------------")
		f(data)
		fmt.Println("----------------  END ----------------------")
	}

	//身份证校验 正则
	fmt.Println("time.hour =>", time.Now().Hour())

	printMys(GetAgeWithIdentificationNumber, "430929199811113115")

	printMys(IsIdCard, "430929199811113115")

	printMys(testRegexp, "221033199902022222")
	printMys(testRegexp2, "221033199902022222")

	// str := "time='2021-07-24T05:50:51+08:00' level=debug msg='xxxxx' Acc=xx UID=1 fid=1627077050 pet=291425748946 pos=15"
	// reg := regexp.MustCompile(`Acc=\w+ UID=\d+ fid=\d+`)
	// fmt.Println(reg.FindAllString(str, -1))

	////cmd测试
	//cmd := exec.Command("rysnc", "--version")
	//f, err := exec.LookPath("rsync")
	//if err != nil {
	//	fmt.Println(err)
	//}

	//随机数测试
	printMy(GenerateRandomNumber)

	//测试slice[:]
	TestSlice()

	//测试map指向数据,改变数据
	testMapWithSameData()

	//测试反射
	TestReflect()

	//测试byteToNum
	a := []byte{'1', '4', '4'}
	fmt.Print(toNum(a))

	//求模,去小数
	fmt.Print(ClipAdd(199, 22))

	//测试二进制与或运算,reflect.flag
	TestFlag()

	//splitToInt测试
	items := strings.Split("-1", "|")
	for _, v := range items {
		vals := splitToInt(v, ":")
		fmt.Print(vals)
	}
	fmt.Print(items)

	//测试指针存入slice再改变slice中的值会不会修改元数据
	TestPtrInSlice()

	//测试(指针作为struct的字段)和(值作为struct的字段,指针作为方法接收者返回字段的指针)两者是否都由可以修改字段内容
	TestValueReceiver()

	//测试 slice 传值问题
	TestSliceWithPass()
	/*
		  	如果传递slice,不同地方都可能改变slice的值的话,使用copy等赋值的方法,而不是直接传递原slice.
			就算只是读取,但是原slice有重新排序或者改变值的操作,也要使用复制而不是值传递,
			因为都指向一个原slice,原slice的改变会影响值传递获得的slice的值,数据不可信.
	*/

	//测试在rangemap的时候delete参数
	TestMapRangeWithDel()

	fmt.Println(randNSlice(4, 5))

	testSliceCopy()

	printMy(demo)

	printMy(testNil)

	printMy(testNumDay)

	printMy(mapTest)

	printMy(timeDay)
}

//-----------------------test  utils func-------------

func toNum(c []byte) byte {
	return ((c[0]-'0')*8+c[1]-'0')*8 + c[2] - '0'
}

func splitToInt(val string, sep string) []int {
	vals := strings.Split(val, sep)
	ret := make([]int, len(vals))
	for i, v := range vals {
		vint, _ := strconv.Atoi(v)
		ret[i] = vint
	}
	return ret
}

//-------------------test  regexp-----------------------

func GetAgeWithIdentificationNumber(identification_number string) {
	if identification_number == "" {
		return
	}
	reg := regexp.MustCompile(`^[1-9]\d{5}(18|19|20)(\d{2})((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`)
	//reg := regexp.MustCompile(`^[1-9]\d{5}(18|19|20)`)
	params := reg.FindStringSubmatch(identification_number)
	if len(params) == 0 {
		fmt.Errorf("reg error")
		return
	}
	birYear, _ := strconv.Atoi(params[1] + params[2])
	birMonth, _ := strconv.Atoi(params[3])
	age := time.Now().Year() - birYear
	if int(time.Now().Month()) < birMonth {
		age--
	}
	fmt.Println(identification_number, " age=> ", age)
}

func IsIdCard(idCard string) {
	res, err := regexp.Match("^[1-9]\\d{7}((0\\d)|(1[0-2]))(([0|1|2]\\d)|3[0-1])\\d{3}$|^[1-9]\\d{5}[1-9]\\d{3}((0\\d)|(1[0-2]))(([0|1|2]\\d)|3[0-1])\\d{3}([0-9]|X)$", []byte(idCard))
	if res {
		fmt.Println(idCard, " 验证通过")
	} else {
		fmt.Println(idCard, " 验证失败", err)
	}
	return
}

func testRegexp(inf string) {
	//解析正则表达式，如果成功返回解释器
	reg1 := regexp.MustCompile(`^[1-9]\d{5}(18|19|20)(\d{2})(0[1-9]|1[0-2])([0-2][1-9]|10|20|30|31)\d{3}[0-9Xx]$`)
	if reg1 == nil { //解释失败，返回nil
		fmt.Println("regexp err")
		return
	}
	//根据规则提取关键信息
	result1 := reg1.FindStringSubmatch(inf)
	fmt.Println(inf, " result1 = ", result1)
}

// reg := regexp.MustCompile(`^[1-9]\d{5}(18|19|20)(\d{2})((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`)
// 	params := reg.FindStringSubmatch(identification_number)

func testRegexp2(s string) {
	res, err := regexp.Match("[1-9]\\d{5}(18|19|20)(\\d{2})(0[1-9]|1[0-2])([0-2][1-9]|10|20|30|31)\\d{3}[0-9Xx]$", []byte(s))
	fmt.Println(s, " result2 = ", res)
	if err != nil {
		fmt.Println("regexp err", err)
	}
}

//---------------------test random method-----------------
func GenerateRandomNumber() {
	generateRandomNumber(nil, 2, 2, 1)
}

func generateRandomNumber(randObj *rand.Rand, start int, end int, count int) {
	//范围检查
	if end < start || (end-start) < count {
		return
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := randObj
	if r == nil {
		r = rand.New(rand.NewSource(int64(time.Now().UnixNano())))
	}

	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}

		if !exist {
			nums = append(nums, num)
		}
	}

	fmt.Printf("start %d, end %d, count %d,ret => %v", start, end, count, nums)

}

//----------------------------test  reflect----------------

type Foo struct {
	A int `tag1:"Tag1" tag2:"Second Tag"`
	B string
}

func TestReflect() {
	fmt.Println("---------------TestReflect---------------")
	// Struct
	f := &Foo{A: 10, B: "Salutations"}

	// // Struct类型的指针
	// fPtr := &f
	// // Map
	// m := map[string]int{"A": 1, "B": 2}
	// // channel
	// ch := make(chan int)
	// // slice
	// sl := []int{1, 32, 34}
	// //string
	// str := "string var"
	// // string 指针
	// strPtr := &str

	// tMap := examiner(reflect.TypeOf(f), 0)
	// tMapPtr := examiner(reflect.TypeOf(fPtr), 0)
	// tMapM := examiner(reflect.TypeOf(m), 0)
	// tMapCh := examiner(reflect.TypeOf(ch), 0)
	// tMapSl := examiner(reflect.TypeOf(sl), 0)
	// tMapStr := examiner(reflect.TypeOf(str), 0)
	// tMapStrPtr := examiner(reflect.TypeOf(strPtr), 0)

	// fmt.Println("tMap :", tMap)
	// fmt.Println("tMapPtr: ", tMapPtr)
	// fmt.Println("tMapM: ", tMapM)
	// fmt.Println("tMapCh: ", tMapCh)
	// fmt.Println("tMapSl: ", tMapSl)
	// fmt.Println("tMapStr: ", tMapStr)
	// fmt.Println("tMapStrPtr: ", tMapStrPtr)

	ReflectValue := reflect.ValueOf(f).Elem()
	reflectT := reflect.TypeOf(f)
	if ReflectValue.CanAddr() {
		fmt.Println(ReflectValue.Kind(), ReflectValue.Addr(), ReflectValue.Addr().Interface())
		fmt.Println(reflectT, reflectT.Elem(), reflect.New(reflectT), reflect.New(reflectT.Elem()))
	} else {
		fmt.Print("error")
	}

	z := make([]*Foo, 0)
	z = append(z, f)
	zT := reflect.TypeOf(z)
	fmt.Println(zT.Elem(), zT.Elem().Elem(), reflect.New(zT.Elem().Elem()))
}

func examiner(t reflect.Type, depth int) map[int]map[string]string {
	outType := make(map[int]map[string]string)

	// 如果是以下类型，重新验证
	switch t.Kind() {
	case reflect.Array, reflect.Chan, reflect.Map, reflect.Ptr, reflect.Slice:
		fmt.Println("这几种类型Name是空字符串：", t.Name(), ", Kind是：", t.Kind(), "elem是: ", t.Elem())
		// 递归查询元素类型
		tMap := examiner(t.Elem(), depth)
		for k, v := range tMap {
			outType[k] = v
		}

	case reflect.Struct:
		fmt.Println("Name是：", t.Name(), ", Kind是：", t.Kind())
		for i := 0; i < t.NumField(); i++ {
			f := t.Field(i) // reflect字段
			outType[i] = map[string]string{
				"Name": f.Name,
				"Kind": f.Type.String(),
			}
		}
	default:
		// 直接验证类型
		fmt.Println("Name是：", t.Name(), ", Kind是：", t.Kind())
		outType = map[int]map[string]string{depth: {"Name": t.Name(), "Kind": t.Kind().String()}}
	}

	return outType
}

//--------------------------test  求模,去小数 -------------

func ClipAdd(i int64, j int64) int64 {
	m := i % j
	fmt.Println(m)
	return j - m
}

//------------- flag with << >> -------------
const (
	flagKindWidth        = 5 // there are 27 kinds
	flagKindMask    flag = 1<<flagKindWidth - 1
	flagStickyRO    flag = 1 << 5
	flagEmbedRO     flag = 1 << 6
	flagIndir       flag = 1 << 7
	flagAddr        flag = 1 << 8
	flagMethod      flag = 1 << 9
	flagMethodShift      = 10
	flagRO          flag = flagStickyRO | flagEmbedRO
)

type Kind uint

const (
	Invalid Kind = iota
	Bool
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Float32
	Float64
	Complex64
	Complex128
	Array
	Chan
	Func
	Interface
	Map
	Ptr
	Slice
	String
	Struct
	UnsafePointer
)

type flag uintptr

func TestFlag() {
	fmt.Println("---------------TestFlag---------------")
	ll := make([]flag, 0)
	ll = []flag{flagKindWidth, flagKindMask, flagStickyRO, flagEmbedRO, flagIndir, flagAddr, flagMethod, flagMethodShift, flagRO}
	for _, f := range ll {
		fmt.Println("f:", f, "kind: ", "f&flagKindMask: ", Kind(f&flagKindMask))

	}
	fmt.Println(Kind(22 & flagKindMask))
}

//--------------test ptr in slice change value-------------
type xx struct {
	Id   int
	id2  int
	Name string
	a    []*b
}

func TestPtrInSlice() {
	fmt.Println("---------------TestPtrInSlice---------------")
	ac := makeXx()
	// ll := make([]*xx, 0)

	// ll = append(ll, ac)

	// for _, v := range ll {
	// 	v.Id = 3
	// 	v.id2 = 4
	// 	v.Name = "yeeeee"
	// 	fmt.Println(v)
	// }
	// fmt.Println(ac)
	ac.TestPtrInSliceMethod()
	fmt.Println(ac.a)
	fmt.Println(*ac.a[0])
}

func makeXx() *xx {
	ac := new(xx)
	ac.Id = 1
	ac.id2 = 2
	ac.Name = "asd"
	ac.a = make([]*b, 0)
	ac.a = append(ac.a, &b{name: "31", id: 222})
	ac.a = append(ac.a, &b{name: "zz", id: 333})
	return ac
}

func (x *xx) TestPtrInSliceMethod() {
	fmt.Println("---------------TestPtrInSliceMethod---------------")
	ll := make([]*b, 0)
	ll = x.a //ll=append(ll,x.a[0])
	for _, v := range ll {
		v.id = 0
		v.name = "yesss"
	}
	fmt.Println(ll)
	fmt.Println(*ll[0])
}

type yy struct {
	x xx
}

func (y *yy) TestMethodWithValueReceiver() *xx {
	return &y.x
}
func (y yy) Changex() {
	y.x = xx{}
}
func (y yy) changeX() {
	z := &y
	z.x = xx{}
}
func (y *yy) changeY() {
	y.x = xx{}
}
func TestValueReceiver() {
	fmt.Println("---------------TestValueReceiver---------------")
	z := new(yy)
	z.x = *makeXx()
	fmt.Println(z.TestMethodWithValueReceiver())
	z.Changex()
	fmt.Println(z.TestMethodWithValueReceiver())
	z.changeX()
	fmt.Println(z.TestMethodWithValueReceiver())
	z.changeY()
	fmt.Println(z.TestMethodWithValueReceiver())
}

//-----------------------------
type t20210203 struct {
	DD []int
}

func TestSliceWithPass() {
	fmt.Println("---------------TestSliceWithPass---------------")
	ll := make([]int, 0)
	ll = append(ll, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	fmt.Println(ll) //[1 2 3 4 5 6 7 8 9]

	zz := sortS(ll)
	fmt.Println(ll, "||", zz) //like [9 3 1 8 2 4 6 5 7] || [9 3 1 8 2 4 6 5 7]

	l2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//函数里append会让里外slice指向不同值
	z2 := TestAppend(l2)
	fmt.Println(l2, "||", z2) //[1 2 3 4 5 6 7 8 9] || [1 2 3 4 5 6 7 8 9 1 2 34 5 3 67 8 2 90 3]

	mid1(l2)
	mid1(ll)

	fmt.Println("z2", z2) //[1 2 3 4 5 6 7 8 9 1 2 34 5 3 67 8 2 90 3]
	fmt.Println("zz", zz) //[3 3 3 3 3 3 3 3 3]
	fmt.Println(ll)       //[3 3 3 3 3 3 3 3 3]
	fmt.Println(l2)       //[3 3 3 3 3 3 3 3 3]

	X := new(t20210203)
	X.DD = ll
	X.mid2()
	fmt.Println(X.DD) //[3 3 3 3 3 3 3 3 3]
	fmt.Println(ll)   //[3 3 3 3 3 3 3 3 3]
	X.mid3()
	fmt.Println(X.DD) //[4 4 4 4 4 4 4 4 4]
	fmt.Println(ll)   //[4 4 4 4 4 4 4 4 4]
	X.mid4()
	fmt.Println(X.DD) //[]
	fmt.Println(ll)   //[4 4 4 4 4 4 4 4 4]

	//copy()改引用为复制可以避免slice共用值导致的异常
}

func mid1(s []int) {
	for _, x := range s { //range复制再循环 不改变原值
		if x != 0 {
			x = 2
		}
	}
	fmt.Println(s)                //[1 2 3 4 5 6 7 8 9]
	for i := 0; i < len(s); i++ { //改变s
		if s[i] != 0 {
			s[i] = 2
		}
	}
	fmt.Println(s) //[2 2 2 2 2 2 2 2 2]
	zz := s
	for _, x := range zz { //range复制再循环 不改变原值
		if x != 0 {
			x = 3
		}
	}
	fmt.Println(zz) //[2 2 2 2 2 2 2 2 2]
	for i := 0; i < len(zz); i++ {
		if zz[i] != 0 {
			zz[i] = 3
		}
	}
	fmt.Println(s) //zz,s全变 //[3 3 3 3 3 3 3 3 3]
}

func (x t20210203) mid2() { //值接收者直接该字段不生效
	x.DD = []int{}
}

func (x t20210203) mid3() { //修改值内容生效并且影响原生数组
	for i := 0; i < len(x.DD); i++ {
		if x.DD[i] != 0 {
			x.DD[i] = 4
		}
	}
}

func (x *t20210203) mid4() { //修改值内容生效不影响原生数组
	x.DD = []int{}
}

func sortS(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	j := 0
	for i := len(s) - 1; i > 0; i-- {
		j = RandomN(i)
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func TestAppend(s []int) []int {
	s = append(s, []int{1, 2, 34, 5, 3, 67, 8, 2, 90, 3}...)
	return s
}

func TestMapRangeWithDel() {
	fmt.Println("--------------TestMapRangeWithDel-----------")
	m := map[int]int{}

	m[0] = 0
	m[1] = 1
	m[2] = 2
	for k, _ := range m {
		if k == 1 {
			delete(m, k)
			continue
		}
		delete(m, k)
		//m[k] = v + 1
	}
	fmt.Println(m) //[0:1 2:3]
}

func randNSlice(len, n int) []int {
	ll := []int{}
	for i := 0; i < len; i++ {
		ll = append(ll, i)
	}
	if len <= n {
		return ll
	}
	return SliceIntRandomShuffle(ll)[:n]
}

func SliceIntRandomShuffle(s []int) []int {
	if len(s) <= 1 {
		return s
	}
	j := 0
	for i := len(s) - 1; i > 0; i-- {
		j = rand.Intn(i)
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func testSliceCopy() {
	p := &a{}
	p.data = []*b{}
	for i := 0; i < 10; i++ {
		p.data = append(p.data, &b{id: i})
	}

	x := []*b{}
	x = SliceCopy3(p, 5)
	for _, v := range x {
		fmt.Println(v.id) //SliceCopy 01234
	}

	j := 0
	for i := len(sliceTest) - 1; i > 0; i-- {
		j = RandomN(i + 1)
		sliceTest[i], sliceTest[j] = sliceTest[j], sliceTest[i]
	}

	for _, v := range x {
		fmt.Println(v.id) // SliceCopy 02143 SliceCopy2 40312 SliceCopy3 01234
	}
}

var sliceTest = []*b{}

func SliceCopy(data *a, len int) []*b {
	sliceTest = sliceTest[:0]
	sliceTest = data.data[:len]
	return sliceTest
}

//新建一个slice直接指向sliceTest,相当于一个指向sliceTest的指针,sliceTest改变还是影响外层使用
func SliceCopy2(data *a, len int) []*b {
	ret := []*b{}
	sliceTest = sliceTest[:0]
	sliceTest = data.data[:len]
	ret = sliceTest
	return ret
}

//append添加底层指针到ret中,ret和sliceTest不再相关
func SliceCopy3(data *a, len int) []*b {
	ret := []*b{}
	sliceTest = sliceTest[:0]
	sliceTest = data.data[:len]
	ret = append(ret, sliceTest...)
	return ret
}

func demo() {
	num := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	num2 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	for _, v := range num {
		for _, vv := range num2 {
			fmt.Println(fmt.Sprintf("%d|%d =>%d , %d&%d =>%d)", v, vv, v|vv, v, vv, v&vv))
		}
	}
	for _, v := range num {
		fmt.Println(fmt.Sprintf("%d<<(1) => %d(^=> %d) , %d<<(0) => %d(^=> %d)", v, (v << 1), ^(v << 1), v, ^(v << 0), ^(v << 0)))
	}
}

func testNil() {
	var x interface{} = nil
	var y *int = nil //有数据结构原型
	interfaceIsNil(x)
	interfaceIsNil(y)
}

func interfaceIsNil(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

type Mynum struct {
	num     int
	endtime int64
}

func testNumDay() {
	v := &Mynum{}
	v.first()
}

func (p *Mynum) first() { //num 85 1610726400
	num := (GetOpenDay() + 1) / 2
	p.Second(num, true)
	num2 := (GetOpenDay()) / 2
	p.Second(num2, false)
}
func (p *Mynum) Second(i int, ck bool) {
	p.num = i
	if ck {
		p.endtime = time.Unix(1625068800, 0).AddDate(0, 0, p.num*2).Unix()
	} else {
		p.endtime = time.Unix(1625068800, 0).AddDate(0, 0, p.num*2+1).Unix()
	}

}
func GetOpenDay() int {
	sec := time.Now().Unix() - 1625068800
	return int(sec/(3600*24)) + 1
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func timeDay() {
	//今天0点+29天
	println(time.Unix(GetTodayZero(), 0).AddDate(0, 0, 1).Unix())
}

func BeginningOfDay() time.Time {
	y, m, d := time.Now().Date()
	return time.Date(y, m, d, 0, 0, 0, 0, time.Local)
}

func GetTodayZero() int64 {
	return BeginningOfDay().Unix()
}
