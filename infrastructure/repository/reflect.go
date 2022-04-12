package repository

import (
	"fmt"
	"reflect"
	"unsafe"
)

//通过反射获取私有字段值
func getUnExportedField(ptr interface{}, fieldName string) reflect.Value {
	v := reflect.ValueOf(ptr).Elem().FieldByName(fieldName)
	return v
}

func setUnExportedStrField(ptr interface{}, fieldName string, newFieldVal interface{}) error {
	v := reflect.ValueOf(ptr).Elem().FieldByName(fieldName)
	v = reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
	nv := reflect.ValueOf(newFieldVal)
	if v.Kind() != nv.Kind() {
		return fmt.Errorf("expected kind %v, got kind:%v", v.Kind(), nv.Kind())
	}
	v.Set(nv)
	return nil
}
