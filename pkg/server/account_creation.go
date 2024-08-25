package server

import (
	"database/sql"
	"net/http"
	"orchard/pkg/persistence"
)

func HandleAccountCreation(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	hash := r.FormValue("password_hash")

	if username == "" || email == "" || hash == "" {
		http.Error(w, "Missing required fields", http.StatusBadRequest)
		return
	}

	err = persistence.CreateAccount(db, username, email, hash)
	if err != nil {
		http.Error(w, "Failed to create account", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Account successfully created"))
}
