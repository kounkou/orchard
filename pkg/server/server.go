package server

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"orchard/pkg/notifier"
	"os"
	"strings"
)

type Notification struct {
	Notification string `json:"notification"`
}

func handleNotificationRequest(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	hash := strings.TrimPrefix(r.URL.Path, "/notification-")
	if hash == "" {
		http.Error(w, "Missing or invalid hash", http.StatusBadRequest)
		return
	}

	notifier.CreateDefaultNotifications(db, hash)

	file, err := os.Open(hash + "-notification-default.json")
	if err != nil {
		http.Error(w, "Could not open JSON file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var notification Notification
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&notification); err != nil {
		http.Error(w, "Could not decode JSON file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(notification); err != nil {
		http.Error(w, "Could not encode response to JSON", http.StatusInternalServerError)
	}
}

func handleNewFruitDiscovered(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	hash := strings.TrimPrefix(r.URL.Path, "/new-fruit-discovered-")
	if hash == "" {
		http.Error(w, "Missing or invalid hash", http.StatusBadRequest)
		return
	}

	notifier.CreateStatsNotifications(db, hash)

	file, err := os.Open(hash + "-notification-stats.json")
	if err != nil {
		http.Error(w, "Could not open JSON file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var notification Notification
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&notification); err != nil {
		http.Error(w, "Could not decode JSON file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(notification); err != nil {
		http.Error(w, "Could not encode response to JSON", http.StatusInternalServerError)
	}
}

func Start(db *sql.DB) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/notification-") {
			handleNotificationRequest(db, w, r)
		} else if strings.HasPrefix(r.URL.Path, "/new-fruit-discovered-") {
			handleNewFruitDiscovered(db, w, r)
		} else {
			http.NotFound(w, r)
		}
	})

	log.Println("Server starting on :8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
