package test

import "fmt"

//-----------------test  slice [:]------------------
type s2021 struct {
	d []*ins2021
}
type ins2021 struct {
	a int
}

func TestSlice() { //测试slice[:]
	_test1()
	fmt.Println("---------------testSlice---------------")
	x := new(s2021)
	fmt.Println(*x)
	x.d = append(x.d, &ins2021{a: 1}, &ins2021{a: 2}, &ins2021{a: 3})
	fmt.Println(x.d)

	for k, v := range x.d {
		if v.a == 2 {
			x.d[k] = nil
			x.d = append(x.d[:k], x.d[k+1:]...)
		} else {
			v.a++
		}
	}
	fmt.Println(x.d)

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
	if len(s) > 6 {
		s = s[:6]
	}
	//fmt.Println(s[:2], s[2+1:]) //[:]的机制是 [1:2) // [1 2] [4 5 6 7 8 9]
	//fmt.Println(s[:len(s)], "xx")

	// list := make([]interface{}, 0)
	// list = append(list, s)
	// fmt.Println(list)
	// for i := 0; i < len(s); {
	//
	// 	s = s[:len(s)-1]
	// }
	for i := 0; i < len(s); i++ {
		//if InArray(s[i], unsame) {
		s = append(s[:i], s[i+1:]...)
		fmt.Println(s)
		i--
		//}
	}

	l1, l2 := []int{1, 2, 3, 4}, []int{7, 8, 9, 10}
	i := 0
	l3 := []int{}
outer:
	for _, k := range l1 {
		for ; i < len(l2); i++ {
			if l2[i] == 8 {
				continue outer // continue i!++,l1->->
			}
			l3 = append(l3, k)
		}
	}
	fmt.Println(l3)

}

type __test1 [5][6]*Equip

func _test1() {
	var x = new(__test1)
	fmt.Println(x)
}
