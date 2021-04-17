package test

import (
	"reflect"
)

type dbStruct struct {
	name string
	Id   int
}

//todo
func copyStruct(dst, src interface{}) {
	dstVal := reflect.ValueOf(dst)
	srcVal := reflect.ValueOf(src)
	if dstVal.Kind() != reflect.Ptr || srcVal.Kind() != reflect.Ptr {
		panic("dst, src must be ptr")
	}

	dstVal = dstVal.Elem()
	srcVal = srcVal.Elem()
	if srcVal.IsNil() {
		panic("can't copy nil struct")
	}

	srcTyp := srcVal.Type()
	for i := 0; i < srcVal.NumField(); i++ {
		srcName := srcTyp.Field(i).Name
		dstField := dstVal.FieldByName(srcName)
		if dstField.CanSet() {
			srcField := srcVal.FieldByName(srcName)
			if srcField.IsValid() {
				//todo 判断slice
				copyField(dstField, srcField)
			}
		}
	}

}

//todo 判断map
func copyField(dst, src reflect.Value) {
	
}

//todo
func copySlice() {

}
