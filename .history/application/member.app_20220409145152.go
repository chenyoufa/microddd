package application

import "github.com/google/uuid"

type MemberApi interface {
	Login(loginName string, password string) (bool, error)
	GetList()
	Get(uuid uuid.UUID)
	AddMember(loginName string, password string) (bool, error)
	EditMember() (bool, error)
	ChangeMemberStatus() (bool, error)
	ChangeRoleRelationStatus() (bool, error)
}
