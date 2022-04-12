package dto

import (
	"gomicroddd/domain/aggregate"
	"time"

	"github.com/gofrs/uuid"
)

//vo=>dto
type MemberDto struct {
	ID           uuid.UUID
	UserName     string    `json:" " ;mapper:"_uname"` // 用户名
	RealName     string    `json:""`                   // 真实姓名
	Password     string    `json:""`                   // 密码
	Email        string    `json:""`                   // 邮箱
	Phone        string    `json:""`                   // 手机号
	Status       int       `json:""`                   // 状态(1:启用 2:停用)
	Creator      uint64    `json:""`                   // 创建者
	UpdateTimeAt time.Time `json:""`
	AddTimeAt    time.Time `json:""`
	Remark       string    `json:""`
}

func (m *MemberDto) DtoToDo() *aggregate.Member {
	return &aggregate.Member{}
}
