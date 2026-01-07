package game

import (
		"fmt"
)

func DisplayMenu() int {
	var answer int
	menu := `
    +--------------------------+
    |           MENU           |
    +--------------------------+
    | 1. Jouer                 |
    | 2. Options               |
    | 3. À Propos              |
    | 4. Sortir                |
    +--------------------------+
    `

	valid := false

	for !valid {
		fmt.Println(menu)
		_, err := fmt.Scan(&answer)
		if (answer == 1) || (answer == 2) || (answer == 3) || (answer == 4) && (err == nil) {
			valid = true
		} 
		if !valid {fmt.Println("Option non valide !")}
	}

	return answer
}

