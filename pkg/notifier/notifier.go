package notifier

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"orchard/pkg/persistence"
	"os"
	"path/filepath"
	"strings"
)

type Notification struct {
	Notification string `json:"notification"`
}

func SendNotification(notification Notification, passwordHash string, notifType string) {
	dataDir := "."
	if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
		log.Fatalf("Error creating data directory: %v", err)
	}

	filePath := filepath.Join(dataDir, passwordHash+"-notification-"+notifType+".json")

	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	notificationJSON, err := json.MarshalIndent(notification, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling notification: %v", err)
	}

	if _, err := file.Write(notificationJSON); err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	log.Printf("Notification written to %s", filePath)
}

func CreateDefaultNotifications(db *sql.DB, accountHash string) {
	account, err := persistence.GetAccountByPasswordHash(db, accountHash)
	if err != nil {
		log.Fatalf("Failed to get account: %v", err)
	}

	fv, err := persistence.GetFruitOrVegetableNotInAccount(db, account.Username)
	if err != nil {
		log.Printf("Error retrieving fruit or vegetable for account %d: %v", account.Username, err)
		return
	}

	if fv != nil {
		description, err := persistence.GetDescription(db, fv.Name)
		if err != nil {
			log.Fatal("Failed to get the description ", err)
		}

		notification := Notification{
			Notification: fmt.Sprintf(
				"Discover the %s! %s Learn more about the %s in Orchard!",
				fv.Name, description, fv.Name,
			),
		}

		passwordHash, err := persistence.GetAccountPasswordHash(db, account.Username)
		if err != nil {
			log.Fatal("Failed to get the password hash ", err)
		}

		SendNotification(notification, passwordHash, "default")
	} else {
		log.Printf("Account: %s, No more fruits or vegetables available\n", account.Username)
	}
}

func CreateStatsNotifications(db *sql.DB, accountHash string) {
	account, err := persistence.GetAccountByPasswordHash(db, accountHash)
	if err != nil {
		log.Fatalf("Failed to get account: %v", err)
	}

	fv, err := persistence.GetFruitOrVegetableNotInAccount(db, account.Username)
	if err != nil {
		log.Printf("Error retrieving fruit or vegetable for account %d: %v", account.Username, err)
		return
	}

	if fv != nil {
		percentage, err := persistence.GetDiscoveryPercentage(db, account.Username, fv.Category)
		if err != nil {
			log.Fatal("Error retrieving status ", err)
		}

		notification := Notification{
			Notification: fmt.Sprintf("You have Discovered %.2f%% of all fruits in Canada! Keep it going %s!",
				percentage, account.Username),
		}

		passwordHash, err := persistence.GetAccountPasswordHash(db, account.Username)
		if err != nil {
			log.Fatal("Failed to get the password hash ", err)
		}

		SendNotification(notification, passwordHash, "stats")

		topDiscoveryPercentages, err := persistence.GetTopDiscoveryPercentage(db, account.Username, fv.Category)
		if err != nil {
			log.Fatal("Failed to get the percentages per region ", err)
		}

		// Concatenate all percentages into one string
		var stats []string
		for _, stat := range topDiscoveryPercentages {
			stats = append(stats, fmt.Sprintf("\n- %6.2f%% coming from %s", stat.DiscoveryPercentage, stat.RegionName))
		}

		statsString := strings.Join(stats, "")

		// Create the notification message
		notification = Notification{
			Notification: fmt.Sprintf("You have discovered the following fruits and vegetables per region: %s.\nKeep it going %s!",
				statsString, account.Username),
		}

		SendNotification(notification, passwordHash, "all-stats")

	} else {
		log.Printf("Account: %s, No more fruits or vegetables available\n", account.Username)
	}
}
