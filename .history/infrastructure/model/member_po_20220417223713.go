package model

import (
	"microddd/domain/aggregate"
	"microddd/domain/entity"
	"microddd/domain/valobj"
	"time"

	"github.com/devfeel/mapper"
)

type Userpo struct {
	ID        string `gorm:"primarykey;size:64"`
	LoginName string `gorm:"not null; "`
	Email     string `gorm:"not null; size:30"`
	Password  string `gorm:"not null ;size:50"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Rolepos   []*Rolepo `gorm:"many2many:user_Rolepos;"`
}

type Rolepo struct {
	ID        string    `gorm:"primarykey;"`
	RoleName  string    `gorm:"size:20"`
	Remark    string    `gorm:"size:200"`
	Userpo    []*Userpo `gorm:"many2many:user_Rolepos;"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// type UserRolepo struct {
// 	ID        string `gorm:"primarykey;size:64"`
// 	Status    int
// 	Rolepo    Userpo `gorm:"foreignkey:RoleID;association_foreignkey:ID"`
// 	UserID    string `gorm:"foreignkey:UserID;association_foreignkey:ID"`
// 	RoleID    string `gorm:"size:64"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
// }

type CustomerPo struct {
	User *Userpo
	// Roles     []*Role_po
	// Userroles []UserRolepo
}

func init() {
	mapper.Register(&Userpo{})
	mapper.Register(&Rolepo{})
	// mapper.Register(&UserRolepo{})

	mapper.Register(&entity.UserEntity{})
	// mapper.Register(&entity.RoleEntity{})
	mapper.Register(&valobj.UserRoleValObj{})
}

func (ul *Userpo) ToDo() *aggregate.Member_aggre {

	userEntity := &entity.UserEntity{}

	// userRoles := []valobj.UserRoleValObj{}
	mapper.AutoMapper(ul, userEntity)
	// for _, item := range ul.Rolepos {
	// 	temp := valobj.UserRoleValObj{RoleID: item.ID, UserID: item.UserID}
	// 	userRoles = append(userRoles, temp)
	// }

	// log.Println("userEntity1:", userEntity)
	// log.Println("userRoles1:", userRoles)

	rmodel := &aggregate.Member_aggre{
		User: userEntity,
	}
	// setUnExportedStrField(rmodel, "roles", roles)
	// tools.SetUnExportedStrField(rmodel, "userroles", userRoles)
	// log.Println("rmodel:", *rmodel.User)
	return rmodel
}

func (ul *CustomerPo) ToPo(aggre *aggregate.Member_aggre) {
	userAggre := aggre.User
	// userRoles := tools.GetUnExportedField(aggre, "userroles")
	// roleids := aggre.GetRoleIDs()
	// uluser := &Userpo{}
	// ulroles := []Rolepo{}

	// if len(roleids) > 0 {
	// 	for _, item := range roleids {
	// 		temp := Rolepo{UserID: userAggre.ID, RoleID: item}
	// 		// 	ulroles = append(ulroles, temp)
	// 	}
	// 	// ul.Userroles = ulroles
	// }
	// mapper.AutoMapper(userAggre, uluser)

	ul.User = &Userpo{
		ID:        userAggre.ID,
		LoginName: userAggre.LoginName,
		Email:     userAggre.Email,
		Password:  userAggre.Email,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	// mapper.AutoMapper(userRoles, ulrole)
	// log.Println("userRoles:", ul)

}
