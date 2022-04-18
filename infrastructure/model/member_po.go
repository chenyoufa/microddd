package model

import (
	"log"
	"microddd/domain/aggregate"
	"microddd/domain/entity"
	"microddd/domain/valobj"
	"microddd/infrastructure/utils/tools"
	"time"

	"github.com/google/uuid"
)

type Userpo struct {
	ID          string        `gorm:"primarykey;size:64;"`
	LoginName   string        `gorm:"not null; "`
	Email       string        `gorm:"not null; size:30"`
	Password    string        `gorm:"not null ;size:50"`
	CreatedAt   time.Time     `gorm:"uint8"`
	UpdatedAt   time.Time     `gorm:"uint8"`
	UserRolepos []*UserRolepo `gorm:"foreignkey:UserID"`
	Status      int           `gorm:"int;default:0"`
}

type Rolepo struct {
	ID          string        `gorm:"primarykey;size:64;"`
	RoleName    string        `gorm:"size:20"`
	Remark      string        `gorm:"size:200"`
	UserRolepos []*UserRolepo `gorm:"foreignkey:RoleID"`
	CreatedAt   time.Time     `gorm:"uint8"`
	UpdatedAt   time.Time     `gorm:"uint8"`
	Status      int           `gorm:"int;default:0"`
}

type UserRolepo struct {
	ID     string `gorm:"primarykey;;"`
	UserID string `gorm:"size:64"`
	RoleID string `gorm:"size:64"`
}

func (ul *Userpo) ToDo() *aggregate.Member_aggre {
	uid, _ := uuid.Parse(ul.ID)
	userEntity := entity.UserEntity{
		ID:        uid,
		LoginName: ul.LoginName,
		Email:     ul.Email,
		Password:  ul.Password,
		Status:    ul.Status,
	}
	userRoles := []valobj.UserRoleValObj{}
	rmodel := &aggregate.Member_aggre{}
	if len(ul.UserRolepos) > 0 {
		for _, item := range ul.UserRolepos {
			tuserid, _ := uuid.Parse(item.UserID)
			troleid, _ := uuid.Parse(item.RoleID)
			temp := valobj.UserRoleValObj{RoleID: tuserid, UserID: troleid}
			userRoles = append(userRoles, temp)
		}
		tools.SetUnExportedStrField(rmodel, "userroles", userRoles)
	}
	rmodel.User = &userEntity
	return rmodel
}

func (ul *Userpo) ToPo(aggre *aggregate.Member_aggre) {
	userAggre := aggre.User
	roleids := aggre.GetRoleIDs()
	ulroles := []*UserRolepo{}
	log.Println("roleids:", roleids)

	if len(roleids) > 0 {
		for _, item := range roleids {
			temp := UserRolepo{UserID: userAggre.ID.String(), RoleID: item.String()}
			ulroles = append(ulroles, &temp)
		}
	}
	*ul = Userpo{
		ID:          userAggre.ID.String(),
		LoginName:   userAggre.LoginName,
		Email:       userAggre.Email,
		Password:    userAggre.Password,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		UserRolepos: ulroles,
		Status:      userAggre.Status,
	}
}
