package dbcore

import (
	"context"

	"gorm.io/gorm"
)

type ctxTransactionKey struct{}

func FromTransaction(ctx context.Context) (interface{}, bool) {
	v := ctx.Value(ctxTransactionKey{})
	return v, v != nil
}

func CtxWithTransaction(ctx context.Context, tx *gorm.DB) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, ctxTransactionKey{}, tx)
}

func Transaction(ctx context.Context, db *gorm.DB, fn func(txctx context.Context) error) error {
	if _, ok := FromTransaction(ctx); ok {
		return fn(ctx)
	}
	return db.Transaction(func(tx *gorm.DB) error {
		txctx := CtxWithTransaction(ctx, tx)
		return fn(txctx)
	})
}
