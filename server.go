package main

import (
	"fmt"
	"html"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
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

	log.Fatal(http.ListenAndServe(":8080", nil))

}
