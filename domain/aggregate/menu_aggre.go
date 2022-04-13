package aggregate

import (
	"microddd/domain/entity"
	"microddd/domain/valobj"
)

type Menu_aggre struct {
	Menu        *entity.MenuEntity
	menuActions []valobj.ActionValObj
}
