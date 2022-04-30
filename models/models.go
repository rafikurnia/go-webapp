package models

import (
	"fmt"
	"os"

	"github.com/rafikurnia/go-webapp/utils"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// A variable placeholders for models struct that will be accessible from apipackage
var DB *models

// Provides access to the database and available data, to hide implementation details from api
// package, so that it becomes loosely coupled
type models struct {
	Contacts contactLinker
}

// Initialize tables, either a mock database or an actual database
// It accepts one argument to connect with DB object
// It returns any error encountered
func initTables(db *models) error {
	if err := db.Contacts.Init(); err != nil {
		return fmt.Errorf("Contacts initTables(db *models) -> %w", err)
	}

	return nil
}

// Initialize database, either a mock database or an actual database
// It accepts one argument to configure DB connection mode, i.e., mock or not
// It returns any error encountered
func InitDB(mode string) error {
	if mode == utils.DBModeMock {
		DB = &models{
			Contacts: &mockContactModel{},
		}
	} else {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
			os.Getenv("POSTGRES_HOST"),
			os.Getenv("POSTGRES_PASSWORD"),
			os.Getenv("POSTGRES_USER"),
			os.Getenv("POSTGRES_DB"),
			os.Getenv("POSTGRES_PORT"),
			os.Getenv("POSTGRES_SSL"),
			os.Getenv("POSTGRES_TIMEZONE"),
		)

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return fmt.Errorf("OpenDB InitDB(mode string) -> %w", err)
		}

		DB = &models{
			Contacts: &contactModel{DB: db},
		}

		if mode == utils.DBModeTest || os.Getenv("SEED_DATA") == "true" {
			db.Exec("DROP TABLE contacts")
		}
	}

	if err := initTables(DB); err != nil {
		return fmt.Errorf("InitDB(mode string) -> %w", err)
	}

	if mode == utils.DBModeTest || os.Getenv("SEED_DATA") == "true" {
		for _, c := range contactSeedData {
			c.ID = 0
			DB.Contacts.Create(&c)
		}
	}

	return nil
}
