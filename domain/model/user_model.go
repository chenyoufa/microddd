package model

import "errors"

type User struct {
	ID       int64
	Name     string
	Password string
	Nickanme string
	Status   int64
}

//工厂创建领域对象
func NewUser(name string, password string, nickanme string) (*User, error) {
	u := &User{}
	u.ID = 1
	if len(u.Name) < 2 {
		return nil, errors.New("name len letter")
	}
	u.Password = password + "125875"
	u.Nickanme = nickanme
	return u, nil
}

//user 方法
func (u *User) ChangePwd(newpwd string) bool {
	u.Password = newpwd + "125875"
	return true
}

//user 方法
func (u *User) ChangeStatus(status int64) bool {
	u.Status = status
	return true
}
