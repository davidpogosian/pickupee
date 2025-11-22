package createOrder

import (
	"fmt"
	"net/http"
)

func Handler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "You have reched createOrder endpoint!")
	}
}
