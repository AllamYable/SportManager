package game

import (
    "database/sql"
    "fmt"
    "math"
)

// --- Menu ASCII + déclenchement de la création de match ---
func DisplayCreerMatch(db *sql.DB) {
    var answer int
    optionJouer := `
    +-------------------------------------------------------------+
    |                            MENU                             |
    +-------------------------------------------------------------+
    |    Jouer   ══════╗                                          |
    |    Option        ▼                                          |
    |    À Propos      Créer un match !  ════════════════╗        |
    |    Sortir        Consulter son équipe              ▼        |
    |                  Retour au menu          1. Créer un match  |
    |                                          2. Retour jouer    |
    +-------------------------------------------------------------+
    `

    fmt.Println(optionJouer)
    fmt.Scan(&answer)

    if answer == 1 {
        DisplayCreationMatch(db)
    }
}

// Structure et fonctions BDD pour la création du match

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

// Création du match puis saisie du score
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

    // Création du match avec scores de base à 0
    res, err := db.Exec(`
        INSERT INTO match (date_match, id_equipe_domicile, id_equipe_exterieur, score_domicile, score_exterieur, gagnant)
        VALUES (CURRENT_DATE, ?, ?, 0, 0, NULL);
    `, idCesi, idAdverse)
    if err != nil {
        fmt.Println("Erreur lors de la création du match :", err)
        return
    }

    matchID, err := res.LastInsertId()
    if err != nil {
        fmt.Println("Impossible de récupérer l'ID du match :", err)
        return
    }

    fmt.Println("Match créé entre CESI et l'équipe choisie.")

    // Saisie et enregistrement du score + mise à jour compteurs + points joueurs
    SaisirScoreMatch(db, int(matchID), idCesi, idAdverse)
    randomBlessure(db, int(matchID), idCesi, idAdverse)
}

// Saisie du score : CESI à gauche, adverse à droite, puis UPDATE du match
func SaisirScoreMatch(db *sql.DB, matchID int, idCesi int, idAdverse int) {
    var scoreCesi, scoreAdv int

    fmt.Printf("\nEntrez le score sous la forme CESI - Adversaire\n")
    fmt.Print("Score CESI : ")
    fmt.Scan(&scoreCesi)
    fmt.Print("Score équipe adverse : ")
    fmt.Scan(&scoreAdv)

    // Déterminer le gagnant (id équipe) ou NULL en cas de nul
    var gagnant interface{}
    if scoreCesi > scoreAdv {
        gagnant = idCesi
    } else if scoreAdv > scoreCesi {
        gagnant = idAdverse
    } else {
        gagnant = nil
    }

    _, err := db.Exec(`
        UPDATE match
        SET score_domicile = ?, score_exterieur = ?, gagnant = ?
        WHERE id_match = ?;
    `, scoreCesi, scoreAdv, gagnant, matchID)
    if err != nil {
        fmt.Println("Erreur lors de la mise à jour du score du match :", err)
        return
    }

    fmt.Printf("Score enregistré : %d - %d\n", scoreCesi, scoreAdv)

    // Mise à jour des compteurs d'équipe CESI
    MettreAJourCompteursEquipe(db, idCesi, scoreCesi, scoreAdv)
    // Mise à jour des compteurs de l'équipe adverse
    MettreAJourCompteursEquipe(db, idAdverse, scoreAdv, scoreCesi)

    // Attribution des points de compétence aux joueurs de CESI
    DistribuerPointsJoueurs(db, matchID, idCesi)

    // -- Modif niveau global de l'equipe CESI (1) et l'autre (2) : 
    eloUpdate(db, idCesi, idAdverse, scoreCesi, scoreAdv)
}

// Mise à jour des compteurs d'une équipe : nb_matchs, nb_victoires, nb_defaites
func MettreAJourCompteursEquipe(db *sql.DB, idEquipe int, scoreEquipe int, scoreAdverse int) {
    // Incrémenter nb_matchs toujours
    _, err := db.Exec(`
        UPDATE equipe
        SET nb_matchs = nb_matchs + 1
        WHERE id_equipe = ?;
    `, idEquipe)
    if err != nil {
        fmt.Println("Erreur lors de la mise à jour de nb_matchs :", err)
        return
    }

    // Incrémenter victoires ou défaites selon le résultat
    if scoreEquipe > scoreAdverse {
        _, err = db.Exec(`
            UPDATE equipe
            SET nb_victoires = nb_victoires + 1
            WHERE id_equipe = ?;
        `, idEquipe)
    } else if scoreEquipe < scoreAdverse {
        _, err = db.Exec(`
            UPDATE equipe
            SET nb_defaites = nb_defaites + 1
            WHERE id_equipe = ?;
        `, idEquipe)
    }
    // Si égalité, on ne touche ni victoires ni défaites

    if err != nil {
        fmt.Println("Erreur lors de la mise à jour des compteurs de victoire/défaite :", err)
    }
}

