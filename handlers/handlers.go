package handlers

import (
	"encoding/json"
	"net/http"
	"receipt-processor/models"
	"receipt-processor/utils"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Store receipts in memory
var receipts = make(map[string]models.Receipt)

// ProcessReceiptHandler handles the POST request to process a receipt
func ProcessReceiptHandler(w http.ResponseWriter, r *http.Request) {
	var receipt models.Receipt
	if err := json.NewDecoder(r.Body).Decode(&receipt); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	// Generate a unique ID for the receipt
	id := uuid.New().String()

	// Store receipt in memory
	receipts[id] = receipt

	// Respond with the generated ID
	response := models.ProcessResponse{ID: id}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetPointsHandler handles the GET request to get points of a receipt by its ID
func GetPointsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// Retrieve the receipt by ID
	receipt, found := receipts[id]
	if !found {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	// Calculate points for the receipt
	points, err := utils.CalculatePoints(receipt)

	if err != nil {
		// Handle error
	}
	// Respond with the points
	response := models.PointsResponse{Points: points}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
