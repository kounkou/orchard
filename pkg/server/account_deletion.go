package server

import (
	"database/sql"
	"net/http"
	"orchard/pkg/persistence"
)

func HandleAccountDeletion(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
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

	err = persistence.DeleteAccount(db, accountID)
	if err != nil {
		http.Error(w, "Failed to delete account", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Account successfully deleted"))
}
