package model

import (
	"microddd/domain/aggregate"
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

type Model struct {
	User_po
	Role_po
	UserRole_po
}

func (ul *Model) ToDo() *aggregate.Member_aggre {

	return &aggregate.Member_aggre{}
}
