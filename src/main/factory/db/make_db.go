package db

import (
	"log"

	"github.com/oderapi/src/main/factory/dsn"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MakePostgresDB() *gorm.DB {
	currentDsn := dsn.MakeDSN()
	db, err := gorm.Open(postgres.Open(currentDsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	return db
}
