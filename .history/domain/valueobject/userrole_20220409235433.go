package valueobject

import (
	"gomicroddd/domain/entity"

	"github.com/google/uuid"
)

type UserRoleRelation struct {
	ID     uuid.UUID   `gorm:"primarykey"`
	User   entity.User `gorm:"foreignkey:UserID"` // 使用 UserID 作为外键
	Role   entity.Role `gorm:"foreignkey:RoleID"` // 使用 RoleID 作为外键
	UserID uuid.UUID   `gorm:"index;default:0;"`  // 用户内码
	RoleID uuid.UUID   `gorm:"index;default:0;"`  // 角色内码
	Status int         `gorm:"index;default:0;"`  // 状态
}
