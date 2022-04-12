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
	var roles []entity.Role
	var relations []valueobject.UserRoleRelation
	var err error
	err = dbcore.GetDB(ctx).Debug().Where("user_id=?", uuid).Find(&relations).Error
	err = dbcore.GetDB(ctx).Debug().Model(&relations).Association("Role").Find(&roles)
	err = dbcore.GetDB(ctx).Debug().Model(&relations).Association("User").Find(&user)
	r.User = &user
	err = setUnExportedStrField(r, "roles", roles)
	err = setUnExportedStrField(r, "userRoleRelation", relations)
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
