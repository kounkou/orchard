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

	// Serve static files for images
    // Here, the "/images/" prefix will be stripped, and the actual file path will be served
    fileServer := http.FileServer(http.Dir("./images"))
	mux.Handle("/images/", http.StripPrefix("/images/", fileServer))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch {

		// Notifications
		// [Challenge] Handle suggestion triggers a challenge, when user has discovered <= 2 fruits/veggies 
		// so that we can gather the current state of knowledge of the user
		case strings.HasPrefix(r.URL.Path, "/notification-suggestion-"):
			HandleSuggestionNotificationRequest(db, w, r)
		case strings.HasPrefix(r.URL.Path, "/notification-stats-"):
			HandleDiscoveryStatistics(db, w, r)
		case strings.HasPrefix(r.URL.Path, "/notification-completion-"):
			HandleCompletionNotificationRequest(db, w, r)

		// Account management handlers
		// [Challenge] Account creation triggers a challenge so that we can gather the current
		// state of knowledge of the user
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
		case strings.HasPrefix(r.URL.Path, "/mark-item-known-"):
			HandleMarkItemAsKnown(db, w, r)

		// TODO: Application level improvement
		case strings.HasPrefix(r.URL.Path, "/submit-new-item-for-review-"):
			HandleSubmitNewItemForReview(db, w, r)

		// Engagement
		// HandleSuggestionsRequest returns a list of Suggestions, we then use the 
		// suggestions to make a challenge to the user
		case strings.HasPrefix(r.URL.Path, "/get-suggestions-"):
			HandleSuggestionsRequest(db, w, r)
		// This API just retrieves the notification that is already written to notification.json
		// notification.json is generated in 2 cases : 
		// 1- When number of fruits/veggies decreases to less or equal to 2
		// 2- When we create an account
		case strings.HasPrefix(r.URL.Path, "/get-notification-"):
			HandleGetNotificationRequest(db, w, r)

		default:
			http.NotFound(w, r)
		}
	})

	log.Println("Server starting on :8000")
	log.Fatal(http.ListenAndServe(":8000", mux))
}
