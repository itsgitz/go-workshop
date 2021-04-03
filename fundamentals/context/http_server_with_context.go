package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("server: Start the web server on port :8000")

	mux := http.NewServeMux()
	mux.Handle("/hello", helloMiddleware(http.HandlerFunc(helloController)))
	mux.Handle("/world", helloMiddleware(http.HandlerFunc(worldController)))

	server := http.Server{
		Addr:         ":8000",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}

	log.Fatal(server.ListenAndServe())
}

func helloController(w http.ResponseWriter, r *http.Request) {
	// define the context
	ctx := r.Context()

	log.Println("server: context value from helloMiddleware (helloController):", ctx.Value("greeting"))
	log.Println("server: hello controller started")
	defer log.Println("server: hello controller ended")

	select {
	case <-time.After(10 * time.Second):
		fmt.Fprintf(w, "hello bro\n")
	case <-ctx.Done():
		err := ctx.Err()
		log.Println("server:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func worldController(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Println("server: context value from helloMiddleware (worldController):", ctx.Value("greeting"))
	w.Write([]byte("world!"))
}

func helloMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello Middleware")

		// set the context value in helloMiddleware
		ctx := r.Context()
		r = r.WithContext(context.WithValue(ctx, "greeting", "hello world!"))

		next.ServeHTTP(w, r)
	})
}
