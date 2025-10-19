package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type server1APIResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Time    time.Time   `json:"time"`
}

func server1usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users := []map[string]interface{}{
		{"id": 1, "name": "XYZZZ", "email": "abhishek@example.com"},
		{"id": 2, "name": "HJKL", "email": "vv185097@ncr.com"},
	}

	response := server1APIResponse{
		Message: "users lists fetched successfully",
		Data:    users,
		Time:    time.Now(),
	}
	json.NewEncoder(w).Encode(response)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/api/users", server1usersHandler)
	log.Printf("Users Server 1 starting on :8081")
	log.Fatal(http.ListenAndServe(":8081", mux))
}
