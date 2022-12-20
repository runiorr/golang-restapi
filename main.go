package main

import (
	"fmt"
	"net/http"

	"msg-app/profiler"
	API "msg-app/src/api"
	"msg-app/src/db/factory"
)

func main() {
	go profiler.MemoryProfiler()

	database := factory.MemoryStorageFactory()
	api := API.NewAPI(database)

	fmt.Println("App listening at :8080")
	if err := http.ListenAndServe(":8080", api.Router); err != nil {
		fmt.Println(err)
	}
}
