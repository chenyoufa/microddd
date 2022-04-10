package aggregate

import (
	"errors"
	"gomicroddd/domain/entity"
	"gomicroddd/domain/valueobject"
	"gomicroddd/infrastructure/utils/hash"

	"github.com/google/uuid"
)

var (
	ErrInvalidMember = errors.New("a menber has to have an valid menber")
)

type Member struct {
	User             *entity.User
	roles            []*entity.Role
	userRoleRelation []valueobject.UserRoleRelation
}

// func NewMember(name string, password string, email string) (Member, error) {
// 	// Validate that the Name is not empty
// 	if name == "" || password == "" || email == "" {
// 		return Member{}, ErrInvalidMember
// 	}

// 	// Create a new person and generate ID
// 	user := &entity.User{
// 		UserName: name,
// 		Email:    email,
// 		Password: hash.MD5String(password),
// 		ID:       uuid.New(),
// 	}

	// Create a customer object and initialize all the values to avoid nil pointer exceptions
	return Member{
		User: user,
	}, nil
}

func (m *Member) AddRoles(r entity.Role) {
	m.roles = append(m.roles, &r)
	m.userRoleRelation = append(m.userRoleRelation, valueobject.UserRoleRelation{UserID: m.User.ID, RoleID: r.ID})
}

func (m *Member) RemoveRoles(r entity.Role) {
	var newroles []*entity.Role
	var rolation []valueobject.UserRoleRelation
	for _, item := range m.roles {
		tempid := item.ID
		if tempid == r.ID {
			continue
		}

		rolation = append(rolation, valueobject.UserRoleRelation{UserID: m.User.ID, RoleID: item.ID})
		newroles = append(newroles, item)
	}
	m.userRoleRelation = rolation
	m.roles = newroles
}

func (m *Member) GetRoles() []*entity.Role {
	return m.roles
}
