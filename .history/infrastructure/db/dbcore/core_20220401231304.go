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
