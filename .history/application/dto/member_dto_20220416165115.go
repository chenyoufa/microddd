package dto

import (
	"log"
	"microddd/domain/aggregate"
	"microddd/domain/entity"
	"microddd/domain/valobj"

	"github.com/devfeel/mapper"
	"github.com/google/uuid"
)

type Member_dto struct {
	ID        string `json:""`
	LoginName string `json:""`
	Email     string `json:""`
	Password  string `json:""`
	// roles []*entity.RoleEntity
	Userroles []UserRoleDto
}
type UserRoleDto struct {
	RoleID string
	UserID string
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

		userdtos = append(userdtos, UserRoleDto{item.String(), aggre.User.ID.String()})
	}
	log.Println("userRoles:", userRoles)

	mapper.AutoMapper(userEntity, m)
	return
	m.Userroles = userdtos
	m.Password = ""
}

func (m *Member_dto) ToAggre() *aggregate.Member_aggre {
	// var aggre *aggregate.Member_aggre
	aggre, err := aggregate.NewMember(m.LoginName, m.Email, m.Password)
	if err != nil {
		return nil
	}
	var roleids []uuid.UUID
	if m.Userroles == nil {
		m.Userroles = make([]UserRoleDto, 0)
	}
	for _, item := range m.Userroles {
		uid, _ := uuid.Parse(item.RoleID)
		roleids = append(roleids, uid)
	}
	mapper.AutoMapper(m, aggre)

	aggre.AddRoles(roleids...)
	return aggre
}
