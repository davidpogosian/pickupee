package listOrdersForUser

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/davidpogosian/pickupee/service"
)

type responseStruct struct {
	OrderIDs []int `json:"order_ids"`
}

// Request:
// /orders?user_id=42
//
// JSON output:
//
//	{
//		"order_ids": [1,2,3]
//	}
func Handler(orderService *service.OrderService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Only allow GET
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Extract user_id from URL query string
		userIDStr := r.URL.Query().Get("user_id")
		if userIDStr == "" {
			http.Error(w, "Missing user_id", http.StatusBadRequest)
			return
		}

		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			http.Error(w, "Invalid user_id", http.StatusBadRequest)
			return
		}

		// Call service layer
		orders, err := orderService.ListOrdersForUser(userID)
		if err != nil {
			http.Error(w, "Failed to list orders: "+err.Error(), http.StatusInternalServerError)
			return
		}

		orderIDs := make([]int, len(orders))
		for i, order := range orders {
			orderIDs[i] = order.ID
		}

		resp := responseStruct{OrderIDs: orderIDs}

		// Output JSON
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
