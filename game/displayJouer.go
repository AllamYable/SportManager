package game

import (
		"fmt"
)

func DisplayJouer() int {
	var answer int
	optionJouer := `
	+----------------------------------------+
	|                  MENU                  |
	+----------------------------------------+
	|    Jouer   ══════╗                     |
	|    Option        ▼                     |
	|    À Propos   1. Créer un match !      |
	|    Sortir     2. Consulter son équipe  |
	|               3. Retour au menu        |
	|                                        |
	+----------------------------------------+
	`
	
	valid := false

	for !valid {
		fmt.Println(optionJouer)
		_, err := fmt.Scan(&answer)
		if (answer == 1) || (answer == 2) || (answer == 3) && (err == nil) {
			valid = true
		} 
		if !valid {fmt.Println("Option non valide !")}
	}

	return answer
}