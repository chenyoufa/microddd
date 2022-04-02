package dbcore

type Config struct {
	Debug        bool
	DbType       string
	DSN          string
	MaxLifetime  int
	MaxOpenConns int
	MaxIdleConns int
	TablePrefix  string
	AutoMigrate  bool
}
