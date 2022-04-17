package repository

import (
	"log"
	"microddd/domain/repository"
	"microddd/infrastructure/db/dbcore"
	"microddd/infrastructure/model"

	"gorm.io/gorm"
)

// type Repository struct {
// 	MemberRepos memberRepos
// }

func NewRepository(cdb *gorm.DB) *repository.AuthFactory {

	db := dbcore.GetDBConfig()
	if db.AutoMigrate {
		dbMigrate(cdb)
	}
	return &repository.AuthFactory{

		MRepo: &memberRepos{db: cdb},
	}
}

func dbMigrate(db *gorm.DB) *gorm.DB {

	db.AutoMigrate(&model.Userpo{}, &model.Rolepo{}, &model.UserRolepo)
	log.Println("Schema migration has been procceed")

	return db
}
