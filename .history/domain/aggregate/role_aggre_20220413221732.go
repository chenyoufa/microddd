package aggregate

import "microddd/domain/entity"

type Role_aggre struct {
	Role     *entity.RoleEntity
	role     []*entity.MenuEntity
	userrole []*entity.UserRoleEntity
}