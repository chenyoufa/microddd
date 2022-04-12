package mainterface

import (
	"fmt"
	app "gomicroddd/application"
	"gomicroddd/interface/dto"
)

func main() {

	var dto = dto.MemberDto{}

	memberdo := dto.DtoToDo()
	app.AddMember(memberdo)
	fmt.Print("ok")

}
