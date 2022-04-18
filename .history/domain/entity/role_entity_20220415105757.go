package entity

import "github.com/google/uuid"

type RoleEntity struct {
	ID       uuid.UUID
	RoleName string
	Remark   string
}
