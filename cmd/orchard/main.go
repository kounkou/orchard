package main

import (
	"database/sql"
	"log"
	"orchard/internal/config"
	"orchard/pkg/migration"
	"orchard/pkg/server"

	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	cfg := config.LoadConfig()
	db, err := sql.Open("sqlite3", cfg.DatabaseName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = migration.Migrate(db)
	if err != nil {
		log.Fatal(err)
	}

	go server.Start(db)

	select {}
}
