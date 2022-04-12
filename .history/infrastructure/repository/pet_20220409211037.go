package repository

import (
	"gorm.io/gorm"

	petmodel "gomicroddd/domain/entity/pet"
	"gomicroddd/infrastructure/db/dbcore"

	"github.com/pkg/errors"
)

func init() {
	dbcore.RegisterInjector(func(db *gorm.DB) {
		// dbcore.SetupTableModel(db, &petmodel.Pet{})
	})
}

type petDb struct {
	db *gorm.DB
}

func (s *petDb) List(query *petmodel.Pet, offset, limit int) ([]*petmodel.Pet, error) {
	var r []*petmodel.Pet

	db := dbcore.WithOffsetLimit(s.db, offset, limit)

	err := db.Where(query).Find(&r).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return r, nil
}

func (s *petDb) Get(id string) (*petmodel.Pet, error) {
	var r petmodel.Pet
	err := s.db.Where("id = ?", id).First(&r).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &r, nil
}

func (s *petDb) Create(in *petmodel.Pet) (*petmodel.Pet, error) {
	err := s.db.Create(in).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return in, nil
}

func (s *petDb) Update(in *petmodel.Pet) (*petmodel.Pet, error) {
	err := s.db.Updates(in).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return in, nil
}

func (s *petDb) Delete(in *petmodel.Pet) error {
	err := s.db.Where(in).Delete(&petmodel.Pet{}).Error
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
