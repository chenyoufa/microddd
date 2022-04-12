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

// func NewProductService(mRepo repository.MemberRepo) *memberApp {
// 	return &memberApp{
// 		mRepo: mRepo,
// 	}
// }
func (u *memberApp) Get(ctx context.Context, uuid uuid.UUID) (*aggregate.Member_aggre, error) {
	m, err := u.mRepo.Get(ctx, uuid)

	if err != nil {
		return nil, err
	}
	return m, nil
}
