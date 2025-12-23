package database

import (
	"database/sql"
	"fmt"
)

const FK_ON = "PRAGMA foreign_keys = ON;"

func ResetDatabase(conn *sql.DB) error{

	if _, err := conn.Exec("PRAGMA foreign_keys = OFF;"); err != nil {
		return err
	}

    conn.Exec("DROP TABLE IF EXISTS equipe;")
    conn.Exec("DROP TABLE IF EXISTS poste;")
    conn.Exec("DROP TABLE IF EXISTS joueur;")
    conn.Exec("DROP TABLE IF EXISTS match;")
    conn.Exec("DROP TABLE IF EXISTS performance;")
    conn.Exec("DROP TABLE IF EXISTS dispute;")
    conn.Exec("DROP TABLE IF EXISTS blessure;")

	if _, err := conn.Exec(FK_ON); err != nil {
		return err
	}

	// -- Désactiver temporairement les vérifications de clés étrangères
	if _, err := conn.Exec("PRAGMA foreign_keys = OFF;"); err != nil {
		return err
	}

	// -- Réinsérer les données initiales
	err := InitDatabase(conn) 
	if err != nil{ fmt.Println("Erreur Init DB:",err) }
	PushDatabase(conn)

	// -- Réactiver les vérifications de clés étrangères
	if _, err := conn.Exec(FK_ON); err != nil {
		return err
	}

	fmt.Println("> Base Données réinitialisée avec succès.")
	return nil

}
