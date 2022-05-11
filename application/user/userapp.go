package user

import (
	"myddd/domain/model"
	"myddd/domain/repository"
)

type UserApp struct {
	repo repository.UserRepoer
}

func (u *UserApp) Add(usermodel *model.User) string {
	u.repo.Add()
	return "add"
}
func (u *UserApp) Del(id string) string {
	u.repo.Del()
	return "Del"
}
func (u *UserApp) Update(usermodel *model.User) string {
	u.repo.Update()
	return "Update"
}
func (u *UserApp) Query(id string) string {
	u.repo.Query()
	return "Query"
}
