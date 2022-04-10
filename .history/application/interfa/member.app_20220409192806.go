package interfa

import (
	"gomicroddd/domain/aggregate"

	"github.com/google/uuid"
)

type MemberApi interface {
	GetList()
	AddMember(mdo *aggregate.Member) (bool, error)
	EditMember(mdo *aggregate.Member) (bool, error)
	Login(loginName string, password string) (bool, error)
	Get(uuid uuid.UUID)
	ChangeMenbStatus(uuid uuid.UUID, status int) (bool, error)
	AddRole() (bool, error)
	RemoveRole() (bool, error)
	ChangeUsRoleStatus(uuid uuid.UUID, status int) (bool, error)
}
