package repository

import (
	"context"
	"microddd/domain/aggregate"

	"microddd/infrastructure/db/dbcore"
	"microddd/infrastructure/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type memberRepos struct {
	db *gorm.DB
}

func (u *memberRepos) Get(ctx context.Context, uuid uuid.UUID) (*aggregate.Member_aggre, error) {
	var err error
	// dbcore.Transaction(ctx, u.db, func(txctx context.Context) error {
	// 	err = u.db.Where("id=?", uuid).Find(&userpo).Error

	// 	return err
	// })
	var user model.User_po
	var roles []*model.Role_po
	var relations []*model.UserRole_po
	err1 := dbcore.GetDB(ctx, u.db).Where("user_id=?", uuid).Find(&relations).Error
	err2 := dbcore.GetDB(ctx, u.db).Model(&relations).Association("Role").Find(&roles)
	err3 := dbcore.GetDB(ctx, u.db).Model(&relations).Association("User").Find(&user)
	if err1 != nil || err2 != nil || err3 != nil {
		return nil, nil
	}
	cmpo := model.CustomerPo{User: user, Roles: roles, Userroles: relations}
	cmdo := cmpo.ToDo()
	return cmdo, err
}
