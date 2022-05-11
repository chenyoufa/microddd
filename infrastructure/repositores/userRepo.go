package repositores

import "gorm.io/gorm"

type UserPo struct {
	db gorm.DB
}

// var _repo UserRepoer = UserPo{}

func (u *UserPo) Add() string {

	return " UserPo add"
}
func (u *UserPo) Del() string {
	return "UserPo Del"
}
func (u *UserPo) Update() string {
	return "UserPo Update"
}
func (u *UserPo) Query() string {
	return "UserPo Query"
}
