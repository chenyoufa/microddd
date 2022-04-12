package aggregate

import (
	"errors"
	"gomicroddd/domain/entity"
	"gomicroddd/domain/valueobject"
)

var (
	// ErrInvalidPerson is returned when the person is not valid in the NewCustome factory
	ErrInvalidResources = errors.New("a menber has to have an valid person")
)

type Resources struct {
	Role             *entity.Role
	menu             *entity.Menu
	roleMenuRelation []valueobject.RoleMenuRelation
}
