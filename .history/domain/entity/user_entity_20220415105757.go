package entity

import "github.com/google/uuid"

type UserEntity struct {
	ID        uuid.UUID
	LoginName string
	Email     string
	Password  string
}
