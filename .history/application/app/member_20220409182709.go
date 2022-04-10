package app

import (
	"gomicroddd/domain/aggregate"

	"github.com/gofrs/uuid"
)

type MemberApi struct {
}

func (mapi *MemberApi) GetList() {

}
func (mapi *MemberApi) AddMember(mdo *aggregate.Member) (bool, error) {

}
func (mapi *MemberApi) EditMember(mdo *aggregate.Member) (bool, error) {

}
func (mapi *MemberApi) Login(loginName string, password string) (bool, error) {

}
func (mapi *MemberApi) Get(uuid uuid.UUID) {

}
func (mapi *MemberApi) ChangeMemberStatus(uuid uuid.UUID, status int) (bool, error) {

}
func (mapi *MemberApi) ChangeRoleRelationStatus(uuid uuid.UUID, status int) (bool, error) {

}
