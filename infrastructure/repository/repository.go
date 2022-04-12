package repository

import (
	"gomicroddd/domain/entity"
	"gomicroddd/domain/valueobject"
	"gomicroddd/infrastructure/db/dbcore"

	"gorm.io/gorm"
)

func init() {

	dbcore.RegisterInjector(func(db *gorm.DB) {
		dbcore.SetupTableModel(db, &entity.User{})
		dbcore.SetupTableModel(db, &entity.Role{})
		dbcore.SetupTableModel(db, &valueobject.UserRoleRelation{})
	})
}

type repository struct {
	MemberRepo *Merchant
	AuthCode   *AuthCode
	AuthToken  *AuthToken
}

func NewRepository() *Repository {
	var a = repository{mgo, r}
	return &Repository{
		&Merchant{a},
		&AuthCode{a},
		&AuthToken{a},
	}
}
