package repository

import (
	"context"
	"log"
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
	var user model.Userpo
	// var roles []*model.Role_po
	// var relations []model.UserRolepo
	err1 := dbcore.GetDB(ctx, u.db).Where("id=?", uuid).Preload("UserRolepos").Find(&user).Error

	// err2 := dbcore.GetDB(ctx, u.db).Model(&relations).Association("Role").Find(&roles)
	// err3 := dbcore.GetDB(ctx, u.db).Model(&relations).Association("User").Find(&user)

	for _, item := range user.UserRolepos {
		log.Println("aggre:", item)
	}
	// log.Println("relations:", relations)
	if err1 != nil {
		return nil, nil
	}
	// cmpo := model.CustomerPo{User: &user, Userroles: relations}
	cmdo := user.ToDo()
	return cmdo, err
}

func (u *memberRepos) GetList(ctx context.Context, uuid uuid.UUID) ([]*aggregate.Member_aggre, error) {
	var users []model.Userpo
	var cmdos []*aggregate.Member_aggre
	err := dbcore.GetDB(ctx, u.db).Where("user_id=?", uuid).Preload("UserRoles").Find(&users)
	for _, user := range users {
		// userrole := user.UserRoles

		// cmpo := model.CustomerPo{User: &user, Userroles: userrole}
		cmdo := user.ToDo()
		cmdos = append(cmdos, cmdo)
	}
	if err != nil {
		return nil, err.Error
	}
	return cmdos, nil
}
func (u *memberRepos) Add(ctx context.Context, aggre *aggregate.Member_aggre) (bool, error) {
	var userpo = &model.Userpo{}
	userpo.ToPo(aggre)
	var err error
	userroles := model.UserRolepo{}
	dbcore.GetDB(ctx, u.db).Where("user_id=?", userpo.ID).Find(&userroles)
	dbcore.Transaction(ctx, u.db, func(txctx context.Context) error {
		dbcore.GetDB(ctx, u.db).Delete(userroles)
		err = dbcore.GetDB(ctx, u.db).Debug().Create(userpo).Error
		return err
	})
	if err != nil {
		log.Println(err)
		return false, err
	}
	return true, nil
}
func (u *memberRepos) Edit(ctx context.Context, aggre *aggregate.Member_aggre) (bool, error) {
	var userpo = &model.Userpo{}
	userpo.ToPo(aggre)
	var err error

	userroles := []*model.UserRolepo{}

	// dbcore.GetDB(ctx, u.db).Where("user_id=? and role_id in (?)", userpo.ID, aggre.GetRoleIDs()).Find(&userroles)

	dbcore.Transaction(ctx, u.db, func(txctx context.Context) error {
		// err = dbcore.GetDB(ctx, u.db).Updates(customerPo.Userroles).Error

		dbcore.GetDB(ctx, u.db).Delete(userroles)
		userpo.UserRolepos = userroles
		err = dbcore.GetDB(ctx, u.db).Updates(userpo).Error
		return err
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func (u *memberRepos) Login(ctx context.Context, usname string, pwd string) (bool, error) {
	var count int64
	err := dbcore.GetDB(ctx, u.db).Where(" (username=? or email=?)  and password=?", usname, pwd).Find(&model.Userpo{}).Count(&count).Error
	return count > 0, err
}
