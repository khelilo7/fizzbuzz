package db

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	log "github.com/sirupsen/logrus"

	"lbc/fizzbuzz/internal/utils"
)

func NewDBConn() *pg.DB {
	address := fmt.Sprintf("%s:%s", utils.Getenv("DB_HOST", "localhost"), utils.Getenv("DB_PORT", "5432"))
	options := &pg.Options{
		User:     utils.Getenv("DB_USER", "postgres"),
		Password: utils.Getenv("DB_PASS", "123456"),
		Addr:     address,
		Database: utils.Getenv("DB_NAME", "api_stats"),
		PoolSize: 50,
	}
	db := pg.Connect(options)
	if db == nil {
		log.Error("cannot connect to postgres")
	}
	return db
}
