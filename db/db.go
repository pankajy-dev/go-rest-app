package db

import (
	"database/sql"

	// _ tells that this import is needed in final build, don't auto remove it, the below dependecy is needed for database/sql
	_ "github.com/mattn/go-sqlite3"
)

// The DB will be used in other files to interact with database.
var DB *sql.DB

func InitDb() {
	// Open expects two params:
	// 1. driver name in our case it is : sqlite3
	// 2. datasource name, if it doesn't exist then it will be auto created
	DB, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Error during Database object creation.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
}
