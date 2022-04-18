package valobj

import "github.com/google/uuid"

type UserRoleValObj struct {
	RoleID uuid.UUID
	UserID uuid.UUID
}
