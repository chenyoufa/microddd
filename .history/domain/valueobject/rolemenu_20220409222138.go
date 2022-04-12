package valueobject

type RoleMenuRelation struct {
	// Role     entity.Role `gorm:"foreignkey:RoleID"` // 使用 RoleID 作为外键
	// Menu     entity.Menu `gorm:"foreignkey:MenuID"` // 使用 MenuID 作为外键
	RoleID   uint64 `gorm:"index;not null;"` // 角色ID
	MenuID   uint64 `gorm:"index;not null;"` // 菜单ID
	ActionID uint64 `gorm:"index;not null;"` // 动作ID
}
