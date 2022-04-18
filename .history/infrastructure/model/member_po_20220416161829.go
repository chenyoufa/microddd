package model

import (
	"log"
	"microddd/domain/aggregate"
	"microddd/domain/entity"
	"microddd/domain/valobj"
	"time"

	tools "microddd/infrastructure/utils/tools"

	"github.com/devfeel/mapper"
	"github.com/google/uuid"
)

type Userpo struct {
	ID        uuid.UUID `gorm:"primarykey"`
	LoginName string    `gorm:"not null; "`
	Email     string    `gorm:"not null; size:30"`
	Password  string    `gorm:"not null ;size:50"`
	CreatedAt time.Time
	UpdatedAt time.Time
	UserRoles []UserRolepo `gorm:"foreignkey:UserId;association_foreignkey:ID"`
}

type Rolepo struct {
	ID        uuid.UUID `gorm:"primarykey"`
	RoleName  string    `gorm:"size:20"`
	Remark    string    `gorm:"size:200"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRolepo struct {
	ID     uuid.UUID `gorm:"primarykey"`
	Status int
	// Userpo    Userpo `gorm:"foreignkey:UserId;association_foreignkey:Id"`
	UserId    uuid.UUID
	RoleId    uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
}

type CustomerPo struct {
	User *Userpo
	// Roles     []*Role_po
	Userroles []UserRolepo
}

func init() {
	mapper.Register(&Userpo{})
	mapper.Register(&Rolepo{})
	mapper.Register(&UserRolepo{})

	mapper.Register(&entity.UserEntity{})
	// mapper.Register(&entity.RoleEntity{})
	mapper.Register(&valobj.UserRoleValObj{})
}

func (ul *Userpo) ToDo() *aggregate.Member_aggre {

	userEntity := &entity.UserEntity{}
	userRoles := make([]valobj.UserRoleValObj, 0)
	log.Println("ul.userRoles:", ul.UserRoles)

	mapper.AutoMapper(ul, userEntity)
	mapper.AutoMapper(ul.UserRoles, userRoles)

	log.Println("userEntity1:", userEntity)
	log.Println("userRoles1:", userRoles)
	return nil
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
