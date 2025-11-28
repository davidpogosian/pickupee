package initialization

import (
	"database/sql"
)

func InitTables(db *sql.DB) error {
	ordersTable := `
    CREATE TABLE IF NOT EXISTS orders (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id INTEGER NOT NULL,
        placed_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
    );`

	orderItemsTable := `
    CREATE TABLE IF NOT EXISTS order_items (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        order_id INTEGER NOT NULL,
        item_id INTEGER NOT NULL,
        FOREIGN KEY(order_id) REFERENCES orders(id)
    );`

	if _, err := db.Exec(ordersTable); err != nil {
		return err
	}

	if _, err := db.Exec(orderItemsTable); err != nil {
		return err
	}

	return nil
}
