package app

import (
	"gomicroddd/domain/aggregate"
	"gomicroddd/domain/repository"

	"github.com/google/uuid"
)

type MemberApp struct {
	repo repository.MemberRepository
}

func NewMemberApp(repo repository.MemberRepository) *MemberApp {
	return &MemberApp{repo: repo}
}

func (mapp *MemberApp) GetList() ([]*aggregate.Member, error) {

}
func (mapi *MemberApp) AddMember(mdo *aggregate.Member) (bool, error) {

}
func (mapi *MemberApp) EditMember(mdo *aggregate.Member) (bool, error) {

}
func (mapi *MemberApp) Login(loginName string, password string) (bool, error) {

}
func (mapp *MemberApp) Get(uuid uuid.UUID) (*aggregate.Member, error) {
	model, err := mapp.repo.Get(uuid)
	if err != nil {
		return nil, err
	}
	return &model, nil
}
func (mapi *MemberApp) ChangeMemberStatus(uuid uuid.UUID, status int) (bool, error) {

}
func (mapi *MemberApp) ChangeRoleRelationStatus(uuid uuid.UUID, status int) (bool, error) {

}
