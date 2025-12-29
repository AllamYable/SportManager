package game

import (
    "database/sql"
    "fmt"
)

// --- Menu ASCII + déclenchement de la création de match ---
func DisplayCreerMatch(db *sql.DB) int {
    var answer int
    optionJouer := `
    +-------------------------------------------------------------+
    |                            MENU                             |
    +-------------------------------------------------------------+
    |    Jouer   ══════╗                                          |
    |    Option        ▼                                          |
    |    À Propos      Créer un match !                           |
    |    Sortir        Consulter son équipe ══════╗                |
    |                   Retour au menu           ▼                |
    |                                          1. Créer un match  |
    |                                          2. Retour jouer    |
    +-------------------------------------------------------------+
    `

    fmt.Println(optionJouer)
    fmt.Scan(&answer)

    if answer == 1 {
        DisplayCreationMatch(db)
    }

    return answer
}

// --- Structures et fonctions BDD pour la création du match ---

type Equipe struct {
    ID  int
    Nom string
}

func ObtenirEquipesAdverses(db *sql.DB) ([]Equipe, error) {
    rows, err := db.Query(`
        SELECT id_equipe, nom_equipe
        FROM equipe
        WHERE nom_equipe <> 'Cesi'
		ORDER BY id_equipe ASC;
    `)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var equipes []Equipe
    for rows.Next() {
        var e Equipe
        if err := rows.Scan(&e.ID, &e.Nom); err != nil {
            return nil, err
        }
        equipes = append(equipes, e)
    }

    return equipes, nil
}

func ObtenirEquipeCesiID(db *sql.DB) (int, error) {
    var id int
    err := db.QueryRow(`
        SELECT id_equipe
        FROM equipe
        WHERE nom_equipe = 'Cesi';
    `).Scan(&id)
    if err != nil {
        return 0, err
    }
    return id, nil
}

func DisplayCreationMatch(db *sql.DB) {
    equipes, err := ObtenirEquipesAdverses(db)
    if err != nil {
        fmt.Println("Erreur lors de la récupération des équipes adverses :", err)
        return
    }
    if len(equipes) == 0 {
        fmt.Println("Aucune équipe adverse disponible.")
        return
    }

    fmt.Println("=== Choisissez l'équipe adverse ===")
    for _, e := range equipes {
        fmt.Printf("ID: %d | %s\n", e.ID, e.Nom)
    }

    var idAdverse int
    fmt.Print("Entrez l'ID de l'équipe adverse : ")
    fmt.Scan(&idAdverse)

    existe := false
    for _, e := range equipes {
        if e.ID == idAdverse {
            existe = true
            break
        }
    }
    if !existe {
        fmt.Println("ID d'équipe adverse invalide.")
        return
    }

    idCesi, err := ObtenirEquipeCesiID(db)
    if err != nil {
        fmt.Println("Impossible de trouver l'équipe CESI :", err)
        return
    }

    _, err = db.Exec(`
        INSERT INTO match (date_match, id_equipe_domicile, id_equipe_exterieur, score_domicile, score_exterieur, gagnant)
        VALUES (CURRENT_DATE, ?, ?, 0, 0, NULL);
    `, idCesi, idAdverse)
    if err != nil {
        fmt.Println("Erreur lors de la création du match :", err)
        return
    }

    fmt.Println("Match créé entre CESI et l'équipe choisie.")
}
