package main

import (
	"log"
	"net/http"

	"github.com/davidpogosian/pickupee/platform/initialization"
)

func main() {
	// Make r
	db, r := initialization.CreateServer("./mydb.db")
	defer db.Close()

	// Launch server
	log.Println("Server starting at http://localhost:8080")
	err := http.ListenAndServe(":8080", r)
	if err != nil {
		panic(err)
	}
}
