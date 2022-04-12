package repository

import (
	"context"
	"microddd/domain/aggregate"
	"microddd/infrastructure/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepos struct {
	db *gorm.DB
}

func (u *userRepos) Get(ctx context.Context, uuid uuid.UUID) (*aggregate.Member_aggre, error) {
	userpo := model.User_po{}
	err := u.db.Where("id=?", uuid).Find(&userpo).Error
	userdo := &aggregate.Member_aggre{}
	return userdo, err
}
