package game

import (
		"fmt"
)

func DisplayConsulterEquipe() int {
	var answer int
	optionJouer := `
	+-------------------------------------------------------------+
	|                            MENU                             |
	+-------------------------------------------------------------+
	|    Jouer   ══════╗                                          |
	|    Option        ▼                                          |
	|    À Propos      Créer un match !                           |
	|    Sortir        Consulter son équipe ══════╗               |
	|                  Retour au menu             ▼               |
	|                                          1. Modifier équipe |
	|                                          2. Modifier joueur |
	|                                          3. Retour jouer    |
	+-------------------------------------------------------------+
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
