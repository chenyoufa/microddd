package application

type MemberApi interface {
	Login(loginName string, password string) (bool, error)
	GetList()
	Get()
	AddMember() (bool, error)
	EditMember() (bool, error)
	ChangeMemberStatus() (bool, error)
	ChangeRoleRelationStatus() (bool, error)
}
