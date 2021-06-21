package copy

import (
	"fmt"
	"reflect"
	"testing"
)

type dbStruct struct {
	Name   string
	Id     int
	Parmas []int
	C      *C
	Cs     []*C
}
type msgStruct struct {
	Name   string
	Id     int
	Parmas []int
	C      *C
	Cs     []*C
}

type C struct {
	Id int
}

func (p *msgStruct) CLear() {
	p.Id = 0
	p.Name = ""
	p.Parmas = p.Parmas[:0]
	p.C = &C{}
	p.Cs = p.Cs[:0]
}
func String(data interface{}) string {
	ret := "{"
	val := reflect.ValueOf(data).Elem()
	typ := val.Type()

	ret += typ.Name() + ": "
	for i := 0; i < typ.NumField(); i++ {
		Name := typ.Field(i).Name
		v := val.FieldByName(Name)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
			mid := make([]interface{}, 0)
			for i := 0; i < v.Len(); i++ {
				elem := v.Index(i)
				if elem.Kind() == reflect.Ptr {
					elem = elem.Elem()
				}
				mid = append(mid, elem.Interface())
			}
			ret += fmt.Sprintf("%s:%v,", Name, mid)
			continue
		}
		ret += fmt.Sprintf("%s:%v,", Name, v)
	}

	return ret[:len(ret)-1] + "}"
}

//todo
func Benchmark1(b *testing.B) {

}

func TestCopyStruct(t *testing.T) {
	a := &dbStruct{Name: "初号机", Id: 001, Parmas: []int{1, 2, 3}, C: &C{Id: 002}, Cs: []*C{{Id: 003}}}
	b := new(msgStruct)
	copyStruct(b, a)
	fmt.Println(String(b))
	b.CLear()
	fmt.Println(String(b))
	fmt.Println(String(a))
}
