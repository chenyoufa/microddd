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
	err1 := dbcore.GetDB(ctx, u.db).Where("id=?", uuid).Preload("UserRoles").Find(&user).Error

	// err2 := dbcore.GetDB(ctx, u.db).Model(&relations).Association("Role").Find(&roles)
	// err3 := dbcore.GetDB(ctx, u.db).Model(&relations).Association("User").Find(&user)

	// log.Println("user:", user)
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
	var customerPo = &model.CustomerPo{}
	customerPo.ToPo(aggre)

	var err error
	log.Println("customerPo:", customerPo.Userroles)

	// uid := uuid.New().String()
	// userpo := model.Userpo{
	// 	ID:        uid,
	// 	LoginName: "fage22",
	// 	Email:     "79756530@qq.com",
	// 	Password:  "12548",
	// 	UpdatedAt: time.Now(),
	// 	CreatedAt: time.Now(),
	// 	// UserRoles: nil,
	// }
	// log.Println(userpo)

	roles := dbcore.GetDB(ctx, u.db).Where("User_ID=?", customerPo.User.ID)

	dbcore.Transaction(ctx, u.db, func(txctx context.Context) error {

		// dbcore.GetDB(ctx, u.db).Delete(roles)
		customerPo.User.UserRolepos = append(customerPo.User.UserRolepos, roles)
		err = dbcore.GetDB(ctx, u.db).Debug().Create(customerPo.User).Error

		return err
	})
	// dbcore.Transaction(ctx, u.db, func(txctx context.Context) error {

	// 	err = dbcore.GetDB(ctx, u.db).Create(customerPo.Userroles).Error
	// 	return err
	// })
	if err != nil {
		log.Println(err)
		return false, err
	}
	return true, nil
}
func (u *memberRepos) Edit(ctx context.Context, aggre *aggregate.Member_aggre) (bool, error) {
	var customerPo = &model.CustomerPo{}
	customerPo.ToPo(aggre)
	var err error
	err = dbcore.GetDB(ctx, u.db).Updates(customerPo.User).Error
	// dbcore.Transaction(ctx, u.db, func(txctx context.Context) error {
	// 	// err = dbcore.GetDB(ctx, u.db).Updates(customerPo.Userroles).Error
	// 	return err
	// })
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
