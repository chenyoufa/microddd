package entity

import "github.com/google/uuid"

type RoleMenu struct {
	RoleID   uuid.UUID // 角色ID
	MenuID   uuid.UUID // 菜单ID
	ActionID uuid.UUID // 动作ID
}
