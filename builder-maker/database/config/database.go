package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	GormDb *gorm.DB
	SqlDb  *sql.DB
}

func Open() (db DB, err error) {
	db.SqlDb, err = sql.Open("postgres", os.Getenv("DATABASE"))
	if err != nil {
		return
	}
	if db.GormDb, err = gorm.Open(postgres.New(postgres.Config{Conn: db.SqlDb}), &gorm.Config{}); err != nil {
		return
	}
	log.Printf("Connected To DB")
	return
}
