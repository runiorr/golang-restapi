package main

import (
	API "msg-app/src"

	c "msg-app/src/config"
	um "msg-app/src/core/users/model"
	"msg-app/src/database"
)

func main() {
	var config c.Conf
	config.GetConfig()

	db := database.GetDB(&config)
	db.AutoMigrate(&um.User{})

	api := API.NewAPI(db, config.Http)
	api.SetupRouter()
	api.Start()
}
