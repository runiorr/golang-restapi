package database

import (
	"fmt"

	c "msg-app/src/config"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDB(config *c.Conf) *gorm.DB {
	if config.Env["name"] == "local" {
		db, err := gorm.Open(sqlite.Open(config.Database["path"]), &gorm.Config{})
		if err != nil {
			fmt.Println(err)
			panic("failed to connect to sqlite")
		}
		return db
	}

	if config.Env["name"] == "docker" {
		conString := fmt.Sprintf("postgresql://%s:%s@postgres:%s/%s?sslmode=disable",
			config.Database["user"],
			config.Database["pass"],
			config.Database["port"],
			config.Database["name"])

		db, err := gorm.Open(postgres.Open(conString), &gorm.Config{})
		if err != nil {
			fmt.Println(err)
			panic("failed to connect to postgres")
		}
		return db
	}

	panic("no database specified")
}
