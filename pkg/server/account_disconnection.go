package server

import (
	"database/sql"
	"net/http"
	"time"
)

func HandleAccountDisconnection(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Assume session-token is sent in the request header
	cookie, err := r.Cookie("session-token")
	if err != nil {
		http.Error(w, "Session token not found", http.StatusUnauthorized)
		return
	}

	sessionToken := cookie.Value

	// Check if session exists in memory
	sessionStore.Lock()
	sessionData, sessionExists := sessionStore.sessions[sessionToken]
	if !sessionExists || sessionData.Expiry.Before(time.Now()) {
		sessionStore.Unlock()
		http.Error(w, "Invalid or expired session", http.StatusUnauthorized)
		return
	}

	// Remove session from in-memory store
	delete(sessionStore.sessions, sessionToken)
	sessionStore.Unlock()

	// Clear the session cookie
	http.SetCookie(w, &http.Cookie{
		Name:   "session-token",
		Value:  "",
		MaxAge: -1,
		Path:   "/",
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User logged off successfully\n"))
}