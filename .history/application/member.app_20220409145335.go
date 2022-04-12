package application

import "github.com/google/uuid"

type MemberApi interface {
	Login(loginName string, password string) (bool, error)
	Get(uuid uuid.UUID)

	GetList()
	AddMember(loginName string, password string) (bool, error)
	EditMember() (bool, error)

	ChangeMemberStatus(uuid uuid.UUID, status int) (bool, error)
	ChangeRoleRelationStatus(uuid uuid.UUID, status int) (bool, error)
}
