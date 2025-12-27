package game

import (
		"fmt"
		"database/sql"
)

func DisplayModifierEquipe(db *sql.DB) int {
	var answer int

	prenomJoueurAtk1 := "J1"
	nomJoueurAtk1 := "J1"
	prenomJoueurAtk2 := "J2"
	nomJoueurAtk2 := ""
	prenomJoueurDef1 := "J3"
	nomJoueurDef1 := ""
	prenomJoueurDef2 := "J4"
	nomJoueurDef2 := ""
	prenomJoueurGoal := "J5"
	nomJoueurGoal := ""

	// -- Atk 1
	err := db.QueryRow(`
	SELECT j.prenom_joueur, j.nom_joueur FROM joueur j
	WHERE id_equipe = 5 and id_poste = 2
	LIMIT 1
	OFFSET 0
	`).Scan(&prenomJoueurAtk1, &nomJoueurAtk1)
	if err != nil {
		fmt.Println("Erreur:", err)
	}

	// -- Atk 2
	err = db.QueryRow(`
	SELECT j.prenom_joueur, j.nom_joueur FROM joueur j
	WHERE id_equipe = 5 and id_poste = 2
	LIMIT 1
	OFFSET 1
	`).Scan(&prenomJoueurAtk2, &nomJoueurAtk2)
	if err != nil {
		fmt.Println("Erreur2:", err)
	}

	// -- Def 1
	err = db.QueryRow(`
	SELECT j.prenom_joueur, j.nom_joueur FROM joueur j
	WHERE id_equipe = 5 and id_poste = 1
	LIMIT 1
	OFFSET 0
	`).Scan(&prenomJoueurDef1, &nomJoueurDef1)
	if err != nil {
		fmt.Println("Erreur3:", err)
	}

	// -- Def 2
	err = db.QueryRow(`
	SELECT j.prenom_joueur, j.nom_joueur FROM joueur j
	WHERE id_equipe = 5 and id_poste = 1
	LIMIT 1
	OFFSET 1
	`).Scan(&prenomJoueurDef2, &nomJoueurDef2)
	if err != nil {
		fmt.Println("Erreur4:", err)
	}

	// -- Goal
	err = db.QueryRow(`
	SELECT j.prenom_joueur, j.nom_joueur FROM joueur j
	WHERE id_equipe = 5 and id_poste = 0
	LIMIT 1
	`).Scan(&prenomJoueurGoal, &nomJoueurGoal)
	if err != nil {
		fmt.Println("Erreur5:", err)
	}
	


  base := `
                           ┌─`+prenomJoueurAtk1+` `+nomJoueurAtk1+`
                           │
                           │                  ┌─`+prenomJoueurDef1+` `+nomJoueurDef1+`
              _____________│__________________│_____________________
             /             ↓           /      ↓                    /
            /_______       o          /       o          _________/
        /| / ___   /      "|\        /       \|/        /  ____/|\
       /_|/    /  /-,      (\  _.--"/"-,     / \   _.--/  o   / |_\
      /       /  /  :         /    /    :         /   /  ~|\ /    / ←-`+prenomJoueurGoal+` `+nomJoueurGoal+`
     /|      /  /  /         (    /    /         (   /  / (\ |\  /  
    /_|     /  /.-'          '._ /_.-'           '. /  /     |_\/
     /-----'  /            o    /         o        /   '-----/
    /--------'         ┌→ "|"  /         \|/ ←┐    '--------/
   /                   │   (\ /          /)   │            /
  /____________________│_____/________________│___________/
                       │                      │
                       │                      └─`+prenomJoueurDef2+` `+nomJoueurDef2+`
                       │
                       └─`+prenomJoueurAtk2+` `+nomJoueurAtk2+`
  `

	fmt.Println(base)
	fmt.Scan(&answer)

	return answer
}
