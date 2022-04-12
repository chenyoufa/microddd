package application

import (
	"context"
	"microddd/domain/aggregate"
	"microddd/domain/repository"

	"github.com/google/uuid"
)

type memberApper interface {
	Get(ctx context.Context, uuid uuid.UUID) (*aggregate.Member_aggre, error)
}
type memberApp struct {
	mRepo repository.MemberRepoer
}

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
