package entity

import (
	"time"

	"github.com/google/uuid"
)

type Menu struct {
	ID         uuid.UUID `gorm:"primaryKey;"`
	Name       string    `gorm:"size:50;index;default:'';not null;"` // 菜单名称
	Icon       *string   `gorm:"size:255;"`                          // 菜单图标
	Router     *string   `gorm:"size:255;"`                          // 访问路由
	ParentID   *uint64   `gorm:"index;default:0;"`                   // 父级内码
	ParentPath *string   `gorm:"size:512;index;default:'';"`         // 父级路径
	IsShow     int       `gorm:"index;default:0;"`                   // 是否显示(1:显示 2:隐藏)
	Status     int       `gorm:"index;default:0;"`                   // 状态(1:启用 2:禁用)
	Sequence   int       `gorm:"index;default:0;"`                   // 排序值
	Memo       *string   `gorm:"size:1024;"`                         // 备注
	Creator    uint64    `gorm:""`                                   // 创建人
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type MenuAction struct {
	ID        uuid.UUID `gorm:"primaryKey;"`
	MenuID    uint64    `gorm:"index;not null;"` // 菜单ID
	Code      string    `gorm:"size:100;"`       // 动作编号
	Name      string    `gorm:"size:100;"`       // 动作名称
	CreatedAt time.Time
	UpdatedAt time.Time
}
type MenuActionResource struct {
	ID        uuid.UUID `gorm:"primaryKey;"`
	ActionID  uint64    `gorm:"index;not null;"` // 菜单动作ID
	Method    string    `gorm:"size:50;"`        // 资源请求方式(支持正则)
	Path      string    `gorm:"size:255;"`       // 资源请求路径（支持/:id匹配）
	CreatedAt time.Time
	UpdatedAt time.Time
}
