package repository

import (
	"database/sql"

	"github.com/davidpogosian/pickupee/platform/models"
)

type OrderRepository struct {
	db sql.DB
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
// 【 Update 】
// 【 Delete 】
