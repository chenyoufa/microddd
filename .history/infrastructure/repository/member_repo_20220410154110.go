package repository

import (
	"context"
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

func (m *MemberDomain) Get(ctx context.Context, uuid uuid.UUID) (*aggregate.Member, error) {
	var r = new(aggregate.Member)
	var user entity.User
	var roles []*entity.Role
	var relations []valueobject.UserRoleRelation

	err1 := dbcore.GetDB(ctx).Debug().Where("user_id=?", uuid).Find(&relations).Error
	err2 := dbcore.GetDB(ctx).Debug().Model(&relations).Association("Role").Find(&roles)
	err3 := dbcore.GetDB(ctx).Debug().Model(&relations).Association("User").Find(&user)
	r.User = &user
	err4 := setUnExportedStrField(r, "roles", roles)
	err5 := setUnExportedStrField(r, "userRoleRelation", relations)
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || err5 != nil {
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
