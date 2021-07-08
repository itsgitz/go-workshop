package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Go HTTP server with goroutine function for start the server")

	mux := http.NewServeMux()
	mux.HandleFunc("/hello", helloHandler)

	server := &http.Server{
		Addr:              ":8000",
		Handler:           mux,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 60 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	go func() {
		log.Fatal(server.ListenAndServe())
	}()
	select {}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello Handler"))
}
