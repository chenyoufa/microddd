package mainterface

import (
	"fmt"
	"gomicroddd/application/interfa"
	"gomicroddd/interface/dto"
)

func main() {

	var dto = dto.MemberDto{}

	memberdo := dto.DtoToDo()
	var memberapi interfa.MemberApi
	memberapi.AddMember(memberdo)
	fmt.Print("ok")

}
