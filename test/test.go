package test

import (
	"fmt"
	"math/rand"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func Api() {
	//身份证校验 正则
	fmt.Println(time.Now().Hour())
	fmt.Println(GetAgeWithIdentificationNumber("430929199811113115")) //5224261981110555X 0 error
	res, _ := IsIdCard("430929199811113115")
	if res {
		fmt.Println("验证通过")
	} else {
		fmt.Println("验证失败")
	}
	testRegexp("221033199902022222")
	testRegexp2("221033199902022222")

	////cmd测试
	//cmd := exec.Command("rysnc", "--version")
	//f, err := exec.LookPath("rsync")
	//if err != nil {
	//	fmt.Println(err)
	//}

	//随机数测试
	fmt.Print(GenerateRandomNumber(nil, 2, 2, 1))

	//测试slice[:]
	testSlice()

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

func GetAgeWithIdentificationNumber(identification_number string) int {
	if identification_number == "" {
		return 0
	}
	reg := regexp.MustCompile(`^[1-9]\d{5}(18|19|20)(\d{2})((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`)
	//reg := regexp.MustCompile(`^[1-9]\d{5}(18|19|20)`)
	params := reg.FindStringSubmatch(identification_number)
	if len(params) == 0 {
		fmt.Errorf("reg error")
		return 0
	}
	birYear, _ := strconv.Atoi(params[1] + params[2])
	birMonth, _ := strconv.Atoi(params[3])
	age := time.Now().Year() - birYear
	if int(time.Now().Month()) < birMonth {
		age--
	}
	return age
}

func IsIdCard(idCard string) (res bool, err error) {
	res, err = regexp.Match("^[1-9]\\d{7}((0\\d)|(1[0-2]))(([0|1|2]\\d)|3[0-1])\\d{3}$|^[1-9]\\d{5}[1-9]\\d{3}((0\\d)|(1[0-2]))(([0|1|2]\\d)|3[0-1])\\d{3}([0-9]|X)$", []byte(idCard))
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
	fmt.Println("result1 = ", result1)
}

// reg := regexp.MustCompile(`^[1-9]\d{5}(18|19|20)(\d{2})((0[1-9])|(1[0-2]))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$`)
// 	params := reg.FindStringSubmatch(identification_number)

func testRegexp2(s string) {
	res, err := regexp.Match("[1-9]\\d{5}(18|19|20)(\\d{2})(0[1-9]|1[0-2])([0-2][1-9]|10|20|30|31)\\d{3}[0-9Xx]$", []byte(s))
	fmt.Println(res)
	if err != nil {
		fmt.Println(err)
	}
}

//---------------------test random method-----------------

func GenerateRandomNumber(randObj *rand.Rand, start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
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

	return nums
}

//-----------------test  slice [:]------------------
func testSlice() { //测试slice[:]
	// ll := make([]int, 0)
	// for i := 0; i < 20; i++ {
	// 	ll = append(ll, i)
	// 	if len(ll) > 10 {
	// 		// fmt.Print(len(ll))
	// 		// ll = ll[len(ll)-9:]
	// 		ll = ll[:10]
	// 	}
	// }
	// fmt.Print(len(ll))
	// fmt.Print(ll)
	// fmt.Print(ll[:1])
	// //slice测试
	// fmt.Print(f)
	// v := []int{1, 2, 3}
	// fmt.Print(v[2:], v[2], v[:0])
	// for i := 0; i < 10; i++ {
	// 	continue
	// 	fmt.Print("太弱智了")
	// }
	// actionAgainU := []int{1, 2, 3, 4}
	// for i := 0; i < 100; i++ {
	// 	if len(actionAgainU) != 0 {
	// 		fmt.Println(actionAgainU[0])
	// 		actionAgainU = actionAgainU[1:]
	// 		fmt.Println(actionAgainU)
	// 	} else {
	// 		break
	// 	}
	// }
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//fmt.Println(effectTyp[:2], effectTyp[2+1:]) //:的机制是 [1:2)
	// list := make([]interface{}, 0)
	// list = append(list, effectTyp)
	// fmt.Println(list)
	// for i := 0; i < len(effectTyp); {
	//
	// 	effectTyp = effectTyp[:len(effectTyp)-1]
	// }
	for i := 0; i < len(s); i++ {
		//if InArray(s[i], unsame) {
		s = append(s[:i], s[i+1:]...)
		fmt.Println(s)
		i--
		//}
	}
}

//-------------test  map => slice's value--------

type a struct {
	data  []*b
	mapId map[int]*b
}

func (a *a) newA(ll []*b) {
	a.data = ll
	a.mapId = make(map[int]*b, len(ll))
	for _, v := range ll {
		a.mapId[v.id] = v
	}
}

type b struct {
	name string
	id   int
}

var defaultA = &a{}

func testMapWithSameData() { //指针指向相同内存,修改data数据,map的val也变化,delet原来的key,map[newKey]=>oldVal就可以了
	ll := []*b{}
	for i := 0; i < 10; i++ {
		ll = append(ll, &b{name: "aaa", id: i})
	}
	defaultA.newA(ll)
	for _, v := range defaultA.data {
		v.id = -1
	}
	fmt.Print(defaultA)
	fmt.Print(defaultA.mapId[1])
}

//----------------------------test  reflect----------------

type Foo struct {
	A int `tag1:"Tag1" tag2:"Second Tag"`
	B string
}

func TestReflect() {
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

func TestValueReceiver() {
	z := new(yy)
	z.x = *makeXx()
	fmt.Print(z.TestMethodWithValueReceiver())
}

//-----------------------------
type t20210203 struct {
	DD []int
}

func TestSliceWithPass() {
	ll := make([]int, 0)
	ll = append(ll, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	fmt.Println(ll) //[1 2 3 4 5 6 7 8 9]

	zz:=sortS(ll)
	fmt.Println(ll,"||",zz) //like [9 3 1 8 2 4 6 5 7] || [9 3 1 8 2 4 6 5 7]

	l2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}\
	//函数里append会让里外slice指向不同值
	z2:=TestAppend(l2)
	fmt.Println(l2,"||",z2)//[1 2 3 4 5 6 7 8 9] || [1 2 3 4 5 6 7 8 9 1 2 34 5 3 67 8 2 90 3]
	
	mid1(l2)
	mid1(ll)

	fmt.Println("z2",z2) //[1 2 3 4 5 6 7 8 9 1 2 34 5 3 67 8 2 90 3]
	fmt.Println("zz",zz) //[3 3 3 3 3 3 3 3 3]
	fmt.Println(ll) //[3 3 3 3 3 3 3 3 3]
	fmt.Println(l2) //[3 3 3 3 3 3 3 3 3]

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

func sortS(s []int) []int{
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


func TestAppend(s []int) []int{
	s = append(s, []int{1,2,34,5,3,67,8,2,90,3}...)
	return s
}