package main

import (
	"sportmanager/database"
	"fmt"
	"database/sql"
	_ "modernc.org/sqlite"
)

func main () {
	dbPath := "./database/sqlite.db"

	// -- Ouvrir la DB
	db, err := sql.Open("sqlite", dbPath)
	if err != nil {
		fmt.Println("Erreur ouverture DB:", err)
		return
	}
	defer db.Close() 

	database.CreerSchema(db)
}