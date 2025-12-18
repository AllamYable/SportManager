package main

import (
	"database/sql"
	"fmt"
	"sportmanager/database"
	"sportmanager/game"

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

	err = database.InitDatabase(db)
	if err != nil{
		fmt.Println("Erreur Init DB:",err)
	}
	database.PushDatabase(db)

	answer := game.DisplayMenu()

	switch answer {
		case 1 :
			fmt.Println("Lancement de la partie...")
		case 2 :
			fmt.Println("Ouverture des options...")
		case 3 :
			fmt.Println("Affichage des règles")
	}
}
