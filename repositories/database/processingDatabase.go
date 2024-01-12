package database

import (
	"database/sql"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
)

var (
	db   *sql.DB
	once sync.Once
)

func ProcessingCreateTablesIfNotExists() {
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS processing (id INTEGER PRIMARY KEY AUTOINCREMENT, process DOUBLE, free_memory DOUBLE, swap DOUBLE)")

	if err != nil {
		log.Fatal(err)
	}

	statement.Exec()
}

func ProcessingInitDBConnection(path string) (*sql.DB, error) {
	once.Do(func() {
		var err error

		db, err = sql.Open("sqlite3", path)

		if err != nil {
			log.Fatal(err)
		}

		err = db.Ping()

		if err != nil {
			log.Fatal(err)
		}

		ProcessingCreateTablesIfNotExists()
	})

	return db, nil
}

func ProcessingGetDBConnection() *sql.DB {
	return db
}

func ProcessingCloseDBConnection() {
	db.Close()
}
