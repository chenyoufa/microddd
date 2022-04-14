package dto

import (
	"microddd/domain/aggregate"
	"microddd/domain/entity"
	"microddd/domain/valobj"

	"github.com/devfeel/mapper"
	"github.com/google/uuid"
)

type Member_dto struct {
	ID        uuid.UUID `json:""`
	LoginName string    `json:""`
	Email     string    `json:""`
	Password  string    `json:"_"`
	// roles []*entity.RoleEntity
	userroles []UserRoleDto
}
type UserRoleDto struct {
	RoleID uuid.UUID
	UserID uuid.UUID
}

func init() {
	mapper.Register(&Member_dto{})
	mapper.Register(&UserRoleDto{})

	mapper.Register(&entity.UserEntity{})
	// mapper.Register(&entity.RoleEntity{})
	mapper.Register(&valobj.UserRoleValObj{})
}

func (m *Member_dto) ToDto(aggre *aggregate.Member_aggre) {
	userEntity := aggre.User
	userRoles := aggre.GetRoleIDs()
	userdtos := []UserRoleDto{}
	for _, item := range userRoles {
		userdtos = append(userdtos, UserRoleDto{item, aggre.User.ID})
	}
	mapper.AutoMapper(userEntity, m)
	m.userroles = userdtos
}

func (m *Member_dto) ToAggre() *aggregate.Member_aggre {
	var aggre *aggregate.Member_aggre
	var roleids []uuid.UUID
	for _, item := range m.userroles {
		roleids = append(roleids, item.RoleID)
	}
	mapper.AutoMapper(m, aggre)

	aggre.AddRoles(roleids...)
	return aggre
}
