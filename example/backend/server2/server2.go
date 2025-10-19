package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type server2APIResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Time    time.Time   `json:"time"`
}

func server2usersHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	users := []map[string]interface{}{
		{"id": 1, "name": "SERVER20", "email": "abhishek@example.com"},
		{"id": 2, "name": "SERVER21", "email": "vv185097@ncr.com"},
	}

	response := server2APIResponse{
		Message: "users lists fetched successfully",
		Data:    users,
		Time:    time.Now(),
	}
	json.NewEncoder(w).Encode(response)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/api/users", server2usersHandler)
	log.Printf("Users Server 2 starting on :8082")
	log.Fatal(http.ListenAndServe(":8082", mux))

}
