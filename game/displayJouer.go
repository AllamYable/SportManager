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

	fmt.Println(optionJouer)
	fmt.Scan(&answer)

	return answer
}