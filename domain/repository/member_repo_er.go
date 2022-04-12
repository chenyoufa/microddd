package repository

import (
	"context"
	"microddd/domain/aggregate"

	"github.com/google/uuid"
)

type MemberRepoer interface {
	Get(ctx context.Context, uuid uuid.UUID) (*aggregate.Member_aggre, error)
}
