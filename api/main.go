package main

import (
	"log"
	"net/http"
	"fmt"
)

func main() {
	server := http.Server{}
    server.Addr = ":8080"
    http.HandleFunc("/health", healthHandler)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %s", err)
	}
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    r.Close = true
    if r.Method != http.MethodGet {
        w.WriteHeader(http.StatusMethodNotAllowed)
        header := w.Header()
        header.Set("Allowed", http.MethodGet)
        w.Write([]byte(fmt.Sprintf("Bad Request: Method %s not allowed for resource.", r.Method)))
        return
    }
    
    _, err := w.Write([]byte("OK"))
    if err != nil {
        log.Fatalf("Failed to write response: %s", err)   
    }
}