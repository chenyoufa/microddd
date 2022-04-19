package repository

import (
	"context"
	"microddd/domain/aggregate"

	"github.com/google/uuid"
)

type MemberRepoer interface {
	Get(ctx context.Context, uuid uuid.UUID) (*aggregate.Member_aggre, error)
	GetList(ctx context.Context, uuid uuid.UUID) ([]*aggregate.Member_aggre, error)
	Add(ctx context.Context, aggre *aggregate.Member_aggre) (bool, error)
	Edit(ctx context.Context, aggre *aggregate.Member_aggre) (bool, error)
	Login(ctx context.Context, usname string, pwd string) (bool, error)
}
