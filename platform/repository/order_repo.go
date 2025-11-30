package repository

import (
	"database/sql"

	"github.com/davidpogosian/pickupee/platform/models"
)

type OrderRepository struct {
	db *sql.DB
}

// 【 Create 】
// POST request
func (r *OrderRepository) Insert(order *models.Order) (int, error) {
	result, err := r.db.Exec(
		"INSERT INTO orders (user_id) VALUES (?)",
		order.UserID,
	)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
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
