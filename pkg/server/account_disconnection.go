package server

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/sessions"
)

func HandleAccountDisconnection(db *sql.DB, w http.ResponseWriter, r *http.Request, store *sessions.CookieStore) {
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

	var req LoginRequest

	err = r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")

	if username == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	req.Username = username

	var sessionToken string

	query := `SELECT session_token FROM sessions WHERE username = ?`
	err = db.QueryRow(query, req.Username).Scan(&sessionToken)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Invalid username", http.StatusUnauthorized)
		} else {
			http.Error(w, "Failed to fetch user details", http.StatusInternalServerError)
		}
		return
	}

	if err != nil || sessionToken != cookie.Value {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}

	session, err := store.New(r, "session-token")
	// Invalidate the session
	session.Values["authenticated"] = false
	session.Save(r, w)

	query = `DELETE FROM sessions WHERE session_token = ?`
	_, err = db.Exec(query, sessionToken)
	if err != nil {
		http.Error(w, "Failed to delete session ", http.StatusInternalServerError)
		return
	}

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