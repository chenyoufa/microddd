package dbcore

type DBConfig struct {
	Debug        bool
	DbType       string
	DSN          string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
	TablePrefix  string
	AutoMigrate  bool
}

func defaultDbConfig(cfg *DBConfig) *DBConfig {
	newCfg := *cfg
	if newCfg.MaxIdleConns == 0 {
		newCfg.MaxIdleConns = 10
	}

	if newCfg.MaxOpenConns == 0 {
		newCfg.MaxOpenConns = 20
	}
	return &newCfg
}
