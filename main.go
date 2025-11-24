package main

import (
	"database/sql"
	"net/http"

	"github.com/davidpogosian/pickupee/platform/initialization"
	"github.com/davidpogosian/pickupee/platform/logger"
	"github.com/davidpogosian/pickupee/platform/router"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open database
	db, err := sql.Open("sqlite3", "./mydb.db")
	if err != nil {
		logger.Fatalf(err.Error())
	}
	defer logger.Info("Hello I am defer") // Ctrl+C does not allow main() to finish executing
	defer db.Close()                      // Might not be necessary?

	// Initialize database
	initialization.InitTables(db)

	// Make router
	router := router.Create()

	// Launch server
	logger.Info("Server starting at http://localhost:8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		logger.Error("Error starting server:", err)
	}
}
