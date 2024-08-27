package server

import (
	"database/sql"
	"net/http"
	"orchard/pkg/persistence"
	"sync"
	"time"

	"github.com/gorilla/sessions"
)

type LoginRequest struct {
	Username string
	Hash     string
}

type User struct {
	Username     string
	PasswordHash string
}

// In-memory session store
var sessionStore = struct {
	sync.RWMutex
	sessions map[string]SessionData
}{
	sessions: make(map[string]SessionData),
}

type SessionData struct {
	Username string
	Expiry   time.Time
}

func HandleAccountConnection(db *sql.DB, w http.ResponseWriter, r *http.Request, store *sessions.CookieStore) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var req LoginRequest
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	hash := r.FormValue("hash")

	if username == "" || hash == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	req.Username = username
	req.Hash = hash

	// Check if session already exists in memory
	sessionStore.RLock()
	_, sessionExists := sessionStore.sessions[req.Username]
	sessionStore.RUnlock()
	if sessionExists {
		http.Error(w, "Session already exists", http.StatusBadRequest)
		return
	}

	// Fetch user details from the database
	user, err := persistence.GetAccount(db, req.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		} else {
			http.Error(w, "Failed to fetch user details", http.StatusInternalServerError)
		}
		return
	}

	if user.PasswordHash != req.Hash {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
		return
	}

	// Create session
	session, err := store.Get(r, "session-token")
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	sessionToken := session.ID
	session.Values["username"] = user.Username
	session.Values["authenticated"] = true
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "Failed to save session", http.StatusInternalServerError)
		return
	}

	// Store session in memory
	expiry := time.Now().Add(1 * time.Hour)
	sessionStore.Lock()
	sessionStore.sessions[sessionToken] = SessionData{
		Username: user.Username,
		Expiry:   expiry,
	}
	sessionStore.Unlock()

	// Set the session token as a cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "session-token",
		Value:   sessionToken,
		Expires: expiry,
		Path:    "/",
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User logged in successfully\n"))
}
