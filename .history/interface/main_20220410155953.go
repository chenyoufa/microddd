package inteface

import (
	"fmt"
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

	MemberAppInterface

	do := &entity.User{}

	var dto = dto.MemberDto{
		UName:        "fage",
		RName:        "cyf",
		Pwd:          "123456",
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

}