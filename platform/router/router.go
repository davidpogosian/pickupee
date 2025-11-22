package router

import (
	"net/http"

	"github.com/davidpogosian/pickupee/web/api/createOrder"
)

func Create() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/createOrder", createOrder.Handler())

	return mux
}
