package model

import "github.com/google/uuid"

type User_po struct {
	ID        uuid.UUID
	LoginName string
	Email     string
	Password  string
}

type Role_po struct {
	ID       uuid.UUID
	RoleName string
	Remark   string
}

type UserRole_po struct {
	ID     uuid.UUID
	RoleID uuid.UUID
	UserID uuid.UUID
	Status int
}