// Distribution de 20 points (max 5 par stat par joueur) pour les joueurs de CESI
func DistribuerPointsJoueurs(db *sql.DB, matchID int, idCesi int) {
    // Récupérer les joueurs de CESI
    joueurs, err := ObtenirJoueursCesi(db)
    if err != nil {
        fmt.Println("Erreur lors de la récupération des joueurs CESI :", err)
        return
    }

    fmt.Println("\n=== Attribution des points de compétence ===")
    fmt.Println("Vous disposez de 20 points à distribuer (5 max par stat par joueur)")

    pointsRestants := 20

    for i := 0; i < len(joueurs) && pointsRestants > 0; i++ {
        j := &joueurs[i]
        fmt.Printf("\n--- %s %s (VIT:%d END:%d FOR:%d TEC:%d) ---\n", j.Prenom, j.Nom, j.Vitesse, j.Endurance, j.Force, j.Technique)
        fmt.Printf("Points restants à distribuer : %d\n", pointsRestants)

        // Distribution par stat avec limite de 5 par stat
        pointsVitesse := demanderPoints("Vitesse", pointsRestants)
        pointsRestants -= pointsVitesse
        j.Vitesse += pointsVitesse

        if pointsRestants > 0 {
            pointsEndurance := demanderPoints("Endurance", pointsRestants)
            pointsRestants -= pointsEndurance
            j.Endurance += pointsEndurance
        }

        if pointsRestants > 0 {
            pointsForce := demanderPoints("Force", pointsRestants)
            pointsRestants -= pointsForce
            j.Force += pointsForce
        }

        if pointsRestants > 0 {
            pointsTechnique := demanderPoints("Technique", pointsRestants)
            pointsRestants -= pointsTechnique
            j.Technique += pointsTechnique
        }

        // Mise à jour en BDD
        _, err := db.Exec(`
            UPDATE joueur
            SET vitesse = ?, endurance = ?, force = ?, technique = ?
            WHERE id_joueur = ?;
        `, j.Vitesse, j.Endurance, j.Force, j.Technique, j.ID)
        if err != nil {
            fmt.Println("Erreur lors de la mise à jour du joueur :", err)
        }
    }

    fmt.Println("\nTous les points ont été distribués.")
}

// Demande d'attribution de points pour une stat donnée (max 5)
func demanderPoints(nomStat string, pointsDisponibles int) int {
    var points int
    maxPoints := 5
    if pointsDisponibles < maxPoints {
        maxPoints = pointsDisponibles
    }

    for {
        fmt.Printf("Points à ajouter à %s (0-%d) : ", nomStat, maxPoints)
        fmt.Scan(&points)

        if points >= 0 && points <= maxPoints {
            return points
        }
        fmt.Printf("Erreur : vous devez entrer entre 0 et %d points.\n", maxPoints)
    }
}

// Insérer le format aléatoire de blessure

func eloUpdate(db *sql.DB, idCesi int, idAdverse int, scoreCesi int, scoreAdv int) {
    // recup score global : 

    var levelCESI float64
    var levelADV float64

    var resultatEloCesi float64
    var resultatEloAdv float64
    var probaEquipeCesiWin float64
    var probaEquipeAdvWin float64
    poids := 30.0

    var resultatReelCesi float64
    var resultatReelAdv float64


    err := db.QueryRow(`
        SELECT niveau_global
        FROM equipe
        WHERE id_equipe = ?;
    `, idCesi).Scan(&levelCESI)
    if err != nil {
            fmt.Println("Erreur lors de la récupération niveau equipe CESI :", err)
    }
    err = db.QueryRow(`
        SELECT niveau_global
        FROM equipe
        WHERE id_equipe = ?;
    `, idAdverse ).Scan(&levelADV)
    if err != nil {
            fmt.Println("Erreur lors de la récupération niveau equipe ADV :", err)
    }


    // -- Proba que CESI gagne : 

    var base float64
    base = 10
    exposant := (levelADV-levelCESI)/400
    exposant64 := float64(exposant)
    result := math.Pow(base, exposant64)

    probaEquipeCesiWin = 1/(1+result)

    // -- Proba que ADV gagne : 

    probaEquipeAdvWin = 1-probaEquipeCesiWin

    fmt.Println("Proba CESI win : ", probaEquipeCesiWin, " vs ", probaEquipeAdvWin)

    // -- Resultat Float : 
    if scoreCesi > scoreAdv {
        resultatReelCesi = 1
    } else if scoreCesi < scoreAdv {
        resultatReelCesi = 0
    } else {
        resultatReelCesi = 0.5
    }

    resultatReelAdv = 1 - resultatReelCesi

    // -- MAJ rating :
    resultatEloCesi = levelCESI + (poids * (resultatReelCesi-probaEquipeCesiWin))
    resultatEloAdv = levelADV + (poids * (resultatReelAdv-probaEquipeAdvWin))

    fmt.Println("Elo CESI avant : ", levelCESI, " après : ", resultatEloCesi)
    fmt.Println("Elo ADV avant : ", levelADV, " après : ", resultatEloAdv)

    var eloCesiFinal int
    var eloAdvFinal int

    eloCesiFinal = int(math.Round(resultatEloCesi))
    eloAdvFinal = int(math.Round(resultatEloAdv))

    // -- push sur la BDD : 

    _, err = db.Exec(`
        UPDATE equipe
        SET niveau_global = ?
        WHERE id_equipe = ?;
    `, eloCesiFinal, idCesi)
    if err != nil {
        fmt.Println("Erreur lors de la mise à jour du ELO equipe CESI :", err)
        return
    }
    _, err = db.Exec(`
        UPDATE equipe
        SET niveau_global = ?
        WHERE id_equipe = ?;
    `, eloAdvFinal, idAdverse)
    if err != nil {
        fmt.Println("Erreur lors de la mise à jour du ELO equipe adverse :", err)
        return
    }
}   

func randomBlessure(db *sql.DB, matchID int, idCesi int, idAdverse int) {
    fmt.Println("Aucune Blessure durant ce match.")
}