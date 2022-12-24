package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDB(dbConf map[string]string) *gorm.DB {
	if dbConf["type"] == "sqlite" {
		db, err := gorm.Open(sqlite.Open(dbConf["url"]), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		return db
	}

	if dbConf["type"] == "postgresql" {
		db, err := gorm.Open(postgres.Open(dbConf["url"]), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		return db
	}

	panic("No database specified")
}
