package reflect

import "reflect"

//通过反射获取私有字段值
func getUnExportedField(ptr interface{}, fieldName string) reflect.Value {
	v := reflect.ValueOf(ptr).Elem().FieldByName(fieldName)
	return v
}
