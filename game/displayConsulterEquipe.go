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

	fmt.Println(optionJouer)
	fmt.Scan(&answer)

	return answer
}
