package copy

import (
	"reflect"

	"github.com/HimawariSunFlower/goTest/zaplog"
)

//todo 一个结构存储copy过的结构,方法,再次copy直接取存的方法
func copyStruct(dst, src interface{}) {
	dstVal := reflect.ValueOf(dst)
	srcVal := reflect.ValueOf(src)
	if dstVal.Kind() != reflect.Ptr || srcVal.Kind() != reflect.Ptr {
		zaplog.Errorf("dst, src must be ptr")
	}
	if srcVal.IsNil() {
		zaplog.Errorf("can't copy nil struct")
	}

	dstVal = dstVal.Elem()
	srcVal = srcVal.Elem()

	srcTyp := srcVal.Type()
	for i := 0; i < srcVal.NumField(); i++ {
		srcName := srcTyp.Field(i).Name
		dstField := dstVal.FieldByName(srcName)
		if dstField.CanSet() {
			srcField := srcVal.FieldByName(srcName)
			if srcField.IsValid() {
				if srcField.Kind() == reflect.Array || srcField.Kind() == reflect.Slice {
					copySlice(&dstField, &srcField)
				} else {
					copyField(dstField, srcField)
				}
			}
		}
	}

}

//todo
func copyField(dst, src reflect.Value) {
	if dst.Kind() == reflect.Ptr {
		if src.Kind() == reflect.Ptr {
			dst.Set(src)
		} else {
			zaplog.Errorf("%v %v can't copy", src.Kind(), dst.Kind())
		}
	} else if dst.Kind() == src.Kind() {
		dst.Set(src)
	} else {
		zaplog.Errorf("%v %v can't copy", src.Kind(), dst.Kind())
	}
}

//todo
func copySlice(dst, src *reflect.Value) {
	if dst.Type().Elem().Kind() == src.Type().Elem().Kind() {
		dst.Set(reflect.MakeSlice(dst.Type(), src.Len(), src.Len()))
		reflect.Copy(*dst, *src)
	} else {
		zaplog.Errorf("%v %v can't copy", src.Kind(), dst.Kind())
	}
}

func mustBe(val reflect.Value, kind reflect.Kind) bool {
	if val.Type().Kind() != kind {
		zaplog.Errorf("%s.Kind must be %s", val.Type().Name(), kind.String())
		return false
	}
	return true
}
