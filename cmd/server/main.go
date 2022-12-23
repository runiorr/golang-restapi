package main

import (
	"fmt"
	// "io/ioutil"
	"net/http"

	API "msg-app/internal"
	"msg-app/profiler"

	um "msg-app/internal/core/users/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	go profiler.MemoryProfiler()

	db, err := gorm.Open(sqlite.Open("database/test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&um.User{})

	api := API.NewAPI(db)
	router := api.GetRouter()

	fmt.Println("App listening at :8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Println(err)
	}
}

// // Load configuration from config/config.yaml which contains details such as DB connection params
// cfg, err := config.Load(ctx)
// if err != nil {
// 	return nil, err
// }

// // Connect to the postgres DB
// db, err := initDatabase(ctx, cfg, a)
// if err != nil {
// 	return nil, err
// }

// // Run our migrations which will update the DB or create it if it doesn't exist
// if err := db.MigratePostgres(ctx, "file://migrations"); err != nil {
// 	return nil, err
// }
// a.OnShutdown(func() {
// 	// Temp for development so database is cleared on shutdown
// 	if err := db.RevertMigrations(ctx, "file://migrations"); err != nil {
// 		logging.From(ctx).Error("failed to revert migrations", zap.Error(err))
// 	}
// })

// // Instantiate and connect all our classes
// us := store.New(db.GetDB())
// e := events.New()
// u := users.New(us, e)

// httpServer := httptransport.New(u, db.GetDB())

// // Create a HTTP server
// h, err := http.New(httpServer, cfg.HTTP)
// if err != nil {
// 	return nil, err
// }

// // Start listening for HTTP requests
// return []app.Listener{
// 	h,
// }, nil
