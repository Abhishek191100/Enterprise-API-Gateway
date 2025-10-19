package main

import (
	"log"
	"net/http"

	"github.com/Abhishek191100/Enterprise-API-Gateway/internal/proxy"
	"github.com/Abhishek191100/Enterprise-API-Gateway/internal/router"
)

func main() {
	routingTable, err := router.LoadRoutingTable("internal/config/gateway.yaml")
	if err != nil {
		log.Fatal("Failed to load routing table:", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		route, err := routingTable.Match(r)
		if err != nil {
			http.Error(w, "No matching route found", http.StatusNotFound)
			return
		}

		backend := route.NextBackend()
		if backend == "" {
			http.Error(w, "No backend available", http.StatusBadGateway)
			return
		}

		rp, err := proxy.NewReverseProxy(backend)
		if err != nil {
			http.Error(w, "Failed to create reverse proxy", http.StatusInternalServerError)
			return
		}
		rp.ServeHTTP(w, r)
	})

	/*rp, err := proxy.NewReverseProxy("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}*/

	log.Println("Gateway starting on :9090")
	if err := http.ListenAndServe(":9090", nil); err != nil {
		log.Fatal("Failed to start gateway:", err)
	}
}
