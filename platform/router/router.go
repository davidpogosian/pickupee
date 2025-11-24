package router

import (
	"net/http"

	"github.com/davidpogosian/pickupee/web/api/createOrder"
	"github.com/davidpogosian/pickupee/web/api/getOrder"
)

func Create() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/createOrder", createOrder.Handler())
	mux.HandleFunc("/getOrder", getOrder.Handler())

	return mux
}
