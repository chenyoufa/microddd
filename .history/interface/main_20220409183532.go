package mainterface

import (
	"fmt"
	"gomicroddd/application/interfa"
	"gomicroddd/interface/dto"
	"time"
)

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

	memberdo := dto.DtoToDo()
	var memberapi interfa.MemberApi
	memberapi.AddMember(memberdo)
	fmt.Print("ok")

}