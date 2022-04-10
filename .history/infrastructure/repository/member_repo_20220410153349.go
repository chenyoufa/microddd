package repository

import (
	"context"
	"fmt"
	"gomicroddd/domain/aggregate"
	"gomicroddd/domain/entity"
	"gomicroddd/domain/valueobject"
	"gomicroddd/infrastructure/db/dbcore"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func init() {

	dbcore.RegisterInjector(func(db *gorm.DB) {
		dbcore.SetupTableModel(db, &entity.User{})
		dbcore.SetupTableModel(db, &entity.Role{})
		dbcore.SetupTableModel(db, &valueobject.UserRoleRelation{})
	})
}

type MemberDomain struct{}

// func init() {
// 	mapper.Register(&.MemberDto{})
// 	mapper.Register(&entity.User{})
// }

func (m *MemberDomain) Get(ctx context.Context, uuid uuid.UUID) (*aggregate.Member, error) {
	var r *aggregate.Member
	// var user Usert
	// profile := Profile{}

	var user entity.User
	var roles []entity.Role
	var relations []valueobject.UserRoleRelation

	err := dbcore.GetDB(ctx).Debug().Where("user_id=?", uuid).Find(&relations).Error
	dbcore.GetDB(ctx).Debug().Model(&relations).Association("Role").Find(&roles).Error()
	dbcore.GetDB(ctx).Debug().Model(&relations).Association("User").Find(&user).Error()
	// db = db.Preload("user_id").Find(&user)
	// db.Preload("role").Find(&roles)
	//
	// if db != nil {

	// }
	fmt.Println(user)
	fmt.Println(roles)
	for _, item := range relation {
		fmt.Println(item.Role)
	}

	// err = dbcore.GetDB(ctx).Find(&roles).Error
	// err = dbcore.GetDB(ctx).Find(&relations).Error
	// err = setUnExportedStrField(r, "roles", roles)
	// err = setUnExportedStrField(r, "userRoleRelation", relations)
	var err error
	if err != nil {
		return nil, nil
	}
	return r, nil
}
func (m *MemberDomain) Add(nm aggregate.Member) error {
	return nil
}
func (m *MemberDomain) Update(aggregate.Member) error {
	return nil
}