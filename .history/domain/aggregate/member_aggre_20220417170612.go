package aggregate

import (
	"errors"
	"microddd/domain/entity"
	"microddd/domain/valobj"
	"microddd/infrastructure/utils/encrypt"

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
		ID:        uuid.New().String(),
		LoginName: loginname,
		Email:     string(encrypt.AesEncryptCFB(email, "")),
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

func (m *Member_aggre) GetRoleIDs() []string {

	var rids []string
	for _, item := range m.userroles {
		rids = append(rids, item.RoleID)
	}
	return rids
}

func (m *Member_aggre) AddRoles(roleids ...string) {
	if len(roleids) <= 0 {
		return
	}
	for _, roleid := range roleids {

		temp := valobj.UserRoleValObj{RoleID: roleid, UserID: m.User.ID}
		m.userroles = append(m.userroles, temp)
	}

}
func (m *Member_aggre) RemoveRoles(roleids ...string) {

	for index, oldroleid := range m.userroles {
		if ok := isExit(oldroleid.RoleID, roleids); ok {
			newlist := append(m.userroles[:index], m.userroles[(index+1):]...)
			m.userroles = newlist
		}
	}
}
func isExit(searchval string, roleids []string) bool {
	for _, hvaid := range roleids {
		if hvaid == searchval {
			return true
		}
	}
	return false
}
