package configs

type PostgresConfig struct {
	Host, Port, Username, Password, DBName string
	SSLMode                                string
}

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
