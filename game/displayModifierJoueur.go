package game

import (
    "database/sql"
    "fmt"
)

type Joueur struct {
    ID        int
    Nom       string
    Prenom    string
    IdPoste   int
    Vitesse   int
    Endurance int
    Force     int
    Technique int
    Blesse    bool
}

// Requete de selection pour obtenir les joueurs de l'equipe CESI

func ObtenirJoueursCesi(db *sql.DB) ([]Joueur, error) {
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
        WHERE e.nom_equipe = 'Cesi';
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

// Modification des stats ds joueurs + chercher le joueur dans le slice

func ModifierStatsJoueur(db *sql.DB, joueurs []Joueur, id int)  {
    var joueur *Joueur
    for i := range joueurs {
        if joueurs[i].ID == id {
            joueur = &joueurs[i]
            break
        }
    }
    if joueur == nil {
        fmt.Println("Aucun joueur CESI avec cet ID.")
        return
    }

    for {
        fmt.Printf("\nModification de %s %s (ID %d)\n", joueur.Prenom, joueur.Nom, joueur.ID)
        fmt.Printf("1) Vitesse    : %d\n", joueur.Vitesse)
        fmt.Printf("2) Endurance  : %d\n", joueur.Endurance)
        fmt.Printf("3) Force      : %d\n", joueur.Force)
        fmt.Printf("4) Technique  : %d\n", joueur.Technique)
        fmt.Println("5) Terminer")
        

        var choix int
        valid := false

        for !valid {
            fmt.Print("Choisissez la stat à modifier : ")
            _, err := fmt.Scan(&choix)
            if choix==1 || choix==2 || choix==3 || choix==4 || choix==5 && err != nil {
                valid = true
            } 
            if !valid {fmt.Println("Option non valide !")}
        }
        valid = false

        if choix == 5 {
            break
        }


        var nouvelleValeur int

        for !valid {
            fmt.Print("Nouvelle valeur (0-100) : ")
            _, err := fmt.Scan(&nouvelleValeur)
            if nouvelleValeur>=0 && nouvelleValeur<=100 && err != nil {
                valid = true
            } 
            if !valid {fmt.Println("Option non valide !")}
        }

        
        

        switch choix {
        case 1:
            joueur.Vitesse = nouvelleValeur
        case 2:
            joueur.Endurance = nouvelleValeur
        case 3:
            joueur.Force = nouvelleValeur
        case 4:
            joueur.Technique = nouvelleValeur
        default:
            fmt.Println("Choix invalide.")
            continue
        }

        _, err := db.Exec(`
            UPDATE joueur
            SET vitesse = ?, endurance = ?, force = ?, technique = ?
            WHERE id_joueur = ?;
        `, joueur.Vitesse, joueur.Endurance, joueur.Force, joueur.Technique, joueur.ID)
        if err != nil {
            fmt.Println("Erreur lors de la mise à jour du joueur :", err)
        } else {
            fmt.Println("Stat mise à jour.")
        }
    }
}

// Appel du menu existant + changement de stats des joueurs

func DisplayModifierJoueur(db *sql.DB) int {
    joueurs, err := ObtenirJoueursCesi(db)
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


    // 
    valid := false
    for !valid {
        fmt.Println("Rentrez l'id du joueur que vous voulez modifier :")
        _, err := fmt.Scan(&answer)

        for i := 0; i < len(joueurs); i++ {
			if answer == joueurs[i].ID && err != nil {
				valid = true
				break
			}
		}
        if !valid {fmt.Println("Option non valide !")}
    }

    ModifierStatsJoueur(db, joueurs, answer)

    return answer
}
