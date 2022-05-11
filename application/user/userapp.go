package user

import (
	"myddd/domain/model"
	"myddd/domain/repository"
)

type UserApp struct {
	Repo repository.UserRepoer
}

func (u *UserApp) Add(usermodel *model.User) string {
	return u.Repo.Add(usermodel)
}
func (u *UserApp) Del(id string) string {
	return u.Repo.Del(id)
}
func (u *UserApp) Update(usermodel *model.User) string {
	return u.Repo.Update(usermodel)
}
func (u *UserApp) Query(id string) string {
	return u.Repo.Query(id)
}
