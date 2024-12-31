package server

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"orchard/pkg/notifier"
	"orchard/pkg/persistence"
	"strings"
	"time"
)

func getSuggestions(db *sql.DB, accountHash string) ([]persistence.FruitVegetable, error) {
	account, err := persistence.GetAccountByPasswordHash(db, accountHash)
	if err != nil {
		log.Fatalf("Failed to get account: %v", err)
	}

	fv, err := persistence.GetFruitsOrVegetableNotInAccount(db, account.Username)
	if err != nil {
		log.Printf("Error retrieving fruit or vegetable for account %d: %v", account.Username, err)
		return nil, nil
	}

	return fv, nil
}

func HandleSuggestionsRequest(db *sql.DB, w http.ResponseWriter, r *http.Request) {
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

	hash := strings.TrimPrefix(r.URL.Path, "/get-suggestions-")
	if hash == "" {
		http.Error(w, "Missing or invalid hash", http.StatusBadRequest)
		return
	}

	items, err := getSuggestions(db, hash)
	if err != nil {
		log.Fatalf("Error getting suggestions from the database", err)
		return
	}

	notification, err := notifier.CreateItemsSuggestionNotification(items)
	if err != nil {
		return
	}

	jsonData, err := json.Marshal(notification)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}