package router

import (
	"net/http"

	"github.com/davidpogosian/pickupee/service"
	"github.com/davidpogosian/pickupee/web/api/orders"
	"github.com/davidpogosian/pickupee/web/socket"
)

func Create(orderService *service.OrderService) *http.ServeMux {
	mux := http.NewServeMux()

	// TODO: Make a separate /orders mux
	mux.Handle("/orders", &orders.OrdersHandler{Service: orderService}) // /orders?user_id=42
	mux.HandleFunc("/socket", socket.Handler())

	return mux
}
