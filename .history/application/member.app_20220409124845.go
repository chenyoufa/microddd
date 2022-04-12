package application

type MemberApi interface {
	Login(loginName string, password string) (bool, error)
	AddMember() (bool, error)
	EditMember() (bool, error)
}
