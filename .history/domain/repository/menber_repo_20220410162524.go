package repository

import (
	"context"
	"gomicroddd/domain/aggregate"

	"github.com/google/uuid"
)

// CustomerRepository is a interface that defines the rules around what a customer repository
// Has to be able to perform
type MemberRepository interface {
	Get(ctx context.Context, uuid uuid.UUID) (aggregate.Member, error)
	Add(nm aggregate.Member) error
	Update(nm aggregate.Member) error
}
