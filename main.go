package main

import (
	"fmt"
	"log"
	"net/http"
	"receipt-processor/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new router
	r := mux.NewRouter()

	// Set up routes and their handlers
	r.HandleFunc("/receipts/process", handlers.ProcessReceiptHandler).Methods("POST")
	r.HandleFunc("/receipts/{id}/points", handlers.GetPointsHandler).Methods("GET")

	// Start the server on port 8080
	fmt.Println("Server started at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
