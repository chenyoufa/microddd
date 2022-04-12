package aggregate

import (
	"errors"
	"microddd/domain/entity"

	"github.com/google/uuid"
)

type Member_aggre struct {
	User     entity.User_entity
	role     []*entity.Role_entity
	userrole []*entity.UserRole_entity
}

func NewMember(loginname string, email string, password string) (*Member_aggre, error) {
	if loginname == "" || email == "" || password == "" {
		return nil, errors.New("请输入必填项")
	}
	user := entity.User_entity{
		uuid.New(),
		loginname,
		email,
		password,
	}
	return &Member_aggre{
		user,
		make([]*entity.Role_entity, 0),
		make([]*entity.UserRole_entity, 0)}, nil
}
