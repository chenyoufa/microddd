package repository

import (
	"context"
	"microddd/domain/aggregate"

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
	rolepo := model.Role_po{}
	userrolepo := model.UserRole_po{}
	var err error
	// dbcore.Transaction(ctx, u.db, func(txctx context.Context) error {
	// 	err = u.db.Where("id=?", uuid).Find(&userpo).Error

	// 	return err
	// })
	err = u.db.Where("id=?", uuid).Find(&userpo).Error
	err = u.db.Where("id=?", uuid).Find(&userpo).Error
	err = u.db.Where("id=?", uuid).Find(&userpo).Error

	cmpo := model.CustomerPo{userpo, rolepo, userrolepo}
	cmdo := cmpo.ToDo()
	return cmdo, err
}
