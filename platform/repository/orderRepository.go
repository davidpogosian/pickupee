package repository

import (
	"database/sql"

	"github.com/davidpogosian/pickupee/platform/models"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

// 【 Create 】
// POST request
func (r *OrderRepository) CreateOrder(userID int, itemIDs []int) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return -1, err
	}

	// ensure rollback on any error
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	// 1. Insert the order
	var res sql.Result
	res, err = tx.Exec("INSERT INTO orders (user_id) VALUES (?)", userID)
	if err != nil {
		return -1, err
	}

	var orderID64 int64
	orderID64, err = res.LastInsertId()
	if err != nil {
		return -1, err
	}
	orderID := int(orderID64)

	// 2. Prepare insert for items
	var stmt *sql.Stmt
	stmt, err = tx.Prepare("INSERT INTO order_items (order_id, item_id) VALUES (?, ?)")
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	for _, itemID := range itemIDs {
		_, err = stmt.Exec(orderID, itemID)
		if err != nil {
			return -1, err
		}
	}

	// 3. Commit
	err = tx.Commit()
	if err != nil {
		return -1, err
	}

	return orderID, nil
}

// 【 Read 】
// GET request
func (r *OrderRepository) ListByUserID(userID int) ([]models.Order, error) {
	rows, err := r.db.Query(
		"SELECT id, user_id, placed_at FROM orders WHERE user_id = ?",
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []models.Order{}

	for rows.Next() {
		var o models.Order
		err := rows.Scan(&o.ID, &o.UserID, &o.PlacedAt)
		if err != nil {
			return nil, err
		}
		orders = append(orders, o)
	}

	// Check for scan/iteration errors
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return orders, nil
}

// 【 Update 】
// 【 Delete 】
