package repository

import (
	"context"
	"fmt"
	"gomicroddd/domain/aggregate"
	"gomicroddd/domain/entity"
	"gomicroddd/domain/valueobject"
	"gomicroddd/infrastructure/db/dbcore"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Po struct {
	entity.User
	roles []entity.Role
}
type User struct {
	gorm.Model
	Refer string // 关联外键
	Name  string
}

type Profile struct {
	gorm.Model
	Name      string
	User      User `gorm:"references:Refer"` // 使用 Refer 作为关联外键
	UserRefer string
}

func init() {

	dbcore.RegisterInjector(func(db *gorm.DB) {
		dbcore.SetupTableModel(db, &User{})
		dbcore.SetupTableModel(db, &Profile{})
		dbcore.SetupTableModel(db, &valueobject.UserRoleRelation{})
	})
}

type MemberDomain struct{}

// func init() {
// 	mapper.Register(&.MemberDto{})
// 	mapper.Register(&entity.User{})
// }

func (m *MemberDomain) Get(ctx context.Context, uuid uuid.UUID) (*aggregate.Member, error) {
	var r *aggregate.Member
	var user User
	profile := Profile{}

	var roles []entity.Role
	var relations []valueobject.UserRoleRelation

	// db := dbcore.GetDB(ctx).Debug().Where("user_id=?", uuid).Model(&user).Association("id").Find(&roles)
	// db := dbcore.GetDB(ctx).Debug().Model(&relations).Association("Roleid").Find(&roles)

	db := dbcore.GetDB(ctx).Model(&profile).Association("User").Find(&user)
	// db = db.Preload("user_id").Find(&user)
	// db.Preload("role").Find(&roles)
	//
	if db != nil {

	}
	fmt.Println(user)
	fmt.Println(relations)
	fmt.Println(roles)
	// err = dbcore.GetDB(ctx).Find(&roles).Error
	// err = dbcore.GetDB(ctx).Find(&relations).Error
	// err = setUnExportedStrField(r, "roles", roles)
	// err = setUnExportedStrField(r, "userRoleRelation", relations)
	var err error
	if err != nil {
		return nil, nil
	}
	return r, nil
}
func (m *MemberDomain) Add(nm aggregate.Member) error {
	return nil
}
func (m *MemberDomain) Update(aggregate.Member) error {
	return nil
}
