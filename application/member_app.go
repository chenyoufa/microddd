package application

import (
	"context"
	"microddd/domain/aggregate"
	"microddd/domain/repository"

	"github.com/google/uuid"
)

type memberApp struct {
	mRepo repository.MemberRepo
}

func NewmemberApp(memberRepos repository.MemberRepo) *memberApp {
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
