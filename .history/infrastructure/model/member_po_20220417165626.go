package model

import (
	"log"
	"microddd/domain/aggregate"
	"microddd/domain/entity"
	"microddd/domain/valobj"
	"time"

	tools "microddd/infrastructure/utils/tools"

	"github.com/devfeel/mapper"
)

type Userpo struct {
	ID        string `gorm:"primarykey;size:64"`
	LoginName string `gorm:"not null; "`
	Email     string `gorm:"not null; size:30"`
	Password  string `gorm:"not null ;size:50"`
	CreatedAt time.Time
	UpdatedAt time.Time
	// UserRoles []UserRolepo `gorm:"foreignkey:UserID;association_foreignkey:ID"`
}

type Rolepo struct {
	ID        string `gorm:"primarykey;size:64"`
	RoleName  string `gorm:"size:20"`
	Remark    string `gorm:"size:200"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRolepo struct {
	ID        string `gorm:"primarykey;size:64"`
	Status    int
	Rolepo    Userpo `gorm:"foreignkey:RoleID;association_foreignkey:ID"`
	UserID    string `gorm:"size:64"`
	RoleID    string `gorm:"size:64"`
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
	userRoles := []valobj.UserRoleValObj{}
	mapper.AutoMapper(ul, userEntity)
	// for _, item := range ul.UserRoles {
	// 	temp := valobj.UserRoleValObj{RoleID: item.RoleID, UserID: item.UserID}
	// 	userRoles = append(userRoles, temp)
	// }

	// log.Println("userEntity1:", userEntity)
	// log.Println("userRoles1:", userRoles)

	rmodel := &aggregate.Member_aggre{
		User: userEntity,
	}
	// setUnExportedStrField(rmodel, "roles", roles)
	tools.SetUnExportedStrField(rmodel, "userroles", userRoles)
	// log.Println("rmodel:", *rmodel.User)
	return rmodel
}

func (ul *CustomerPo) ToPo(aggre *aggregate.Member_aggre) {
	userAggre := aggre.User
	// userRoles := tools.GetUnExportedField(aggre, "userroles")
	roleids := aggre.GetRoleIDs()
	uluser := &Userpo{}
	ulroles := []UserRolepo{}

	if len(roleids) > 0 {
		for _, item := range roleids {
			temp := UserRolepo{UserID: userAggre.ID.String(), RoleID: item.String()}
			ulroles = append(ulroles, temp)
		}
		ul.Userroles = ulroles
	}

	mapper.AutoMapper(userAggre, uluser)
	log.Println("userAggre:", userAggre)
	return
	ul.User = uluser

	// mapper.AutoMapper(userRoles, ulrole)
	// log.Println("userRoles:", ul)

}
