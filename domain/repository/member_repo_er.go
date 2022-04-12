package repository

import (
	"context"
	"microddd/domain/aggregate"

	"github.com/google/uuid"
)

type Member_repo_er interface {
	Get(ctx context.Context, uuid uuid.UUID) (*aggregate.Member_aggre, error)
}
