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
	testQueue()
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

const defaultCap = 10

type Queue struct {
	data  []interface{}
	left  int
	right int
	cap   int
}

func newQueue() *Queue {
	return new(Queue).Init()
}
func (q *Queue) Init() *Queue {
	q.data = make([]interface{}, defaultCap, defaultCap)
	q.left, q.right, q.cap = 0, 0, defaultCap
	return q
}

func (q *Queue) isEmpty() bool {
	return q.right == 0
}

func (q *Queue) Push(inf interface{}) {
	if q.checkTyp(inf) {
		q.push(inf)
	} else {
		panic("should't used unsame type interface")
	}
}

func (q *Queue) push(inf interface{}) {
	q.lazyExpand()
	q.data[q.right] = inf
	q.right++
}

func (q *Queue) Pop() (ret interface{}, ck bool) {
	defer func() {
		if ck {
			q.left++
		}
	}()
	q.lazyShrink()
	if q.full() {
		return nil, false
	}
	return q.top(), true
}

func (q *Queue) checkTyp(inf interface{}) bool {
	if !q.isEmpty() {
		if reflect.TypeOf(inf).Kind() != reflect.TypeOf(q.zero()).Kind() {
			return false
		}
	}
	return true
}

func (q *Queue) Top() interface{} {
	if q.isEmpty() || q.full() {
		return nil
	}
	return q.top()
}

func (q *Queue) top() interface{} {
	return q.data[q.left]
}

func (q *Queue) zero() interface{} {
	return q.data[0]
}
func (q *Queue) full() bool {
	return q.left == q.right
}

func (q *Queue) lazyExpand() {
	if q.right == q.cap {
		q.data = append(q.data, q.data...)
		q.cap *= 2
	}
}

func (q *Queue) lazyShrink() {
	if q.left == q.right {
		q.left, q.right = 0, 0
		q.data = q.data[:0]
	}
}

//测试实际效果
func testQueue() {
	t := newQueue()
	for i := 0; i < 99; i++ {
		t.push(i)
	}
	k, _ := t.Pop()
	k2, _ := t.Pop()
	fmt.Print(k, k2)
	fmt.Println(t)
}
