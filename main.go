package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/davidpogosian/pickupee/platform/initialization"
	"github.com/davidpogosian/pickupee/platform/router"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Open database
	db, err := sql.Open("sqlite3", "./mydb.db")
	if err != nil {
		panic(err)
	}
	defer log.Println("Hello I am defer") // Ctrl+C does not allow main() to finish executing
	defer db.Close()                      // Might not be necessary?

	// Initialize database
	err = initialization.InitTables(db)
	if err != nil {
		panic(fmt.Errorf("Failed to initialize database: %w", err))
	}

	// Make router
	router := router.Create()

	// Launch server
	log.Println("Server starting at http://localhost:8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
