package server

import (
	"database/sql"
	"net/http"
	"orchard/pkg/persistence"
	"strings"
)

func HandleAddUnknownItems(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	hash := strings.TrimPrefix(r.URL.Path, "/add-unknown-items-")
	if hash == "" {
		http.Error(w, "Missing or invalid hash", http.StatusBadRequest)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}

	accountID := r.FormValue("account_id")
	if accountID == "" {
		http.Error(w, "Missing account ID", http.StatusBadRequest)
		return
	}

	itemNames := r.Form["fruit_vegetable_name"]

	if len(itemNames) == 0 {
		http.Error(w, "Item names must be provided", http.StatusBadRequest)
		return
	}

	err = persistence.AddUnknownItems(db, accountID, itemNames)
	if err != nil {
		http.Error(w, "Failed to prepare statement", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Items successfully added"))
}
