package entity

import (
	"time"

	"github.com/google/uuid"
)

type Role struct {
	ID           uuid.UUID
	Name         string  `gorm:"size:100;index;default:'';not null;"` // 角色名称
	Sequence     int     `gorm:"index;default:0;"`                    // 排序值
	Memo         *string `gorm:"size:1024;"`                          // 备注
	Status       int     `gorm:"index;default:0;"`                    // 状态(1:启用 2:禁用)
	Creator      uint64  `gorm:""`                                    // 创建者
	UpdateTimeAt time.Time
	AddTimeAt    time.Time
}
