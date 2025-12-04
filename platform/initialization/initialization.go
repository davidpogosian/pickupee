package initialization

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/davidpogosian/pickupee/platform/repository"
	"github.com/davidpogosian/pickupee/platform/router"
	"github.com/davidpogosian/pickupee/service"
)

func initTables(db *sql.DB) error {
	ordersTable := `
    CREATE TABLE IF NOT EXISTS orders (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER NOT NULL,
        placed_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
    );`

	itemsTable := `
    CREATE TABLE IF NOT EXISTS items (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT NOT NULL
    );`

	orderItemsTable := `
    CREATE TABLE IF NOT EXISTS order_items (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        order_id INTEGER NOT NULL,
        item_id INTEGER NOT NULL,
        FOREIGN KEY(order_id) REFERENCES orders(id),
        FOREIGN KEY(item_id) REFERENCES items(id)
    );`

	if _, err := db.Exec(ordersTable); err != nil {
		return err
	}

	if _, err := db.Exec(itemsTable); err != nil {
		return err
	}

	if _, err := db.Exec(orderItemsTable); err != nil {
		return err
	}

	return nil
}

// Note: database is returned purely for the purpose of closing it
func CreateServer(dbLocation string) (*sql.DB, *http.ServeMux) {
	// Open database
	db, err := sql.Open("sqlite3", dbLocation)
	if err != nil {
		panic(err)
	}

	// Initialize database
	err = initTables(db)
	if err != nil {
		panic(fmt.Errorf("Failed to initialize database: %w", err))
	}

	// Instantiate Repositories
	orderRepository := repository.NewOrderRepository(db)

	// Instantiate Services
	orderService := service.NewOrderService(orderRepository)

	// Make r
	return db, router.Create(orderService)
}
