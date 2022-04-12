package app

import (
	"gomicroddd/domain/aggregate"
	"gomicroddd/domain/repository"

	"github.com/google/uuid"
)

type memberApp struct {
	repo repository.MemberRepository
}

func NewMemberApp() *memberApp {
	return &memberApp{}
}
func (mapp *memberApp) GetList() ([]*aggregate.Member, error) {
	return nil, nil
}
func (mapi *memberApp) AddMember(mdo *aggregate.Member) (bool, error) {
	return true, nil
}
func (mapi *memberApp) EditMember(mdo *aggregate.Member) (bool, error) {
	return true, nil
}
func (mapi *memberApp) Login(loginName string, password string) (bool, error) {
	return true, nil
}
func (mapp *memberApp) Get(uuid uuid.UUID) (*aggregate.Member, error) {
	model, err := mapp.repo.Get(uuid)
	if err != nil {
		return nil, err
	}
	return &model, nil
}
func (mapi *memberApp) ChangeMemberStatus(uuid uuid.UUID, status int) (bool, error) {
	return true, nil
}
func (mapi *memberApp) ChangeRoleRelationStatus(uuid uuid.UUID, status int) (bool, error) {
	return true, nil
}
