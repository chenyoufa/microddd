package application

import (
	"context"
	"microddd/domain/aggregate"
	"microddd/domain/repository"

	"github.com/google/uuid"
)

type MemberApper interface {
	Get(ctx context.Context, uuid uuid.UUID) (*aggregate.Member_aggre, error)
	GetList(ctx context.Context, uuid uuid.UUID) ([]*aggregate.Member_aggre, error)
	Add(ctx context.Context, aggre *aggregate.Member_aggre) (bool, error)
	Edit(ctx context.Context, aggre *aggregate.Member_aggre) (bool, error)
	Login(ctx context.Context, usname string, pwd string) (bool, error)
	Logout(ctx context.Context, usname string, pwd string) (bool, error)
}

type memberApp struct {
	mRepo repository.MemberRepoer
}

// var memberAppSet = wire.NewSet(
// 	wire.Struct(new(memberApp)),
// 	wire.Bind(new(MemberApper), new(*memberApp)))

func NewmemberApp(memberRepos repository.MemberRepoer) *memberApp {
	return &memberApp{
		mRepo: memberRepos,
	}
}

func (u *memberApp) Get(ctx context.Context, uuid uuid.UUID) (*aggregate.Member_aggre, error) {
	m, err := u.mRepo.Get(ctx, uuid)

	if err != nil {
		return nil, err
	}
	return m, nil
}

func (u *memberApp) GetList(ctx context.Context, uuid uuid.UUID) ([]*aggregate.Member_aggre, error) {
	list, err := u.mRepo.GetList(ctx, uuid)
	return list, err
}
func (u *memberApp) Add(ctx context.Context, aggre *aggregate.Member_aggre) (bool, error) {

	isbool, err := u.mRepo.Add(ctx, aggre)
	return isbool, err
}
func (u *memberApp) Edit(ctx context.Context, aggre *aggregate.Member_aggre) (bool, error) {

	isbool, err := u.mRepo.Edit(ctx, aggre)
	return isbool, err
}
func (u *memberApp) Login(ctx context.Context, usname string, pwd string) (bool, error) {
	isbool, err := u.mRepo.Login(ctx, usname, pwd)
	return isbool, err
}
func (u *memberApp) Logout(ctx context.Context, usname string, pwd string) (bool, error) {

	return true, nil
}
