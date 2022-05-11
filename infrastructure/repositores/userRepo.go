package repositores

import (
	"myddd/domain/model"

	"gorm.io/gorm"
)

type UserPo struct {
	DB gorm.DB
}

// var _repo UserRepoer = UserPo{}
// var _ repository.UserRepoer = &UserPo{}

func (u *UserPo) Add(usermodel *model.User) string {

	return " UserPo add"
}
func (u *UserPo) Del(id string) string {
	return "UserPo Del"
}
func (u *UserPo) Update(usermodel *model.User) string {
	return "UserPo Update"
}
func (u *UserPo) Query(id string) string {
	return "UserPo Query"
}
