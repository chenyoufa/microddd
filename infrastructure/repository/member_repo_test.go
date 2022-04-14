package repository

import (
	"context"
	"microddd/application"
	"microddd/infrastructure/db/dbcore"
	"microddd/infrastructure/db/dbinit"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMenu(t *testing.T) {

	var abp application.MemberApper
	config, _ := dbinit.LoadConfig()
	db, _ := dbcore.Connect(config)
	factory := NewRepository(db)
	abp = application.NewApps(factory).Mapp

	model, err := abp.Get(context.Background(), uuid.New())
	assert.Nil(t, err)
	assert.Nil(t, model)

}
