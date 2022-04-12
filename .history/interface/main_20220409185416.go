package main

import (
	"fmt"
	"gomicroddd/application/interfa"
	"gomicroddd/interface/dto"
	"reflect"
	"time"
)

func getUnExportedField(ptr interface{}, fieldName string) reflect.Value {

	v := reflect.ValueOf(ptr).Elem().FieldByName(fieldName)
	return v
}
func main() {

	var dto = dto.MemberDto{
		UserName:     "fage",
		RealName:     "cyf",
		Password:     "123456",
		Email:        "879756530@qq.com",
		Phone:        "13348493100",
		Status:       1,
		UpdateTimeAt: time.Now(),
		Remark:       "test",
	}
	t := reflect.TypeOf(dto)
	v := reflect.ValueOf(dto)
	for k := 0; k < t.NumField(); k++ {
		// fmt.Println(t.Field(k).Name)
		// fmt.Println(v.Field(k).Interface())
		// fmt.Println(v.Field(k))
		fmt.Println(t.Field(k).Tag.Get("mapper"))
	}

	return
	memberdo := dto.DtoToDo()
	var memberapi interfa.MemberApi
	memberapi.AddMember(memberdo)
	fmt.Print("ok")

}
