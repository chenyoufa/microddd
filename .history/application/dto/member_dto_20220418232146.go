package dto

import (
	"microddd/domain/aggregate"

	"github.com/google/uuid"
)

type Member_dto struct {
	ID        string `json:""`
	LoginName string `json:""`
	Email     string `json:""`
	Password  string `json:""`
	Status    int    `json:""`
	// roles []*entity.RoleEntity
	Userroles []UserRoleDto
}
type UserRoleDto struct {
	RoleID string
	UserID string
}

type PaginationParam struct {
	Pagination bool `form:"-"`
	OnlyCount  bool `form:"-"`
	Current    int  `form:"current,default=1"`
	PageSize   int  `form:"pageSize,default=10" binding:"max=100"`
}

// UserQueryParam 查询条件
type UserQueryParam struct {
	PaginationParam
	UserName   string   `form:"userName"`   // 用户名
	QueryValue string   `form:"queryValue"` // 模糊查询
	Status     int      `form:"status"`     // 用户状态(1:启用 2:停用)
	RoleIDs    []uint64 `form:"-"`          // 角色ID列表
}

func (m *Member_dto) ToDto(aggre *aggregate.Member_aggre) {
	userRoles := aggre.GetRoleIDs()
	userdtos := []UserRoleDto{}
	for _, item := range userRoles {
		userdtos = append(userdtos, UserRoleDto{item.String(), aggre.User.ID.String()})
	}
	m.ID = aggre.User.ID.String()
	m.LoginName = aggre.User.Email
	m.Email = aggre.User.Email
	m.Userroles = userdtos
	m.Status = aggre.User.Status
	m.Password = ""
}

func (m *Member_dto) ToAggre() *aggregate.Member_aggre {
	aggre, err := aggregate.NewMember(m.LoginName, m.Email, m.Password)
	if err != nil {
		return nil
	}
	var roleids []uuid.UUID
	// if m.Userroles == nil {
	// 	m.Userroles = make([]UserRoleDto, 0)
	// }
	if len(m.Userroles[0].RoleID) > 0 {
		for _, item := range m.Userroles {
			uid, _ := uuid.Parse(item.RoleID)
			roleids = append(roleids, uid)
		}
	}
	aggre.User.LoginName = m.LoginName
	aggre.User.Email = m.Email
	aggre.User.Password = m.Password
	aggre.User.ID, _ = uuid.Parse(m.ID)
	aggre.User.Status = m.Status
	aggre.AddRoles(roleids...)
	return aggre
}
