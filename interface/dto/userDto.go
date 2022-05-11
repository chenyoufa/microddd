package dto

import "myddd/domain/model"

type UserDto struct {
	ID       int64
	Name     string
	Nickanme string
	Status   int64
}

func (u *UserDto) ToDto(do model.User) {
	u.ID = do.ID
	u.Name = do.Name
	u.Nickanme = do.Nickanme
	u.Status = do.Status
}
func (u *UserDto) ToDo() *model.User {
	m := &model.User{}
	m.ID = u.ID
	m.Name = u.Name
	m.Nickanme = u.Nickanme
	m.Status = u.Status
	return m
}
