package valueobject

type RoleMenuRelation struct {
	RoleID   uint64 `gorm:"index;not null;"` // 角色ID
	MenuID   uint64 `gorm:"index;not null;"` // 菜单ID
	ActionID uint64 `gorm:"index;not null;"` // 动作ID
}
