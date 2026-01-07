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
	|    Sortir     1. Reset BDD             |
	|               2. Retour au menu        |
	+----------------------------------------+
	`

	valid := false

	for !valid {
		fmt.Println(optionMenu)
		_, err := fmt.Scan(&answer)
		if (answer == 1) || (answer == 2) && (err == nil) {
			valid = true
		} 
		if !valid {fmt.Println("Option non valide !")}
	}

	return answer
}