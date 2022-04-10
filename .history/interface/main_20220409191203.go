package main

import (
	"fmt"
	"gomicroddd/application/interfa"
	"gomicroddd/domain/aggregate"
	"gomicroddd/domain/entity"
	"gomicroddd/interface/dto"
	"time"

	"github.com/devfeel/mapper"
)

func init() {
	mapper.Register(&dto.MemberDto{})
	mapper.Register(&entity.User{})
}
func main() {

	do := &aggregate.Member{}
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
	mapper.Mapper(&dto, do)
	fmt.Println(do)
	// t := reflect.TypeOf(dto)
	// v := reflect.ValueOf(dto)
	// for k := 0; k < t.NumField(); k++ {
	// fmt.Println(t.Field(k).Name)
	// fmt.Println(v.Field(k).Interface())
	// fmt.Println(v.Field(k))
	// fmt.Println(t.Field(k).Tag.Get("name"))
	// }

	return
	memberdo := dto.DtoToDo()
	var memberapi interfa.MemberApi
	memberapi.AddMember(memberdo)
	fmt.Print("ok")

}
