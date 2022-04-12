package aggregate

import (
	"errors"
	"microddd/domain/entity"

	"github.com/google/uuid"
)

type Member_aggre struct {
	User     entity.UserEntity
	role     []*entity.RoleEntity
	userrole []*entity.UserRoleEntity
}

func NewMember(loginname string, email string, password string) (*Member_aggre, error) {
	if loginname == "" || email == "" || password == "" {
		return nil, errors.New("请输入必填项")
	}
	user := entity.UserEntity{
		uuid.New(),
		loginname,
		email,
		password,
	}
	return &Member_aggre{
		user,
		make([]*entity.RoleEntity, 0),
		make([]*entity.UserRoleEntity, 0)}, nil
}
