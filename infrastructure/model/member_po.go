package model

import (
	"microddd/domain/entity"
	"time"

	"github.com/google/uuid"
)

type User_po struct {
	ID        uuid.UUID `gorm:"primarykey;not null;unqua"`
	LoginName string    `gorm:"not null; "`
	Email     string    `gorm:"not null; size:30"`
	Password  string    `gorm:"not null ;size:50"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Role_po struct {
	ID        uuid.UUID `gorm:"primarykey;not null;unqua"`
	RoleName  string    `gorm:"size:20"`
	Remark    string    `gorm:"size:200"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRole_po struct {
	ID        uuid.UUID `gorm:"primarykey;not null;unqua"`
	RoleID    uuid.UUID `gorm:"not null"`
	UserID    uuid.UUID `gorm:"not null"`
	Status    int       `gorm:""`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ul *UserRole_po) ToDo() *entity.UserRoleEntity {
	return &entity.UserRoleEntity{}
}
