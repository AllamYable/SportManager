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

	answer := game.DisplayMenu()

	switch answer {
		case 1 :
			fmt.Println("> Lancement de la partie...")
		case 2 :
			fmt.Println("> Ouverture des options...")
			answer = game.DisplayOptions()
			switch answer {
				case 1 :
					fmt.Println("> Ouverture de l'historique...")
				case 2 :
					fmt.Println("> Reset de la BDD...")
					err = database.ResetDatabase(db) 
					if err != nil{ fmt.Println("Erreur Reset DB:",err) }
				}
				default :
					break

		case 3 :
			fmt.Println("> Affichage des règles")
	}
}
