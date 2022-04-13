package entity

import "github.com/google/uuid"

type UserRoleEntity struct {
	RoleID uuid.UUID
	UserID uuid.UUID
}
