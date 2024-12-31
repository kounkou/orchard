package notifier

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"orchard/pkg/persistence"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Notification struct {
	Notification string `json:"notification"`
	ImageURL 	 string `json:"imageURL"`
}

type FruitVegetable struct {
	Name     string
	Category string
	ImageURL string
}

type Suggestions struct {
	Items []persistence.FruitVegetable `json:"items"`
}

func generateRandomNotification(name, description string) Notification {
	notifications := []string{
		"Discover the %s! %s Learn more about the %s in Orchard!",
		"Have you tried the %s yet? %s Dive into the details about %s in Orchard!",
		"The %s awaits you! %s Check out all there is to know about %s in Orchard!",
		"Say hello to the %s! %s Uncover more about %s in Orchard!",
		"A new discovery: %s! %s Get the full scoop on %s in Orchard!",
		"Did you know about the %s? %s Explore all the details about %s in Orchard!",
		"Introducing the %s! %s Learn everything you need to know about %s in Orchard!",
		"Time to discover the %s! %s Find out more about %s in Orchard!",
		"Get to know the %s! %s Discover all there is to learn about %s in Orchard!",
		"Curious about the %s? %s Delve into the world of %s in Orchard!",
		"Unearth the wonders of the %s! %s Learn more about %s in Orchard!",
	}

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	randomTemplate := notifications[rng.Intn(len(notifications))]

	formattedNotification := fmt.Sprintf(randomTemplate, name, description, name)

	return Notification{
		Notification: formattedNotification,
		ImageURL: "", // TODO fill this image 
	}
}

func SendNotification(notification Notification, passwordHash string) {
	dataDir := "."
	if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
		log.Fatalf("Error creating data directory: %v", err)
	}

	filePath := filepath.Join(dataDir, passwordHash+"-notification.json")

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

		notification := generateRandomNotification(fv.Name, description)

		passwordHash, err := persistence.GetAccountPasswordHash(db, account.Username)
		if err != nil {
			log.Fatal("Failed to get the password hash ", err)
		}

		SendNotification(notification, passwordHash)
	} else {
		log.Printf("Account: %s, You have discovered all fruits and vegetables!\n", account.Username)
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

		SendNotification(notification, passwordHash)
	} else {
		log.Printf("Account: %s, You have discoverd all possible fruits and vegetables!\n", account.Username)
	}
}

func CreateStats(db *sql.DB, accountHash string) {
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
		topDiscoveryPercentages, err := persistence.GetTopDiscoveryPercentage(db, account.Username, fv.Category)
		if err != nil {
			log.Fatal("Failed to get the percentages per region ", err)
		}

		var stats []string
		for _, stat := range topDiscoveryPercentages {
			stats = append(stats, fmt.Sprintf("\n- %6.2f%% coming from %s", stat.DiscoveryPercentage, stat.RegionName))
		}

		statsString := strings.Join(stats, "")

		var notification Notification

		notification = Notification{
			Notification: fmt.Sprintf("You have discovered the following fruits and vegetables per region: %s.\nKeep it going %s!",
				statsString, account.Username),
		}

		passwordHash, err := persistence.GetAccountPasswordHash(db, account.Username)
		if err != nil {
			log.Fatal("Failed to get the password hash ", err)
		}

		SendNotification(notification, passwordHash)

	} else {
		log.Printf("Account: %s, You have discoverd all possible fruits and vegetables!\n", account.Username)
	}
}

func CreateItemsSuggestionNotification(items []persistence.FruitVegetable) (Suggestions, error) {
	if len(items) == 0 {
		return Suggestions{}, fmt.Errorf("no items provided to create suggestions")
	}

	suggestions := Suggestions{
		Items: items,
	}

	return suggestions, nil
}