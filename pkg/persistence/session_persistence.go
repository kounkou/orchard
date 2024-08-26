package persistence

import (
	"database/sql"
	"log"
)

func IsValidSessionToken(sessionToken string, db *sql.DB) bool {
    query := `
        SELECT EXISTS (
            SELECT 1
            FROM sessions
            WHERE expiry > time(time(), 'localtime') AND session_token = ?
        )
    `

    var exists bool
    err := db.QueryRow(query, sessionToken).Scan(&exists)
    if err != nil {
		if err == sql.ErrNoRows {
            log.Println("Session Token not found or expired:", sessionToken)
            return false
        }
        return false
    }

    return exists
}

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