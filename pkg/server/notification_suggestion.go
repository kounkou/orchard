package server

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"orchard/pkg/notifier"
	"orchard/pkg/persistence"
	"os"
	"strings"
)

/*
TODO:
Basically, notifications could be triggered everyday, at specific time (let the user decide when...)
Another way to have smarter notifications is to trigger when specific conditions are met :
  - Example use case is if the user is at proximity to a grossery store (requires access to locations)
  - Another use case is if the user changes country for instance during vacation, it could make sense to suggestion fruits based on the location
*/
func HandleSuggestionNotificationRequest(db *sql.DB, w http.ResponseWriter, r *http.Request) {
    cookie, err := r.Cookie("session-token")
    if err != nil {
        if err == http.ErrNoCookie {
            http.Error(w, "Unauthorized: Missing session token", http.StatusUnauthorized)
            return
        }
        http.Error(w, "Error retrieving session token", http.StatusBadRequest)
        return
    }

    sessionToken := cookie.Value
    if !persistence.IsValidSessionToken(sessionToken, db) {
        http.Error(w, "Unauthorized: Invalid session token", http.StatusUnauthorized)
        return
    }

	hash := strings.TrimPrefix(r.URL.Path, "/notification-suggestion-")
	if hash == "" {
		http.Error(w, "Missing or invalid hash", http.StatusBadRequest)
		return
	}

	notifier.CreateDefaultNotifications(db, hash)

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
