package entity

import "github.com/google/uuid"

type UserRole_entity struct {
	ID     uuid.UUID
	RoleID uuid.UUID
	UserID uuid.UUID
	Status int
}
