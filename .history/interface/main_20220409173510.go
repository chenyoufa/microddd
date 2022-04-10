package mainterface

import (
	"fmt"
	"gomicroddd/interface/dto"
)

func main() {

	var dto = dto.MemberDto{}

	memberdo := dto.DtoToDo()
	var memberapi MemberApi
	memberapi.AddMember(memberdo)
	fmt.Print("ok")

}
