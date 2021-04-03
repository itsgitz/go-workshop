package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/home", baseMiddleware(http.HandlerFunc(home)))
	mux.Handle("/songs", baseMiddleware(http.HandlerFunc(getSongs)))
	mux.Handle("/playlists", baseMiddleware(http.HandlerFunc(getPlaylists)))

	server := &http.Server{
		Addr:              ":8000",
		ReadHeaderTimeout: 2 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
		Handler:           mux,
	}

	log.Println("[*] Start web server on port :8000")
	log.Fatal(server.ListenAndServe())
}

// baseMiddleware is base middleware handler for this apps
func baseMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("[*] From baseMiddleware")
		next.ServeHTTP(w, r)
	})
}

// Songs data collection
type Songs struct {
	ID     string `json:"id,omitempty"`
	Title  string `json:"title,omitempty"`
	Artist string `json:"artist,omitempty"`
	Album  string `json:"album,omitempty"`
}

// Playlists data collection
type Playlists struct {
	ID    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Songs uint32 `json:"songs,omitempty"`
}

// home() show the home page (as json)
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hey! Welcome home, buddy!"))
}

// getSongs() show the list of songs
func getSongs(w http.ResponseWriter, r *http.Request) {
	// songs stores the list of songs in array struct
	// example output:
	// [{162f14f4-bf53-4b2e-b196-aa8b0307a437 The Reason Hoobastank The Reason} {82c8b8e2-5de0-47d2-bf05-b72a8c090029 Still into You Paramore Paramore}]
	songs := []Songs{
		{
			ID:     uuid.NewString(),
			Title:  "The Reason",
			Artist: "Hoobastank",
			Album:  "The Reason",
		},
		{
			ID:     uuid.NewString(),
			Title:  "Still into You",
			Artist: "Paramore",
			Album:  "Paramore",
		},
	}

	// just print the songs data on the console
	log.Println("songs:", songs)

	// encode to json data format
	encJSON, err := encodeJSON(songs)
	if err != nil {
		showErrorResponse(w)
	}

	// write header on as application/json
	w.Header().Set("Content-Type", "application/json")
	w.Write(encJSON)
}

// getPlaylists show the playlists
func getPlaylists(w http.ResponseWriter, r *http.Request) {
	// playlists stores the playlists in array struct
	// example output:
	// [{7c9558a1-21c8-49b2-b6f4-971d4cd1cc07 Memory Leaks 187} {03806459-a444-4893-933a-eca0ee8b51ce J - Coffee 97}]
	playlists := []Playlists{
		{
			ID:    uuid.NewString(),
			Name:  "Memory Leaks",
			Songs: 187,
		},
		{
			ID:    uuid.NewString(),
			Name:  "J - Coffee",
			Songs: 97,
		},
	}
	log.Println("playlists:", playlists)

	// encode to json data format
	encJSON, err := encodeJSON(playlists)
	if err != nil {
		showErrorResponse(w)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(encJSON)
}

// encodeJSON() encode the data collection, in this case struct data type to json data format
func encodeJSON(collection interface{}) ([]byte, error) {
	enc, err := json.Marshal(collection)
	if err != nil {
		return nil, err
	}
	return enc, nil
}

// showErrorResponse() return or display the error page if any
func showErrorResponse(w http.ResponseWriter) {
	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	return
}
