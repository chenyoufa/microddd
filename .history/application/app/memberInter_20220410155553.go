package interfa

import (
	"gomicroddd/domain/aggregate"

	"github.com/google/uuid"
)

type MemberApi interface {
	Get(uuid uuid.UUID) (*aggregate.Member, error)

	GetList() ([]*aggregate.Member, error)

	AddMember(mdo *aggregate.Member) (bool, error)
	EditMember(mdo *aggregate.Member) (bool, error)
	Login(loginName string, password string) (bool, error)

	ChangeMenbStatus(uuid uuid.UUID, status int) (bool, error)
	AddRole() (bool, error)
	RemoveRole() (bool, error)
	DongjieRoleStatus(uuid uuid.UUID, status int) (bool, error)
}
