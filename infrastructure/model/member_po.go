package model

import (
	"microddd/domain/aggregate"
	"microddd/domain/entity"
	"microddd/domain/valobj"
	"time"

	tools "microddd/infrastructure/utils/tools"

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
	UserRoles []UserRole_po
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
	User_po   User_po   `gorm:"foreignkey:UserID"`
	Role_po   Role_po   `gorm:"foreignkey:UserRoID"`
	RoleID    uuid.UUID `gorm:"not null"`
	UserID    uuid.UUID `gorm:"not null"`
	Status    int       `gorm:""`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CustomerPo struct {
	User *User_po
	// Roles     []*Role_po
	Userroles []UserRole_po
}

func init() {
	mapper.Register(&User_po{})
	mapper.Register(&Role_po{})
	mapper.Register(&UserRole_po{})

	mapper.Register(&entity.UserEntity{})
	// mapper.Register(&entity.RoleEntity{})
	mapper.Register(&valobj.UserRoleValObj{})
}

func (ul *CustomerPo) ToDo() *aggregate.Member_aggre {

	userEntity := &entity.UserEntity{}
	// roles := &[]*entity.RoleEntity{}
	userRoles := &[]*valobj.UserRoleValObj{}

	mapper.AutoMapper(ul.User, userEntity)
	// mapper.AutoMapper(ul.Roles, roles)
	mapper.AutoMapper(ul.Userroles, userRoles)
	rmodel := &aggregate.Member_aggre{
		User: userEntity,
	}
	// setUnExportedStrField(rmodel, "roles", roles)
	tools.SetUnExportedStrField(rmodel, "userroles", userRoles)

	return rmodel
}

func (ul *CustomerPo) ToPo(aggre *aggregate.Member_aggre) {
	userEntity := aggre.User
	userRoles := tools.GetUnExportedField(aggre, "userroles")
	mapper.AutoMapper(userEntity, ul.User)
	mapper.AutoMapper(userRoles, ul.Userroles)
}
