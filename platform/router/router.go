package router

import (
	"net/http"

	"github.com/davidpogosian/pickupee/service"
	"github.com/davidpogosian/pickupee/web/api/listOrdersForUser"
	"github.com/davidpogosian/pickupee/web/api/placeOrder"
	"github.com/davidpogosian/pickupee/web/socket"
)

func Create(orderService *service.OrderService) *http.ServeMux {
	mux := http.NewServeMux()

	// TODO: Make a separate /orders mux
	mux.HandleFunc("/placeOrder", placeOrder.Handler(orderService))
	mux.HandleFunc("/orders", listOrdersForUser.Handler(orderService)) // /orders?user_id=42
	mux.HandleFunc("/socket", socket.Handler())

	return mux
}
