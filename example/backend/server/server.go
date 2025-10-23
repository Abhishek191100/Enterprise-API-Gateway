package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type HealthAPIResponse struct {
	Status    string    `json:"status"`
	Timestamp time.Time `json:"timestamp"`
	Service   string    `json:"service"`
	Version   string    `json:"version"`
}

type APIResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Time    time.Time   `json:"time"`
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/health", healthHandler)

	mux.HandleFunc("/api/users", usersHandler)

	mux.HandleFunc("/api/products", productsHandler)

	mux.HandleFunc("/api/echo", echoHandler)

	log.Printf("Server starting on :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		//log.Printf("Server started on :8080")
		log.Fatalf("Server failed to start: %v", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := HealthAPIResponse{
		Status:    "healthy",
		Timestamp: time.Now(),
		Service:   "ExampleBackend",
		Version:   "1.0.0",
	}
	json.NewEncoder(w).Encode(response)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users := []map[string]interface{}{
		{"id": 1, "name": "Abhishek", "email": "abhishek@example.com"},
		{"id": 2, "name": "Venkat", "email": "vv185097@ncr.com"},
	}

	response := APIResponse{
		Message: "users lists fetched successfully",
		Data:    users,
		Time:    time.Now(),
	}
	json.NewEncoder(w).Encode(response)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Simulate processing delay
	time.Sleep(150 * time.Millisecond)

	products := []map[string]interface{}{
		{"id": 101, "name": "Laptop", "price": 999.99},
		{"id": 102, "name": "Mouse", "price": 29.99},
		{"id": 103, "name": "Keyboard", "price": 79.99},
	}

	response := APIResponse{
		Message: "Products fetched successfully",
		Data:    products,
		Time:    time.Now(),
	}

	json.NewEncoder(w).Encode(response)
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	echo := map[string]interface{}{
		"method":     r.Method,
		"path":       r.URL.Path,
		"headers":    r.Header,
		"remoteAddr": r.RemoteAddr,
		"time":       time.Now(),
	}

	json.NewEncoder(w).Encode(echo)
}
