package application

import (
	"context"
	"log"
	"microddd/application/dto"
	"microddd/domain/repository"

	"github.com/google/uuid"
)

type MemberApper interface {
	Get(ctx context.Context, uid uuid.UUID) (*dto.Member_dto, error)
	GetList(ctx context.Context, uid uuid.UUID) ([]*dto.Member_dto, error)
	Add(ctx context.Context, mdto *dto.Member_dto) (bool, error)
	Edit(ctx context.Context, mdto *dto.Member_dto) (bool, error)
	Login(ctx context.Context, usname string, pwd string) (bool, error)
	Logout(ctx context.Context, usname string, pwd string) (bool, error)
	Remove(ctx context.Context, uid uuid.UUID) (bool, error)
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

func (u *memberApp) Get(ctx context.Context, uid uuid.UUID) (*dto.Member_dto, error) {
	m, err := u.mRepo.Get(ctx, uid)
	// log.Panicln("m:", m)
	rdto := dto.Member_dto{}
	rdto.ToDto(m)
	if err != nil {
		return nil, err
	}
	return &rdto, nil
}

func (u *memberApp) GetList(ctx context.Context, uid uuid.UUID) ([]*dto.Member_dto, error) {
	list, err := u.mRepo.GetList(ctx, uid)
	var dtolist []*dto.Member_dto
	for _, item := range list {
		rdto := dto.Member_dto{}
		rdto.ToDto(item)
		dtolist = append(dtolist, &rdto)
	}
	if err != nil {
		return nil, err
	}
	return dtolist, err
}
func (u *memberApp) Add(ctx context.Context, mdto *dto.Member_dto) (bool, error) {

	aggre := mdto.ToAggre()
	// log.Println("aggre:", aggre)
	isbool, err := u.mRepo.Add(ctx, aggre)
	return isbool, err
}
func (u *memberApp) Edit(ctx context.Context, mdto *dto.Member_dto) (bool, error) {
	aggre := mdto.ToAggre()
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
func (u *memberApp) Remove(ctx context.Context, uid uuid.UUID) (bool, error) {

	model, err := u.mRepo.Get(ctx, uid)
	log.Println("model:", model)
	model.User.Status = -1
	if err != nil {
		return false, err
	}
	ok, err := u.mRepo.Remove(ctx, model)
	if err != nil || !ok {
		return false, err
	}
	return true, nil
}
