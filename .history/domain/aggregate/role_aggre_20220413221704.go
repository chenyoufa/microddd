package aggregate

import "microddd/domain/entity"

type Role_aggre struct {
	User     *entity.RoleEntity
	role     []*entity.RoleEntity
	userrole []*entity.UserRoleEntity
}
