package server

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"orchard/pkg/notifier"
	"os"
	"strings"
)

func HandleCompletionNotificationRequest(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	hash := strings.TrimPrefix(r.URL.Path, "/notification-completion-")
	if hash == "" {
		http.Error(w, "Missing or invalid hash", http.StatusBadRequest)
		return
	}

	notifier.CreateStatsNotifications(db, hash)

	file, err := os.Open(hash + "-notification.json")
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
