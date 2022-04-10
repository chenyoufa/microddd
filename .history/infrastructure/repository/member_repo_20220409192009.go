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

// func init() {
// 	mapper.Register(&.MemberDto{})
// 	mapper.Register(&entity.User{})
// }

func (m *MemberDomain) Get(ctx context.Context, uuid uuid.UUID) (aggregate.Member, error) {
	var r *aggregate.Member
	var user entity.User
	err := dbcore.GetDB(ctx).Find(&user).Error

}
func (m *MemberDomain) Add(nm aggregate.Member) error {

}
func (m *MemberDomain) Update(aggregate.Member) error {

}
