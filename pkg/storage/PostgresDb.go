package storage

import (
	"fmt"
	"github.com/asavt7/nixEducation/pkg/configs"
	"github.com/jmoiron/sqlx"
	"log"
)

func NewPostgreDb(cfg configs.PostgresConfig) *sqlx.DB {

	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}
