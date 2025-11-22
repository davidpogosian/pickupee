package main

import (
	"fmt"
	"net/http"

	"github.com/davidpogosian/pickupee/platform/router"
)

func main() {
	router := router.Create()

	fmt.Println("Server starting at http://localhost:8080")
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
