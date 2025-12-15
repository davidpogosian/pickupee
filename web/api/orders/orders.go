package orders

import (
	"net/http"

	"github.com/davidpogosian/pickupee/service"
)

type OrdersHandler struct {
	Service *service.OrderService
}

func (h *OrdersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.listOrdersForUserHandler(w, r)
	case http.MethodPost:
		h.placeOrderHandler(w, r)
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}
