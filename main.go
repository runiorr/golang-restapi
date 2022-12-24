package main

import (
	"fmt"
	"net/http"

	API "msg-app/src"

	"msg-app/src/config"
	um "msg-app/src/core/users/model"
	"msg-app/src/database"
)

func main() {
	var conf config.Conf
	conf.GetConf()

	fmt.Println(conf.Http)
	fmt.Println(conf.Database)

	db := database.GetDB(conf.Database)

	db.AutoMigrate(&um.User{})

	api := API.NewAPI(db)
	router := api.GetRouter()

	port := fmt.Sprintf(":%s", conf.Http["port"])
	fmt.Printf("App listening at port %s using %s\n", port, conf.Database["type"])
	if err := http.ListenAndServe(port, router); err != nil {
		fmt.Println(err)
	}
}
