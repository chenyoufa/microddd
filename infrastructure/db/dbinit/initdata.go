package dbinit

import (
	"log"
	"microddd/infrastructure/db/dbcore"

	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func InitData(defDB *gorm.DB) error {
	locker := dbcore.NewLockDb("init", dbcore.GetHostname(), dbcore.DefaultLeaseAge, defDB)
	ok, err := locker.Lock()
	if err != nil {
		return errors.WithStack(err)
	}

	if !ok {
		return nil
	}

	defer func() {
		_ = locker.UnLock()
	}()

	return run()
}

func run() error {
	log.Printf("%s begin init data", dbcore.GetHostname())
	return nil
}

func LoadConfig() (*dbcore.DBConfig, error) {

	config := dbcore.DBConfig{
		DbType:      "mysql",
		DSN:         "fage:Fage501526~@(127.0.0.1:3306)/mytest",
		AutoMigrate: true,
		Debug:       true,
	}
	return &config, nil
}
