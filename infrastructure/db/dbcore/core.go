package dbcore

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	mysqlDriver "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	// "github.com/jinzhu/gorm"
)

var (
	globalDB     *gorm.DB
	globalConfig *DBConfig
	injectors    []func(db *gorm.DB)
)

func Connect(cfg *DBConfig) (*gorm.DB, error) {
	cfg = defaultDbConfig(cfg)
	var dialector gorm.Dialector

	switch strings.ToLower(cfg.DbType) {
	case "mysql":
		// create database if not exists
		parseCfg, err := mysqlDriver.ParseDSN(cfg.DSN)
		if err != nil {
			return nil, err
		}
		err = createDatabaseWithMySQL(parseCfg)
		if err != nil {
			return nil, err
		}
		dialector = mysql.Open(cfg.DSN)
	case "postgres":
		dialector = postgres.Open(cfg.DSN)
	default:
		dialector = sqlite.Open(cfg.DSN)
	}

	gconfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   cfg.TablePrefix,
			SingularTable: true,
		},
	}

	db, err := gorm.Open(dialector, gconfig)
	if err != nil {
		return nil, err
	}

	if cfg.Debug {
		db = db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(cfg.MaxLifetime) * time.Second)

	globalDB = db
	globalConfig = cfg
	registerCallback(db)
	callInjector(db)

	return db, nil
}
func createDatabaseWithMySQL(cfg *mysqlDriver.Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/", cfg.User, cfg.Passwd, cfg.Addr)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("sql open err:%v", err)
		return err
	}
	defer db.Close()

	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET = `utf8mb4`;", cfg.DBName)
	_, err = db.Exec(query)

	return err
}
func RegisterInjector(f func(*gorm.DB)) {
	injectors = append(injectors, f)
}
func callInjector(db *gorm.DB) {
	for _, v := range injectors {
		v(db)
	}
}

// 如果使用跨模型事务则传参
func GetDB(ctx context.Context) *gorm.DB {
	iface := ctx.Value(ctxTransactionKey{})

	if iface != nil {
		tx, ok := iface.(*gorm.DB)
		if !ok {
			log.Panicf("unexpect context value type: %s", reflect.TypeOf(tx))
			return nil
		}

		return tx
	}

	return globalDB.WithContext(ctx)
}

func GetDBConfig() DBConfig {
	return *globalConfig
}

// 自动初始化表结构
func SetupTableModel(db *gorm.DB, model interface{}) {
	if GetDBConfig().AutoMigrate {
		err := db.AutoMigrate(model)
		if err != nil {
			log.Fatal(err)
		}
	}
}
