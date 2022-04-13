package aggregate

import (
	"errors"
	"microddd/domain/entity"
	"microddd/domain/valobj"

	"github.com/google/uuid"
)

//用户 角色  用户角色关系删除、冻结
type Member_aggre struct {
	User *entity.UserEntity
	// roles []*entity.RoleEntity
	userroles []valobj.UserRoleValObj
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
		make([]valobj.UserRoleValObj, 0),
		// make([]*entity.UserRoleEntity, 0),
	}, nil
}

func (m *Member_aggre) Delete() {
	m.userroles = nil
}

func (m *Member_aggre) AddRoles(roleids ...uuid.UUID) {
	for roleid := range roleids {
		temp := valobj.UserRoleValObj{roleid, m.User.ID}
		m.userroles = append(m.userroles, temp)
	}

}
func (m *Member_aggre) RemoveRoles(roleids ...uuid.UUID) {
	m.userroles = append(m.roles, roles...)
}
