package repository

import (
	"context"
	aggregate "gomicroddd/domain/aggregate"
	"gomicroddd/infrastructure/db/dbcore"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func init() {
	// t, _ := aggregate.NewCustomer("aa")
	// fmt.Print(t)
	dbcore.RegisterInjector(func(db *gorm.DB) {
		// dbcore.SetupTableModel(db, &t)

	})
}

type memoryProductRepository struct {
	// db *gorm.DB
}

func MemoryProductRepository() *memoryProductRepository {
	return &memoryProductRepository{}
}

// New is a factory function to generate a new repository of customers
// func NewProductRepo(db *gorm.DB) *MemoryProductRepository {
// 	return &MemoryProductRepository{
// 		db: db,
// 	}
// }

func (mpr *memoryProductRepository) GetAll(ctx context.Context) ([]*aggregate.Product, error) {
	var r []*aggregate.Product
	err := dbcore.GetDB(ctx).Find(&r).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return r, nil
}
