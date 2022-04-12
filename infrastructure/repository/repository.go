package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepos *userRepos
}

func NewRepository(cdb *gorm.DB) *Repository {
	return &Repository{
		&userRepos{cdb},
	}
}
