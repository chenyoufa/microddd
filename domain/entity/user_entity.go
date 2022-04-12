package entity

import "github.com/google/uuid"

type User_entity struct {
	ID        uuid.UUID
	LoginName string
	Email     string
	Password  string
}
