package storage

import (
	"database/sql"
	"fmt"
	"github.com/asavt7/nixEducation/pkg/configs"
	"log"
)

func NewPostgreDb(cfg configs.PostgresConfig) *sql.DB {

	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}
