package main

import (
	API "msg-app/src"

	"msg-app/src/config"
	um "msg-app/src/core/users/model"
	"msg-app/src/database"
)

func main() {
	var conf config.Conf
	conf.GetConf()

	db := database.GetDB(conf.Database)
	db.AutoMigrate(&um.User{})

	api := API.NewAPI(db)
	api.Start()
}
