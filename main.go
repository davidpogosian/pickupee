package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/davidpogosian/pickupee/platform/initialization"
	"github.com/davidpogosian/pickupee/platform/repository"
	"github.com/davidpogosian/pickupee/platform/router"
	"github.com/davidpogosian/pickupee/service"
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

	// Instantiate Repositories
	orderRepository := repository.NewOrderRepository(db)

	// Instantiate Services
	orderService := service.NewOrderService(orderRepository)

	// Make r
	r := router.Create(orderService)

	// Launch server
	log.Println("Server starting at http://localhost:8080")
	err = http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
