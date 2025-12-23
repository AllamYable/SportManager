package game

import (
		"fmt"
)

func DisplayRules() int {
	var answer int
	regles := `
    ____________________________
  / \                           \ .
 |   |      RÈGLES DU JEU:      |.
  \_ |                          |.
     | - Vous êtes le coach de  |.
     | l'équipe CESI en foot    |.
     | 5 contre 5 virtuel.      |.
     |                          |.
     | - Après chaque match,    |.
     | saisissez le score et    |.
     | les stats de vos joueurs.|. 
     |                          |.
     | - Vous gagnez 20 pts de  |.
     | compétence (5 par        |.
     | joueur) à répartir.      |. 
     |                          |.
     | - Le niveau de l'équipe  |.
     | évolue selon les         |.
     | résultats et l'adversaire|.
     | pour viser le top.       |. 
     |                          |.
     |           ~ GLHF :D ~    |.
     |   _______________________|___
     |  /                          /.
     \_/__________________________/.
`

	fmt.Println(regles)
	fmt.Scan(&answer)

	return answer
}