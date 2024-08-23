package persistence

import (
	"database/sql"
	"fmt"
	"orchard/pkg/utils"
	"sort"
)

type Account struct {
	Username     string
	Email        string
	PasswordHash string
	CreatedAt    string
	UpdatedAt    string
}

type FruitVegetable struct {
	Name     string
	Category string
}

type RegionStatistic struct {
	RegionName          string
	DiscoveryPercentage float64
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

func GetFruitOrVegetableNotInAccount(db *sql.DB, accountName string) (*FruitVegetable, error) {
	query := `
		SELECT fv.fruit_vegetable_name, fv.category
		FROM fruit_vegetables fv
		JOIN account_fruit_vegetables afv ON fv.fruit_vegetable_name = afv.fruit_vegetable_name
		WHERE afv.account_name = ?
		LIMIT 1;
	`

	var fv FruitVegetable
	err := db.QueryRow(query, accountName).Scan(&fv.Name, &fv.Category)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &fv, nil
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

func GetDescription(db *sql.DB, id string) (string, error) {
	query := `
		SELECT description 
		FROM information 
		WHERE fruit_vegetable_name = ?
	`

	var description string
	err := db.QueryRow(query, id).Scan(&description)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("no description found for item ID %d", id)
		}
		return "", err
	}

	return description, nil
}

func GetDiscoveryPercentagePerRegion(db *sql.DB, accountID string, category string, region string) (float64, error) {
	totalItemsQuery := `
		SELECT COUNT(*)
		FROM fruit_vegetables
		WHERE category = ? AND region_name = ?
	`
	var totalItems int
	err := db.QueryRow(totalItemsQuery, category, region).Scan(&totalItems)
	if err != nil {
		return 0, fmt.Errorf("failed to count total items in category %s and region %s: %w", category, region, err)
	}

	discoveredItemsQuery := `
		SELECT COUNT(*)
		FROM account_fruit_vegetables afv
		JOIN fruit_vegetables fv ON afv.fruit_vegetable_name = fv.fruit_vegetable_name
		WHERE fv.category = ? AND fv.region_name = ? AND afv.account_name = (
			SELECT username FROM accounts WHERE username = ?
		)
	`
	var discoveredItems int
	err = db.QueryRow(discoveredItemsQuery, category, region, accountID).Scan(&discoveredItems)
	if err != nil {
		return 0, fmt.Errorf("failed to count discovered items for account %d in category %s and region %s: %w", accountID, category, region, err)
	}

	if totalItems == 0 {
		return 0, nil
	}

	percentage := (float64(totalItems-discoveredItems) / float64(totalItems)) * 100
	return percentage, nil
}

func GetDiscoveryPercentage(db *sql.DB, accountID string, category string) (float64, error) {
	region, err := utils.GetCurrentRegion()
	if err != nil {
		return 0, fmt.Errorf("failed to get current region: %w", err)
	}

	return GetDiscoveryPercentagePerRegion(db, accountID, category, region)
}

func GetTopDiscoveryPercentage(db *sql.DB, accountID string, category string) ([]RegionStatistic, error) {
	regionsQuery := `
		SELECT region_name 
		FROM regions
	`

	rows, err := db.Query(regionsQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve regions: %w", err)
	}
	defer rows.Close()

	var stats []RegionStatistic
	for rows.Next() {
		var region string
		if err := rows.Scan(&region); err != nil {
			return nil, fmt.Errorf("failed to scan region: %w", err)
		}

		percentage, err := GetDiscoveryPercentagePerRegion(db, accountID, category, region)
		if err != nil {
			return nil, fmt.Errorf("failed to get discovery percentage for region %s: %w", region, err)
		}

		if percentage > 0 {
			stats = append(stats, RegionStatistic{
				RegionName:          region,
				DiscoveryPercentage: percentage,
			})
		}
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over regions: %w", err)
	}

	// Sort descending order
	sort.Slice(stats, func(i, j int) bool {
		return stats[i].DiscoveryPercentage < stats[j].DiscoveryPercentage
	})

	return stats[:3], nil
}
