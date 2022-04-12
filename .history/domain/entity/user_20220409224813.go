package entity

import (
	"gomicroddd/domain/valueobject"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID            uuid.UUID
	UserName      string `gorm:"size:64;uniqueIndex;default:'';not null;"` // 用户名
	RealName      string `gorm:"size:64;index;default:'';"`                // 真实姓名
	Password      string `gorm:"size:40;default:'';"`                      // 密码
	Email         string `gorm:"size:255;"`                                // 邮箱
	Phone         string `gorm:"size:20;"`                                 // 手机号
	Status        int    `gorm:"index;default:0;"`                         // 状态(1:启用 2:停用)
	Creator       uint64 `gorm:""`                                         // 创建者
	UpdateTimeAt  time.Time
	AddTimeAt     time.Time
	Remark        string
	roleRelations []valueobject.UserRoleRelation
}
