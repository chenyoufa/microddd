package repository

import (
	"github.com/pkg/errors"
	"gorm.io/gorm"

	petmodel "gomicroddd/domain/entity/pet"
	"gomicroddd/infrastructure/db/dbcore"
)

func init() {
	dbcore.RegisterInjector(func(db *gorm.DB) {
		dbcore.SetupTableModel(db, &petmodel.OwnerPet{})
	})
}

type ownerPetDb struct {
	db *gorm.DB
}

func (s *ownerPetDb) Query(in *petmodel.OwnerPet) ([]*petmodel.OwnerPet, error) {
	var r []*petmodel.OwnerPet
	err := s.db.Where(in).Find(&r).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return r, nil
}

func (s *ownerPetDb) Create(in *petmodel.OwnerPet) (*petmodel.OwnerPet, error) {
	err := s.db.Create(in).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return in, nil
}

func (s *ownerPetDb) Delete(in *petmodel.OwnerPet) error {
	err := s.db.Where(in).Delete(&petmodel.OwnerPet{}).Error
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
