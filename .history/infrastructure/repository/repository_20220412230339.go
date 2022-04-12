package repository

import (
	"context"
	"microddd/domain/repository"
	"microddd/infrastructure/db/dbcore"

	"gorm.io/gorm"
)

// type Repository struct {
// 	MemberRepos  memberRepos
// }

func NewRepository(cdb *gorm.DB) *repository.AuthFactory {

	return &repository.AuthFactory{

		MRepo: &memberRepos{db: dbcore.GetDB(context.Background())},
	}
}
