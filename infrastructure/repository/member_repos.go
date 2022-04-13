package repository

import (
	"context"
	"microddd/domain/aggregate"

	"microddd/infrastructure/db/dbcore"
	"microddd/infrastructure/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// var memberReposSet = wire.NewSet(
// 	wire.Struct(new(memberRepos)),
// 	wire.Bind(new(repository.MemberRepoer), new(*memberRepos)))

type memberRepos struct {
	db *gorm.DB
}

// func NewmemberRepos(db *gorm.DB) repository.MemberRepo {
// 	return &memberRepos{db: db}
// }

func (u *memberRepos) Get(ctx context.Context, uuid uuid.UUID) (*aggregate.Member_aggre, error) {
	userpo := model.User_po{}
	var err error
	dbcore.Transaction(ctx, u.db, func(txctx context.Context) error {
		err = u.db.Where("id=?", uuid).Find(&userpo).Error

		return err
	})
	userdo := &aggregate.Member_aggre{}
	return userdo, err
}
