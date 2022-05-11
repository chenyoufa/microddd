package user

import (
	"myddd/domain/model"
	"myddd/domain/repository"
)

type UserApp struct {
	repo repository.UserRepoer
}

func (u *UserApp) Add(usermodel *model.User) string {
	return u.repo.Add(usermodel)
}
func (u *UserApp) Del(id string) string {
	return u.repo.Del(id)
}
func (u *UserApp) Update(usermodel *model.User) string {
	return u.repo.Update(usermodel)
}
func (u *UserApp) Query(id string) string {
	return u.repo.Query(id)
}
