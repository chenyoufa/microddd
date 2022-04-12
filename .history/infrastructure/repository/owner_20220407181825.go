package repository

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	petmodel "gomicroddd/domain/entity/pet"
	"gomicroddd/infrastructure/db/dbcore"
)

func init() {
	dbcore.RegisterInjector(func(db *gorm.DB) {
		dbcore.SetupTableModel(db, &petmodel.Owner{})
	})
}

type ownerDb struct {
	db *gorm.DB
}

func (s *ownerDb) List(query *petmodel.Owner, offset, limit int) ([]*petmodel.Owner, error) {
	var r []*petmodel.Owner

	db := dbcore.WithOffsetLimit(s.db, offset, limit)

	err := db.Where(query).Find(&r).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return r, nil
}

func (s *ownerDb) Get(id string) (*petmodel.Owner, error) {
	var r petmodel.Owner
	err := s.db.Where("id = ?", id).First(&r).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return &r, nil
}

func (s *ownerDb) Create(in *petmodel.Owner) (*petmodel.Owner, error) {
	err := s.db.Create(in).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return in, nil
}

func (s *ownerDb) Update(in *petmodel.Owner) (*petmodel.Owner, error) {
	err := s.db.Updates(in).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return in, nil
}

func (s *ownerDb) Delete(in *petmodel.Owner) error {
	err := s.db.Where(in).Delete(&petmodel.Owner{}).Error
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
