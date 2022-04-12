package application

import (
	"gomicroddd/domain/aggregate"

	"github.com/google/uuid"
)

type MemberApi interface {
	GetList()
	AddMember(loginName string, password string) (bool, error)
	EditMember(mdo aggregate.Member) (bool, error)

	Login(loginName string, password string) (bool, error)
	Get(uuid uuid.UUID)
	ChangeMemberStatus(uuid uuid.UUID, status int) (bool, error)
	ChangeRoleRelationStatus(uuid uuid.UUID, status int) (bool, error)
}
