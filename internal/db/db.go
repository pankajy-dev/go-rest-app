package db

import (
	"database/sql"

	"github.com/pankajy-dev/go-rest-app/internal/config"

	// _ tells that this import is needed in final build, don't auto remove it, the below dependecy is needed for database/sql
	_ "github.com/mattn/go-sqlite3"
)

// The DB will be used in other files to interact with database.
var DB *sql.DB

func InitDb() {
	// Open expects two params:
	// 1. driver name in our case it is : sqlite3
	// 2. datasource name, if it doesn't exist then it will be auto created
	var err error
	DB, err = sql.Open("sqlite3", config.DBPath())

	if err != nil {
		panic("Error during Database object creation.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {

	createUserTable := `
		CREATE TABLE IF NOT EXISTS user(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`

	_, err := DB.Exec(createUserTable)

	if err != nil {
		panic("Could not create user table.")
	}

	createEventTable := `
		CREATE TABLE IF NOT EXISTS event(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			dateTime DATETIME NOT NULL,
			user_id INTEGER NOT NULL,
			FOREIGN KEY (user_id) REFERENCES  user (id)
		)
	`
	_, err = DB.Exec(createEventTable)

	if err != nil {
		panic("Could not create event table.")
	}

	createRegistrationTable := `
		CREATE TABLE IF NOT EXISTS registration(
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			event_id INTEGER NOT NULL,
			user_id INTEGER NOT NULL,
			FOREIGN KEY (user_id) REFERENCES  user (id),
			FOREIGN KEY (event_id) REFERENCES event (id)
		)
	`
	_, err = DB.Exec(createRegistrationTable)

	if err != nil {
		panic("Could not create registration table.")
	}

}
