package test

import "fmt"

type bb struct {
	id   int
	data *b
}

func mapTest() {
	base := &bb{id: 1}
	base.data = &b{name: "x"}

	m := make(map[int]*b)
	m[1] = base.data
	base.data = nil
	fmt.Println(m)
	fmt.Println(base)

	m2 := make(map[int]int)
	val := m2[2]
	x := val + 1
	fmt.Println(x)
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
	fmt.Println("---------------testMapWithSameData---------------")
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

func TestMapNil() {
	m := make(map[int]int)
	m2 := make(map[int][]int)

	if m[2] == 0 {
		fmt.Print(0)
	}

	if m2[2] == nil {
		fmt.Print("nil")
	}
	ll := []int{}
	m2[1] = []int{1, 2}

	ll = m2[0]
	ll = m2[1]
	fmt.Println(ll)
}
