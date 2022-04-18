package dbcore

import (
	"log"
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

// tag按首字母排序
type CommonModel struct {
	Id        string    `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
}

// https://github.com/ulid/spec
// uuid sortable by time
func NewUlid() string {
	now := time.Now()
	return ulid.MustNew(ulid.Timestamp(now), ulid.Monotonic(rand.New(rand.NewSource(now.UnixNano())), 0)).String()
}

func registerCallback(db *gorm.DB) {
	// 自动添加uuid
	err := db.Callback().Create().Before("gorm:create").Register("uuid", func(db *gorm.DB) {
		// if _, ok := uuid.Parse(db.Statement.Get("ID")); ok {

		// }
		log.Println(db.Statement.Get("id"))
		// db.Statement.SetColumn("ID", NewUlid())
		db.Statement.SetColumn("ID", uuid.New())
	})
	if err != nil {
		log.Panicf("err: %+v", errors.WithStack(err))
	}
}
func WithOffsetLimit(db *gorm.DB, offset, limit int) *gorm.DB {
	if offset > 0 {
		db = db.Offset(offset)
	}
	if limit > 0 {
		db = db.Limit(limit)
	}
	return db
}
