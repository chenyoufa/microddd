package main

import (
	"fmt"
	"gomicroddd/application/app"
	"gomicroddd/domain/entity"
	"gomicroddd/infrastructure/db/dbcore"
	"gomicroddd/interface/dto"
	"time"

	"github.com/devfeel/mapper"
	"github.com/google/uuid"
)

func init() {
	mapper.Register(&dto.MemberDto{})
	mapper.Register(&entity.User{})
}
func main() {

	err := dbcore.Connect(&dbcore.DBConfig{
		DbType:      "mysql",
		DSN:         "fage:Fage501526~@(127.0.0.1:3306)/mytest",
		AutoMigrate: true,
		Debug:       false,
	})
	if err != nil {
		fmt.Printf("err fail:%v", err)
	}

	memberapi := app.NewMemberApp()
	model, err := memberapi.Get(uuid.Parse("b4fcb682-f023-419f-9c5b-f67124da7256"))

	fmt.Println(model.User)

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
