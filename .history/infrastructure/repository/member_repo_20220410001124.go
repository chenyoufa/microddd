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

type Po struct {
	entity.User
	roles []entity.Role
}

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
	var user entity.User
	var roles []entity.Role
	var relations []valueobject.UserRoleRelation
	var po Po

	err := dbcore.GetDB(ctx).Debug().Preload("Po.User").Preload("Po.Roles").Find(&po)

	fmt.Println(user)
	fmt.Println(relations)
	fmt.Println(roles)
	// err = dbcore.GetDB(ctx).Find(&roles).Error
	// err = dbcore.GetDB(ctx).Find(&relations).Error
	// err = setUnExportedStrField(r, "roles", roles)
	// err = setUnExportedStrField(r, "userRoleRelation", relations)
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
