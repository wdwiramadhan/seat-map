package main

import (
	"database/sql"
	"flag"
	"log"
	"os"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	databaseURL := os.Getenv("DATABASE_URL")

	action := flag.String("action", "migrate", "action: migrate | seed")
	flag.Parse()

	if *action == "migrate" {
		runMigrations(databaseURL)
	} else if *action == "seed" {
		runSeeds(databaseURL)
	}

}

func runMigrations(databaseURL string) {
	m, err := migrate.New(
		"file://db/migrations",
		databaseURL,
	)

	if err != nil {
		log.Fatal("Could not connect to the database:", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to apply migrations: %v", err)
	}

	log.Println("Migrations applied successfully!")
}

func runSeeds(databaseURL string) {
	db, err := sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal("Could not connect to the database for seeding:", err)
	}

	defer db.Close()

	files, err := filepath.Glob("db/seeds/*.sql")
	if err != nil {
		log.Fatal("Failed to find seed files:", err)
	}

	if len(files) == 0 {
		log.Println("No seed files found")
		return
	}

	for _, file := range files {
		log.Printf("Applying seed: %s", file)

		content, err := os.ReadFile(file)
		if err != nil {
			log.Fatalf("Failed to read %s: %v", file, err)
		}

		if _, err := db.Exec(string(content)); err != nil {
			log.Fatalf("Failed to apply seed %s: %v", file, err)
		}

		log.Printf("Seed %s applied successfully", file)
	}

	log.Println("All seeds applied successfully!")
}
