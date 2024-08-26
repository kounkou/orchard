package server

import (
	"database/sql"
	"net/http"
	"orchard/pkg/persistence"
	"time"

	"github.com/gorilla/sessions"
)

type User struct {
	Username     string
	PasswordHash string
}

type LoginRequest struct {
	Username string `json:"username"`
	Hash 	 string `json:"hash"`
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

	// If session already contains data, abort request
	if persistence.SessionExistsForUser(req.Username, db) {
		return
	}

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

	session, err := store.Get(r, "session-token")
	if err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}
	
	session.Values["username"] = user.Username
	session.Values["authenticated"] = true
	err = session.Save(r, w)
	if err != nil {
		http.Error(w, "Failed to save session", http.StatusInternalServerError)
		return
	}
	
	sessionTokenCookie, err := r.Cookie("session-token")
    if err != nil {
        http.Error(w, "Failed to retrieve session cookie", http.StatusInternalServerError)
        return
    }
    sessionToken := sessionTokenCookie.Value

	// Store session token in the database
	if sessionToken == "" {
		http.Error(w, "Failed to generate session ID", http.StatusInternalServerError)
		return
	}
	expiry := time.Now().Add(1 * time.Hour)

	insertQuery := `
		INSERT INTO sessions (username, session_token, expiry) 
		VALUES (?, ?, ?)
	`
	_, err = db.Exec(insertQuery, user.Username, sessionToken, expiry)
	if err != nil {
		http.Error(w, "Failed to save session token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "session-token",
		Value:   sessionToken,
		Expires: time.Now().Add(1 * time.Hour),
		Path:    "/",
	})

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User logged in successfully\n"))
}
