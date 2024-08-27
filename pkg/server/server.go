package server

import (
	"database/sql"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
)

type Notification struct {
	Notification string `json:"notification"`
}

func Start(db *sql.DB) {
	var store = sessions.NewCookieStore([]byte("secret-key"))

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch {

		// Notifications
		case strings.HasPrefix(r.URL.Path, "/notification-suggestion-"):
			HandleSuggestionNotificationRequest(db, w, r)
		case strings.HasPrefix(r.URL.Path, "/notification-stats-"):
			HandleDiscoveryStatistics(db, w, r)
		case strings.HasPrefix(r.URL.Path, "/notification-completion-"):
			HandleCompletionNotificationRequest(db, w, r)

		// Account management handlers
		case strings.HasPrefix(r.URL.Path, "/create-account-"):
			HandleAccountCreation(db, w, r)
		case strings.HasPrefix(r.URL.Path, "/delete-account-"):
			HandleAccountDeletion(db, w, r)
		case strings.HasPrefix(r.URL.Path, "/connect"):
			HandleAccountConnection(db, w, r, store)
		case strings.HasPrefix(r.URL.Path, "/disconnect"):
			HandleAccountDisconnection(db, w, r)

		// Account data managment
		case strings.HasPrefix(r.URL.Path, "/add-unknown-items-"):
			HandleAddUnknownItems(db, w, r)
		case strings.HasPrefix(r.URL.Path, "/delete-unknown-items-"):
			HandleDeleteUnknownItems(db, w, r)

		// TODO: Application level improvement
		case strings.HasPrefix(r.URL.Path, "/submit-new-item-for-review-"):
			HandleSubmitNewItemForReview(db, w, r)

		default:
			http.NotFound(w, r)
		}
	})

	log.Println("Server starting on :8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
