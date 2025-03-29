package storage

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("Error al abrir la base de datos:", err)
	}

	statement := `
	CREATE TABLE IF NOT EXISTS matches (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		homeTeam TEXT,
		awayTeam TEXT,
		matchDate TEXT,
		goals INTEGER DEFAULT 0,
		yellowCards INTEGER DEFAULT 0,
		redCards INTEGER DEFAULT 0,
		extraTime INTEGER DEFAULT 0
	);
	`
	_, err = DB.Exec(statement)
	if err != nil {
		log.Fatal("Error al crear la tabla:", err)
	}
}
