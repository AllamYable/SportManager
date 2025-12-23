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

	fmt.Println(menu)
	fmt.Scan(&answer)

	return answer
}

