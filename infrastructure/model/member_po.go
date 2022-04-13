package model

import (
	"microddd/domain/aggregate"
	"microddd/domain/entity"
	"time"

	"github.com/devfeel/mapper"
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

type CustomerPo struct {
	User_po
	Role_po
	UserRole_po
}

func init() {
	mapper.Register(&User_po{})
	mapper.Register(&Role_po{})
	mapper.Register(&UserRole_po{})

	mapper.Register(&entity.UserEntity{})
	mapper.Register(&entity.RoleEntity{})
	mapper.Register(&entity.UserRoleEntity{})
}

func (ul *CustomerPo) ToDo() *aggregate.Member_aggre {

	userEntity := &entity.UserEntity{}
	roleEntity := &entity.RoleEntity{}
	userRoleEntity := &entity.UserRoleEntity{}

	mapper.AutoMapper(ul.User_po, userEntity)
	mapper.AutoMapper(ul.User_po, roleEntity)
	mapper.AutoMapper(ul.User_po, userRoleEntity)
	rmodel := &aggregate.Member_aggre{
		User: userEntity,
	}

	return rmodel
}
