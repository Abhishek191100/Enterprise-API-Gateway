package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type APIResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Time    time.Time   `json:"time"`
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Simulate processing delay
	time.Sleep(150 * time.Millisecond)

	products := []map[string]interface{}{
		{"id": 101, "name": "server4laptop", "price": 999.99},
		{"id": 102, "name": "server4mouse", "price": 29.99},
		{"id": 103, "name": "server4keyboard", "price": 79.99},
	}

	response := APIResponse{
		Message: "Products fetched successfully",
		Data:    products,
		Time:    time.Now(),
	}

	json.NewEncoder(w).Encode(response)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/api/products", productsHandler)
	log.Printf("Products Server 4 starting on :8084")
	log.Fatal(http.ListenAndServe(":8084", mux))
}
