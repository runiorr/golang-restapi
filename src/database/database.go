package database

import (
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDB(dbConf map[string]string) *gorm.DB {
	if dbConf["type"] == "sqlite" {
		db, err := gorm.Open(sqlite.Open(dbConf["url"]), &gorm.Config{})
		if err != nil {
			panic("failed to connect to sqlite")
		}
		return db
	}

	if dbConf["type"] == "postgresql" {
		conString := os.Getenv("CONSTRING")
		db, err := gorm.Open(postgres.Open(conString), &gorm.Config{})
		if err != nil {
			panic("failed to connect to postgres")
		}
		return db
	}

	panic("no database specified")
}
