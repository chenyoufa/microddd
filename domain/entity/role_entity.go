package entity

import "github.com/google/uuid"

type Role_entity struct {
	ID       uuid.UUID
	RoleName string
	Remark   string
}
