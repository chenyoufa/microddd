package valueobject

import "github.com/google/uuid"

type UserRoleRelation struct {
	ID     uuid.UUID
	UserID uuid.UUID `gorm:"index;default:0;"` // 用户内码
	RoleID uuid.UUID `gorm:"index;default:0;"` // 角色内码
	Status int       `gorm:"index;default:0;"` // 状态
}
