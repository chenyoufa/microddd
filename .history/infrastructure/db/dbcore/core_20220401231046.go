package dbcore

var (
	globalDB *gorm.DB

	globalConfig *DBConfig

	injectors []func(db *gorm.DB)
)

func Connect(cfg *DBConfig) {
	cfg = defaultDbConfig(cfg)
	var dialector gorm.Dialector

	switch strings.ToLower(c.DBType) {
	case "mysql":
		// create database if not exists
		cfg, err := mysqlDriver.ParseDSN(c.DSN)
		if err != nil {
			return nil, err
		}

		err = createDatabaseWithMySQL(cfg)
		if err != nil {
			return nil, err
		}

		dialector = mysql.Open(c.DSN)
	case "postgres":
		dialector = postgres.Open(c.DSN)
	default:
		dialector = sqlite.Open(c.DSN)
	}

	gconfig := &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.TablePrefix,
			SingularTable: true,
		},
	}

	db, err := gorm.Open(dialector, gconfig)
	if err != nil {
		return nil, err
	}

	if c.Debug {
		db = db.Debug()
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(c.MaxIdleConns)
	sqlDB.SetMaxOpenConns(c.MaxOpenConns)
	sqlDB.SetConnMaxLifetime(time.Duration(c.MaxLifetime) * time.Second)

}
