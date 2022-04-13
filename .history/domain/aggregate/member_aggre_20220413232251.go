package aggregate

import (
	"errors"
	"microddd/domain/entity"

	"github.com/google/uuid"
)

//用户 角色  用户角色关系删除、冻结
type Member_aggre struct {
	User  *entity.UserEntity
	roles []*entity.RoleEntity
	// userrole []*entity.UserRoleEntity
}

func NewMember(loginname string, email string, password string) (*Member_aggre, error) {
	if loginname == "" || email == "" || password == "" {
		return nil, errors.New("请输入必填项")
	}
	user := entity.UserEntity{
		ID:        uuid.New(),
		LoginName: loginname,
		Email:     email,
		Password:  password,
	}
	return &Member_aggre{
		&user,
		make([]*entity.RoleEntity, 0),
		// make([]*entity.UserRoleEntity, 0),
	}, nil
}
