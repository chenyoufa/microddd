package repository

import (
	"gomicroddd/domain/aggregate"

	"github.com/google/uuid"
)

// CustomerRepository is a interface that defines the rules around what a customer repository
// Has to be able to perform
type MemberRepository interface {
	Get(uuid.UUID) (aggregate.Member, error)
	Add(aggregate.Member) error
	Update(aggregate.Member) error
}
