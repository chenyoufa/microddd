package dbcore

type Config struct {
	Debug        bool
	DbType       string
	DSN          string
	MaxLifetime  int
	MaxIdleConns int
}
