package game

import (
	"database/sql"
	"fmt"
)

func DisplayModifierEquipe(db *sql.DB) {

	erreurText := "Option non valide !"

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
	idAtk1 := 0
	idAtk2 := 0
	idDef1 := 0
	idDef2 := 0
	idGoal := 0

	// -- Atk 1
	err := db.QueryRow(`
	SELECT j.prenom_joueur, j.nom_joueur, j.id_joueur FROM joueur j
	WHERE id_equipe = 5 and id_poste = 2
	LIMIT 1
	OFFSET 0
	`).Scan(&prenomJoueurAtk1, &nomJoueurAtk1, &idAtk1)
	if err != nil {
		fmt.Println("Erreur:", err)
	}

	// -- Atk 2
	err = db.QueryRow(`
	SELECT j.prenom_joueur, j.nom_joueur, j.id_joueur FROM joueur j
	WHERE id_equipe = 5 and id_poste = 2
	LIMIT 1
	OFFSET 1
	`).Scan(&prenomJoueurAtk2, &nomJoueurAtk2, &idAtk2)
	if err != nil {
		fmt.Println("Erreur2:", err)
	}

	// -- Def 1
	err = db.QueryRow(`
	SELECT j.prenom_joueur, j.nom_joueur, j.id_joueur FROM joueur j
	WHERE id_equipe = 5 and id_poste = 1
	LIMIT 1
	OFFSET 0
	`).Scan(&prenomJoueurDef1, &nomJoueurDef1, &idDef1)
	if err != nil {
		fmt.Println("Erreur3:", err)
	}

	// -- Def 2
	err = db.QueryRow(`
	SELECT j.prenom_joueur, j.nom_joueur, j.id_joueur FROM joueur j
	WHERE id_equipe = 5 and id_poste = 1
	LIMIT 1
	OFFSET 1
	`).Scan(&prenomJoueurDef2, &nomJoueurDef2, &idDef2)
	if err != nil {
		fmt.Println("Erreur4:", err)
	}

	// -- Goal
	err = db.QueryRow(`
	SELECT j.prenom_joueur, j.nom_joueur, j.id_joueur FROM joueur j
	WHERE id_equipe = 5 and id_poste = 0
	LIMIT 1
	`).Scan(&prenomJoueurGoal, &nomJoueurGoal, &idGoal)
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

	valid := false

	for !valid {
		fmt.Println(base)
		fmt.Printf("\n \n 1. Switch 2 joueurs \n 2. Retour \n > ")
		_, err := fmt.Scan(&answer)
		if (answer == 1) || (answer == 2) && (err == nil) {
			valid = true
		} 
		if !valid {fmt.Println("Option non valide !")}
	}


	if answer == 1 {
		fmt.Printf("\nVoici la liste des joueurs : \n")

		fmt.Printf("Attaquant 1 : %s %s | ID: %v \n", prenomJoueurAtk1, nomJoueurAtk1, idAtk1)
		fmt.Printf("Attaquant 2 : %s %s | ID: %v \n", prenomJoueurAtk2, nomJoueurAtk2, idAtk2)
		fmt.Printf("Defenseur 1 : %s %s | ID: %v \n", prenomJoueurDef1, nomJoueurDef1, idDef1)
		fmt.Printf("Defenseur 2 : %s %s | ID: %v \n", prenomJoueurDef2, nomJoueurDef2, idDef2)
		fmt.Printf("Goal : %s %s | ID: %v \n\n", prenomJoueurGoal, nomJoueurGoal, idGoal)

		valid := false
		var id1 int

		for !valid {
			fmt.Println("Entrez l'ID du 1er joueur : ")
			_, err := fmt.Scan(&id1)
			if answer==idAtk1 || answer==idAtk2 || answer==idDef1 || answer==idDef2 || answer==idGoal && err == nil {
				valid = true
			} 
			if !valid {fmt.Println(erreurText)}
		}
		valid = false

		var id2 int
		for !valid {
			fmt.Println("Entrez l'ID du 2eme joueur : ")
			_, err := fmt.Scan(&id2)
			if answer==idAtk1 || answer==idAtk2 || answer==idDef1 || answer==idDef2 || answer==idGoal && err == nil {
				valid = true
			} 
			if !valid {fmt.Println(erreurText)}
		}


		fmt.Printf("Echange : J%v <=> J%v \n", id1, id2)
		if id1 == id2 {
			fmt.Println("Erreur : Vous avez entré le même ID pour les deux joueurs.")
			return
		}
		if id1 != idAtk1 && id1 != idAtk2 && id1 != idDef1 && id1 != idDef2 && id1 != idGoal {
			fmt.Println("Erreur : L'ID du 1er joueur n'est pas valide.")
			return
		}
		if id2 != idAtk1 && id2 != idAtk2 && id2 != idDef1 && id2 != idDef2 && id2 != idGoal {
			fmt.Println("Erreur : L'ID du 2eme joueur n'est pas valide.")
			return
		}

		var posteIdJ1 int
		var vitesseJ1 int
		var enduranceJ1 int
		var forceJ1 int
		var techniqueJ1 int

		var vitesseJ1Exi int
		var enduranceJ1Exi int
		var forceJ1Exi int
		var techniqueJ1Exi int



		var posteIdJ2 int
		var vitesseJ2 int
		var enduranceJ2 int
		var forceJ2 int
		var techniqueJ2 int

		var vitesseJ2Exi int
		var enduranceJ2Exi int
		var forceJ2Exi int
		var techniqueJ2Exi int

		// -- Récupération poste et compétence J1
		err := db.QueryRow(`
		SELECT j.id_poste, vitesse, endurance, force, technique FROM joueur j
		WHERE id_joueur = ?
		`, id1).Scan(&posteIdJ1, &vitesseJ1, &enduranceJ1, &forceJ1, &techniqueJ1)
		if err != nil {
			fmt.Println("Erreur6:", err)
		}

		// -- Récupération poste et compétence J2
		err = db.QueryRow(`
		SELECT j.id_poste, vitesse, endurance, force, technique FROM joueur j
		WHERE id_joueur = ?
		`, id2).Scan(&posteIdJ2, &vitesseJ2, &enduranceJ2, &forceJ2, &techniqueJ2)
		if err != nil {
			fmt.Println("Erreur7:", err)
		}

		// -- Récupération exigences du poste du J1
		err = db.QueryRow(`
		SELECT min_vitesse, min_endurance, min_force, min_technique FROM poste
		WHERE id_poste = ?
		`, posteIdJ1).Scan(&vitesseJ1Exi, &enduranceJ1Exi, &forceJ1Exi, &techniqueJ1Exi)
		if err != nil {
			fmt.Println("Erreur8:", err)
		}

		// -- Récupération exigences du poste du J2
		err = db.QueryRow(`
		SELECT min_vitesse, min_endurance, min_force, min_technique FROM poste
		WHERE id_poste = ?
		`, posteIdJ2).Scan(&vitesseJ2Exi, &enduranceJ2Exi, &forceJ2Exi, &techniqueJ2Exi)
		if err != nil {
			fmt.Println("Erreur9:", err)
		}

		// -- Check si le J1 est capable de jouer avec les exigences du J2 et inversement
		if vitesseJ1 >= vitesseJ2Exi && enduranceJ1 >= enduranceJ2Exi && forceJ1 >= forceJ2Exi && techniqueJ1 >= techniqueJ2Exi {
			if vitesseJ2 >= vitesseJ1Exi && enduranceJ2 >= enduranceJ1Exi && forceJ2 >= forceJ1Exi && techniqueJ2 >= techniqueJ1Exi {
				// -- SI TOUT OK : Faire le switc

				_, err = db.Exec(`
				UPDATE joueur
				SET id_poste = ?
				WHERE id_joueur = ?;
				`, posteIdJ2, id1)
				if err != nil {
					fmt.Println("Erreur10:", err)
				}
				_, err = db.Exec(`
				UPDATE joueur
				SET id_poste = ?
				WHERE id_joueur = ?;
				`, posteIdJ1, id2)
				if err != nil {
					fmt.Println("Erreur11:", err)
				}
				fmt.Println("> Les exigences sont respectés : le changement est en cours...")

			} else {
				fmt.Println("> Le Joueur 2 ne respecte pas les exigences !")
			}
		} else {
			fmt.Println("> Le Joueur 1 ne respecte pas les exigences !")
		}

		// -- SI TOUT OK : Faire le switch

	}

}
