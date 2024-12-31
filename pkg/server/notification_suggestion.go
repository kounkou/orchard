package server

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"orchard/pkg/notifier"
	"os"
	"strings"
	"time"
)

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

	// Validate session token in memory
	sessionStore.RLock()
	sessionData, sessionExists := sessionStore.sessions[sessionToken]
	sessionStore.RUnlock()

	if !sessionExists || sessionData.Expiry.Before(time.Now()) {
		http.Error(w, "Unauthorized: Invalid or expired session token", http.StatusUnauthorized)
		return
	}

	hash := strings.TrimPrefix(r.URL.Path, "/notification-suggestion-")
	if hash == "" {
		http.Error(w, "Missing or invalid hash", http.StatusBadRequest)
		return
	}

	// Create default notifications based on hash
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
