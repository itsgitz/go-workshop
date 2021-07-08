package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

var (
	HTTPError = map[int]string{
		http.StatusInternalServerError: "Internal Server Error",
		http.StatusNotFound:            "Not Found",
	}
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/players", getPlayers)
	r.HandleFunc("/player/{id}", getPlayer)

	s := &http.Server{
		Handler:           r,
		Addr:              ":8000",
		ReadHeaderTimeout: 2 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       30 * time.Second,
	}

	log.Println("[*] Start go http server on port :8000")
	log.Fatal(s.ListenAndServe())
}

// Players is the data collection for player details
type Players struct {
	ID    uint32 `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Role  string `json:"role,omitempty"`
	Level uint32 `json:"level,omitempty"`
	Token string `json:"token,omitempty"`
}

// playerDataSeeder just for create the dummy data of players
var playerDataSeeder = []Players{
	{
		ID:    1,
		Name:  "ItsGitz",
		Role:  "Pusher",
		Level: 70,
		Token: uuid.NewString(),
	},
	{
		ID:    2,
		Name:  "pleasantvan",
		Role:  "Pusher",
		Level: 70,
		Token: uuid.NewString(),
	},
	{
		ID:    3,
		Name:  "MrCrazy",
		Role:  "Support",
		Level: 80,
		Token: uuid.NewString(),
	},
}

// getPlayers show the list of existing players
func getPlayers(w http.ResponseWriter, r *http.Request) {
	log.Println("players:", playerDataSeeder)

	// encode to JSON data format
	j, err := json.Marshal(playerDataSeeder)
	if err != nil {
		http.Error(w, HTTPError[http.StatusInternalServerError], http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}

// getPlayer show a single player with given ID
func getPlayer(w http.ResponseWriter, r *http.Request) {
	index := make(map[bool]int, len(playerDataSeeder))

	getID := mux.Vars(r)["id"]
	id, err := strconv.ParseUint(getID, 10, 32)
	if err != nil {
		http.Error(w, HTTPError[http.StatusInternalServerError], http.StatusInternalServerError)
		return
	}
	log.Println("Get path variables (converted to integer):", id)

	// search the data where ID = vars["id"]
	for i := range playerDataSeeder {
		log.Println("playerDataSeeder:", playerDataSeeder[i])
		if playerDataSeeder[i].ID == uint32(id) {
			index[true] = i
			break
		}
	}

	if len(index) == 0 {
		log.Println("Not found!")
		http.Error(w, HTTPError[http.StatusNotFound], http.StatusNotFound)
		return
	}
	log.Println("Get Index:", index)
	log.Println("Found index at:", index)

	// encode to JSON data format
	j, err := json.Marshal(playerDataSeeder[index[true]])
	if err != nil {
		http.Error(w, HTTPError[http.StatusInternalServerError], http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(j)
}
