package main

import (
	"fmt"
	"net/http"

	"msg-app/profiler"
	API "msg-app/src/api"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	go profiler.MemoryProfiler()

	db, err := gorm.Open(sqlite.Open("database/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	api := API.NewAPI(db)
	router := api.GetRouter()

	fmt.Println("App listening at :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println(err)
	}
}
