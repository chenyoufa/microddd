package aggregate

import "microddd/domain/entity"

type Menu_aggre struct {
	Menu      *entity.MenuEntity
	roleMenus []*entity.RoleMenuEntity
}
