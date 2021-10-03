package storage

import (
	"database/sql"
	"fmt"
)

type Config struct {
	Host, Port, Username, Password, DBName string
	SSLMode                                string
}

func NewPostgreDb(cfg Config) (*sql.DB, error) {

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
