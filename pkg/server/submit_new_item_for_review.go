package server

import (
	"database/sql"
	"net/http"
)

/*
Users can submit new fruits, vegetable or future new categories to enhance the
application glablly. If approved, the submission will be added to the
onboarding/challenge for new users and therefore, make the application more
engaging with new items to discover
*/
func HandleSubmitNewItemForReview(db *sql.DB, w http.ResponseWriter, r *http.Request) {
}
