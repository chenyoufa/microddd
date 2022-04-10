package reflect

import "reflect"

func getUnExportedField(ptr interface{}, fieldName string) reflect.Value {
	v := reflect.ValueOf(ptr).Elem().FieldByName(fieldName)
	return v
}
