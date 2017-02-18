package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net"
	"net/http"
	"os"
)

type StatusUpdate struct {
	Status int `json:"status"`
}

const (
	version = "0.1"
)

var healthStatus int

func init() {
	healthStatus = 200
}

func main() {
	port := ":8080"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	log.Printf("[INFO] %s version %s started", os.Args[0], version)
	log.Printf("[INFO] Listening on %s", port)
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			log.Printf("[INFO] Received health request, returning %d", healthStatus)
			w.WriteHeader(healthStatus)
		case http.MethodPost:
			status := StatusUpdate{}
			decoder := json.NewDecoder(r.Body)
			err := decoder.Decode(&status)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				log.Printf("[INFO] Health status updated to %d", status.Status)
				healthStatus = status.Status
				w.WriteHeader(http.StatusOK)
			}
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)

		}
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q\n", html.EscapeString(r.URL.Path))
		hostname, _ := os.Hostname()
		fmt.Fprintf(w, "HOST: %s\n", hostname)
		fmt.Fprintf(w, "ADDRESSES:\n")
		addrs, _ := net.InterfaceAddrs()
		for _, addr := range addrs {
			fmt.Fprintf(w, "    %s\n", addr.String())
		}
	})

	log.Fatal(http.ListenAndServe(port, nil))
}
