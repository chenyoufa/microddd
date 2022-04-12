package reflect

import (
	"reflect"
	"unsafe"
)

//通过反射获取私有字段值
func getUnExportedField(ptr interface{}, fieldName string) reflect.Value {
	v := reflect.ValueOf(ptr).Elem().FieldByName(fieldName)
	return v
}

func setUnExportedStrField(ptr interface{}, fieldName string, newFieldVal interface{}) {
	v := reflect.ValueOf(ptr).Elem().FieldByName(fieldName)
	v = reflect.NewAt(v.Type(), unsafe.Pointer(v.UnsafeAddr())).Elem()
}
