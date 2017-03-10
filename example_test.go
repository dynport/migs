package migs

import (
	"database/sql"
	"log"
)

func runMigrations(tx *sql.Tx) {
	migs := New(
		"CREATE TABLE users (id SERIAL NOT NULL PRIMARY KEY, email VARCHAR)",
		migrationFunc, // allow executing functions as well
	)
	err := migs.ExecuteTx(tx)
	if err != nil {
		log.Fatal(err)
	}
}

func migrationFunc(tx Con) error {
	_, err := tx.Exec("INSERT INTO users (email) VALUES ($1)", "tobias@phraseapp.com")
	return err
}
