package valobj

import "github.com/google/uuid"

type Action struct {
	MenuID uuid.UUID // 菜单ID
	Code   string    // 动作编号
	Name   string    // 动作名称
}
