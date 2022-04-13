package aggregate

import "microddd/domain/entity"

type Role_aggre struct {
	Role     *entity.RoleEntity
	menus    []*entity.MenuEntity
	userrole []*entity.UserRoleEntity
}
