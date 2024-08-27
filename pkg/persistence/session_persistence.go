package persistence

import (
	"database/sql"
	"log"
)

func SessionExistsForUser(accountID string, db *sql.DB) bool {
	query := `
		SELECT EXISTS (
            SELECT 1
            FROM sessions
            WHERE username = ?
        )
	`

    var exists bool
    err := db.QueryRow(query, accountID).Scan(&exists)
    if err != nil {
		if err == sql.ErrNoRows {
            log.Println("Session Token not found:", accountID)
            return false
        }
        return false
    }

    return exists
}