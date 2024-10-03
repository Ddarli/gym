package db

import (
	"fmt"
	"github.com/Ddarli/gym/userservice/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresConnection(conf *config.Config) (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.DbHost, conf.DbPort, conf.DbUser, conf.DbPassword, conf.DbName)

	return sqlx.Connect("postgres", psqlInfo)
}
