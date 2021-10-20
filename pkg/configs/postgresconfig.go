package configs

// PostgresConfig - configs for connection to postgres db
type PostgresConfig struct {
	Host, Port, Username, Password, DBName string
	SSLMode                                string
}

// InitPostgresConfig read configs from envs\config files and returns PostgresConfig
func InitPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "postgres",
		DBName:   "postgres",
		SSLMode:  "disable",
	}
}
