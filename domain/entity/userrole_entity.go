package entity

import "github.com/google/uuid"

type UserRoleEntity struct {
	ID     uuid.UUID
	RoleID uuid.UUID
	UserID uuid.UUID
	Status int
}
