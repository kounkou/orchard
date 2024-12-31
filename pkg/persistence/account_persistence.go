package persistence

import (
	"database/sql"
	"fmt"
)

type User struct {
	Username     string
	PasswordHash string
}

func GetAccounts(db *sql.DB) ([]Account, error) {
	rows, err := db.Query("SELECT username, email, password_hash, created_at, updated_at FROM accounts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []Account
	for rows.Next() {
		var account Account
		if err := rows.Scan(&account.Username, &account.Email, &account.PasswordHash, &account.CreatedAt, &account.UpdatedAt); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func GetAccountByPasswordHash(db *sql.DB, passwordHash string) (*Account, error) {
	query := `
		SELECT username, email, password_hash, created_at, updated_at
		FROM accounts
		WHERE password_hash = ?
		LIMIT 1;
	`

	var account Account
	err := db.QueryRow(query, passwordHash).Scan(
		&account.Username,
		&account.Email,
		&account.PasswordHash,
		&account.CreatedAt,
		&account.UpdatedAt,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, fmt.Errorf("failed to retrieve account: %w", err)
	}

	return &account, nil
}

func GetAccountPasswordHash(db *sql.DB, accountID string) (string, error) {
	query := `
		SELECT password_hash 
		FROM accounts 
		WHERE username = ?
	`

	var passwordHash string
	err := db.QueryRow(query, accountID).Scan(&passwordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no account found with ID %d", accountID)
		}
		return "", err
	}

	return passwordHash, nil
}

func GetAccount(db *sql.DB, accountID string) (User, error) {
	var user User

	query := `SELECT username, password_hash FROM accounts WHERE username = ?`
	err := db.QueryRow(query, accountID).Scan(&user.Username, &user.PasswordHash)
	if err != nil {
		return User{}, err
	}

	return user, err
}

func CreateAccount(db *sql.DB, username string, email string, hash string) error {
	query := "INSERT INTO accounts (username, email, password_hash) VALUES (?, ?, ?)"
	_, err := db.Exec(query, username, email, hash)
	if err != nil {
		return err
	}

	return nil
}

func DeleteAccount(db *sql.DB, accountID string) error {
	query := "DELETE FROM accounts WHERE username = ?"
	result, err := db.Exec(query, accountID)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return err
	}

	return nil
}