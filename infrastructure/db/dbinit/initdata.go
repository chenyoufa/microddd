package dbinit

import (
	"log"
	"microddd/infrastructure/db/dbcore"

	"github.com/pkg/errors"
)

func InitData() error {
	locker := dbcore.NewLockDb("init", dbcore.GetHostname(), dbcore.DefaultLeaseAge)
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

func initConfig() {

}
