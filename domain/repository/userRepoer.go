package repository

import "myddd/domain/model"

type UserRepoer interface {
	Add(usermodel *model.User) string
	Del(id string) string
	Update(usermodel *model.User) string
	Query(id string) string
}
