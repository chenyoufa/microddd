package application

import (
	"context"
	"microddd/domain/aggregate"
	"microddd/domain/repository"

	"github.com/google/uuid"
)

type memberApper interface {
	Get(ctx context.Context, uuid uuid.UUID) (*aggregate.Member_aggre, error)
}

type App struct {
	Mapp memberApper
}

func NewApps(repo *repository.AuthFactory) *App {
	return &App{
		&memberApp{mRepo: repo.MRepo},
	}
}
