package algo

import (
	"fmt"
	"sort"
)

func makeSlice(len int) []int {
	return make([]int, len, len)
}

func Api() {
	l1 := []int{4, 5, 6, 3, 2, 1}
	l2 := makeSlice(len(l1))
	l3 := makeSlice(len(l1))
	copy(l2, l1)
	copy(l3, l1)

	bubbleSort(l1, len(l1)) //todo 测试一下struct slice  []*struct / []struct
	fmt.Println("bubbleSort", l1)

	retl2 := merge_sort(l2, len(l2))
	fmt.Println("merge_sort 不修改原数组", retl2)
	merge_sort2(l2, len(l2))
	fmt.Println("merge_sort 修改原数组", l2)

	Quicksort(l3, len(l3))
	fmt.Println("quicksort", l3)
}

// 冒泡排序，a表示数组，n表示数组大小
func bubbleSort(a []int, n int) {
	if n <= 1 {
		return
	}

	for i := 0; i < n; i++ {
		// 提前退出冒泡循环的标志位
		var flag = false
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] { // 交换
				var tmp = a[j]
				a[j] = a[j+1]
				a[j+1] = tmp
				flag = true // 表示有数据交换
			}
		}
		if !flag {
			break
		} // 没有数据交换，提前退出
	}
}

// 插入排序，a表示数组，n表示数组大小
//func insertionSort_func(data lessSwap, a, b int) { //官方库
// 	for i := a + 1; i < b; i++ {
// 		for j := i; j > a && data.Less(j, j-1); j-- {
// 			data.Swap(j, j-1)
// 		}
// 	}
// }
func insertionSort(a []int, n int) {
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
		value := a[i]
		j := i - 1
		// 查找插入的位置
		for ; j >= 0; j-- {
			if a[j] > value {
				a[j+1] = a[j] // 数据移动
			} else {
				break
			}
		}
		a[j+1] = value // 插入数据
	}
}

// 归并排序算法, A是数组，n表示数组大小
func merge_sort(A []int, n int) []int {
	return merge_sort_c(A)
}

// 递归调用函数
func merge_sort_c(A []int) []int {
	// 递归终止条件
	if len(A) == 1 {
		return A
	}

	// 取p到r之间的中间位置q
	q := len(A) / 2
	// 分治递归
	lefe := merge_sort_c(A[:q])
	right := merge_sort_c(A[q:])
	// 将A[p...q]和A[q+1...r]合并为A[p...r]
	return merge(lefe, right)
}

func merge(A, B []int) []int {
	temp := []int{}
	i, j := 0, 0
	for i < len(A) && j < len(B) {
		if A[i] <= B[j] {
			temp = append(temp, A[i])
			i++
		} else {
			temp = append(temp, B[j])
			j++
		}
	}
	if i < len(A) {
		temp = append(temp, A[i:]...)
	}
	if j < len(B) {
		temp = append(temp, B[j:]...)
	}
	return temp
}

// 归并排序算法, A是数组，n表示数组大小
func merge_sort2(A []int, n int) {
	merge_sort_my(A, 0, n)
}

func merge_sort_my(A []int, l, r int) {
	if r-l <= 1 {
		return
	}
	mid := (r + l) / 2 // A[l,mid) 和 array[mid,r)
	merge_sort_my(A, l, mid)
	merge_sort_my(A, mid, r)
	merge2(A, l, mid, r)
}

func merge2(A []int, l, mid, r int) {
	temp := []int{}
	i, j := 0, 0
	lefL, rgtL := mid-l, r-mid
	for i < lefL && j < rgtL {
		if A[l+i] <= A[mid+j] {
			temp = append(temp, A[l+i])
			i++
		} else {
			temp = append(temp, A[mid+j])
			j++
		}
	}

	if i < lefL {
		temp = append(temp, A[l+i:mid]...)
	}

	if j < rgtL {
		temp = append(temp, A[j+mid:r]...)
	}
	for i := 0; i < (r - l); i++ {
		A[l+i] = temp[i]
	}
}

func Quicksort(A []int, n int) { //todo quicksort
	quick_sort(A, 0, n-1)
}

func quick_sort(A []int, l, r int) {
	if r-l <= 0 {
		return
	}

	mid := partition(A, l, r)
	quick_sort(A, l, mid-1)
	quick_sort(A, mid+1, r)
}

/*
	魔鬼隐藏在细节中,临界问题
*/
func partition(A []int, l, r int) int { //todo 理解临界问题i,mid的传值
	pivot := A[r]
	i := l - 1
	for j := l; j < r; j++ {
		if A[j] < pivot {
			i++
			A[j], A[i] = A[i], A[j]
		}
	}
	A[i+1], A[r] = A[r], A[i+1]
	return i + 1
}

type persons struct {
	data []*person
}
type person struct {
	age  int
	name string
}

func (p persons) Len() int      { return len(p.data) }
func (p persons) Swap(i, j int) { p.data[i], p.data[j] = p.data[j], p.data[i] }

func (p persons) Less(i, j int) bool {

	return p.data[i].age < p.data[j].age
}

func TestLess() {
	x := new(persons)
	x.data = append(x.data, &person{age: 1, name: "1"})
	x.data = append(x.data, &person{age: 12, name: "12"})
	x.data = append(x.data, &person{age: 15, name: "15"})
	sort.Sort(x)
	fmt.Println(x.data[0].age, x.data[1].age, x.data[2].age)
}
