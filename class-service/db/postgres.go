package db

import (
	"fmt"
	"github.com/Ddarli/gym/common"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresConnection(conf *common.Config) (*sqlx.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.DbHost, conf.DbPort, conf.DbUser, conf.DbPassword, conf.DbName)

	return sqlx.Connect("postgres", psqlInfo)
}
