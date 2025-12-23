package game

import (
        "fmt"
        "database/sql"
)

type Joueur struct {
    ID         int
    Nom        string
    Prenom     string
    IdPoste    int
    Vitesse    int
    Endurance  int
    Force      int
    Technique  int
    Blesse     bool
}

//Fonction pour récupérer les joueurs de l'equipe CESI

func GetJoueursCesi(db *sql.DB) ([]Joueur, error) {
	rows, err := db.Query(`
		SELECT
		j.id_joueur,
       j.nom_joueur,
       j.prenom_joueur,
       j.id_poste,
       j.vitesse,
       j.endurance,
       j.force,
       j.technique,
       j.blesse
	FROM joueur j
	JOIN equipe e ON j.id_equipe = e.id_equipe
	WHERE (e.nom_equipe) = 'Cesi';
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var joueurs []Joueur

	for rows.Next() {
		var j Joueur
		if err := rows.Scan(
			&j.ID,
			&j.Nom,
			&j.Prenom,
			&j.IdPoste,
			&j.Vitesse,
			&j.Endurance,
			&j.Force,
			&j.Technique,
			&j.Blesse,
		); err != nil {
			return nil, err
		}
		joueurs = append(joueurs, j)
	}

	return joueurs, nil
}

//Affichage de l'equipe CESI
func DisplayModifierJoueur(db *sql.DB) int {
    joueurs, err := GetJoueursCesi(db)
    if err != nil {
        fmt.Println("Erreur lors de la récupération des joueurs CESI :", err)
    }

    fmt.Println("=== Joueurs de l'équipe CESI ===")
	fmt.Println(joueurs)
    for _, j := range joueurs {
        fmt.Printf("ID: %d | %s %s | Poste: %d | VIT:%d END:%d FOR:%d TEC:%d | Blesse: %t\n",
            j.ID, j.Nom, j.Prenom, j.IdPoste, j.Vitesse, j.Endurance, j.Force, j.Technique, j.Blesse)
    }

	var answer int
	fmt.Println("Rentrez l id du joueur que vous voulez modifier")
	fmt.Scan(&answer)

	return answer
}


