package orders

import (
	"encoding/json"
	"net/http"
)

type placeOrderRequest struct {
	UserID  int   `json:"user_id"`
	ItemIDs []int `json:"item_ids"`
}

type placeOrderResponse struct {
	OrderID int `json:"order_id"`
}

// JSON input:
//
//	{
//		"user_id": 42,
//		"item_ids": [1,2,3]
//	}
//
// JSON output:
//
//	{
//		"order_id": 2
//	}
func (h *OrdersHandler) placeOrderHandler(w http.ResponseWriter, r *http.Request) {
	var req placeOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	orderID, err := h.Service.PlaceOrder(req.UserID, req.ItemIDs)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := placeOrderResponse{OrderID: orderID}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
