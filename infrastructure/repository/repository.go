package repository

import (
	"microddd/domain/repository"

	"gorm.io/gorm"
)

// type Repository struct {
// 	MemberRepos memberRepos
// }

func NewRepository(cdb *gorm.DB) *repository.AuthFactory {

	return &repository.AuthFactory{

		MRepo: &memberRepos{db: cdb},
	}
}
