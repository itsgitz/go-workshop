package main

import (
	"log"
	"net/http"
	"time"

	"gl.atisicloud.com/dev/go-restful-server/handler"
)

func main() {
	http.HandleFunc("/api/v1", handler.HomeHandler)             // GET
	http.HandleFunc("/api/v1/vm", handler.GetVMServicesHandler) // GET

	log.Println("[*] Run go http server on port :8000")

	// custom server
	myServer := &http.Server{
		Addr:         ":8000",
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	err := myServer.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
