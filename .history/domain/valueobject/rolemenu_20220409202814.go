package valueobject

import "gomicroddd/domain/entity"

type RoleMenuRelation struct {
	User entity.User `gorm:"foreignkey:RoleID"` // 使用 UserID 作为外键
	Role entity.Role `gorm:"foreignkey:MenuID"` // 使用 RoleID 作为外键

	RoleID   uint64 `gorm:"index;not null;"` // 角色ID
	MenuID   uint64 `gorm:"index;not null;"` // 菜单ID
	ActionID uint64 `gorm:"index;not null;"` // 动作ID
}
