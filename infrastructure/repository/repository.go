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
	db *gorm.DB
}

type Repository struct {
	MemberRepo *memberinfra
}

func NewRepository(db *gorm.DB) *Repository {
	var a = repository{db}

	return &Repository{
		MemberRepo: a,
	}
}
