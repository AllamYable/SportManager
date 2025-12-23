package game

import (
		"fmt"
)

func DisplayOptions() int {
	var answer int
	optionMenu := `
	+----------------------------------------+
	|                  MENU                  |
	+----------------------------------------+
	|    Jouer                               |
	|    Options ══════╗                     |
	|    À Propos      ▼                     |
	|    Sortir     1. Historique            |
	|               2. Réinitialiser la BDD  |
	|               3. Retour au menu        |
	+----------------------------------------+
	`

	fmt.Println(optionMenu)
	fmt.Scan(&answer)

	return answer
}