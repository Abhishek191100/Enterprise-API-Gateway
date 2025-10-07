package main

import (
	"log"
	"net/http"

	"github.com/Abhishek191100/Enterprise-API-Gateway/internal/proxy"
)

func main() {
	rp, err := proxy.NewReverseProxy("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Gateway starting on :9090")
	if err := http.ListenAndServe(":9090", rp); err != nil {
		log.Fatal("Failed to start gateway:", err)
	}
}
