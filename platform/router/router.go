package router

import (
	"net/http"

	"github.com/davidpogosian/pickupee/service"
	"github.com/davidpogosian/pickupee/web/api/listOrdersForUser"
	"github.com/davidpogosian/pickupee/web/api/placeOrder"
)

func Create(orderService *service.OrderService) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/placeOrder", placeOrder.Handler(orderService))
	mux.HandleFunc("/listOrdersForUser", listOrdersForUser.Handler(orderService)) // /orders?user_id=42

	return mux
}
