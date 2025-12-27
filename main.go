package main

import (
	"database/sql"
	"fmt"
	"sportmanager/database"
	"sportmanager/game"

	_ "modernc.org/sqlite"
)

func main() {
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
	case 1: // -- Jouer
		fmt.Println("> Lancement de la partie...")
		answer = game.DisplayJouer()
		switch answer {
			case 1:
				fmt.Println("> Création d'un MATCH")
			case 2:
				fmt.Println("> Gestion de l'équipe")
				answer = game.DisplayConsulterEquipe()
				switch answer {
					case 1:
						fmt.Println("> Modifier l équipe CESI")
						answer = game.DisplayModifierEquipe(db)
					case 2:
						fmt.Println("> Modifier les joueurs de CESI")
						answer = game.DisplayModifierJoueur(db)
				}
			default:
				break
		}
	case 2: // -- Options
		fmt.Println("> Ouverture des options...")
		answer = game.DisplayOptions()
		switch answer {
		case 1:
			fmt.Println("> Ouverture de l'historique...")
		case 2:
			fmt.Println("> Reset de la BDD...")
			err = database.ResetDatabase(db)
			if err != nil {
				fmt.Println("Erreur Reset DB:", err)
			}
		default:
			break
		}

	case 3:
		fmt.Println("> Affichage des règles")
		game.DisplayRules()


	}
}
