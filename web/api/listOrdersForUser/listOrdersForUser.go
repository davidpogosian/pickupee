package listOrdersForUser

import (
	"encoding/json"
	"net/http"

	"github.com/davidpogosian/pickupee/service"
)

type requestStruct struct {
	UserID int `json:"user_id"`
}

type responseStruct struct {
	OrderIDs []int `json:"order_ids"`
}

// JSON input:
//
//	{
//		"user_id": 42
//	}
//
// JSON output:
//
//	{
//		"order_ids": [1,2,3]
//	}
func Handler(orderService *service.OrderService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Only allow POST
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		var req requestStruct
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
			return
		}

		orders, err := orderService.ListOrdersForUser(req.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var orderIDs []int
		for _, order := range orders {
			orderIDs = append(orderIDs, order.ID)
		}

		resp := responseStruct{OrderIDs: orderIDs}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
}
