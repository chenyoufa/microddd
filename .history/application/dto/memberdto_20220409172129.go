package dto

import (
	"time"

	"github.com/gofrs/uuid"
)

type memberDto struct {
	ID           uuid.UUID
	userName     string `json:""`                          // 用户名
	realName     string `gorm:"size:64;index;default:'';"` // 真实姓名
	password     string `gorm:"size:40;default:'';"`       // 密码
	email        string `gorm:"size:255;"`                 // 邮箱
	phone        string `gorm:"size:20;"`                  // 手机号
	status       int    `gorm:"index;default:0;"`          // 状态(1:启用 2:停用)
	creator      uint64 `gorm:""`                          // 创建者
	updateTimeAt time.Time
	addTimeAt    time.Time
	remark       string
}
