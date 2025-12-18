package dbPack

import (
	"database/sql"
	"fmt"
)

func PushPoste(db *sql.DB) {
	fmt.Println("pushing posts...")
	_, err := db.Exec(`
        INSERT INTO poste (id_poste, nom_poste, description_poste, min_vitesse, min_endurance, min_force, min_technique) VALUES
        (0, 'Gardien', 'Le but du gardien est d empecher que la balle rentre dans son but', 5, 5, 15, 5),
        (1, 'Defenseur', 'Le but du defenseur est d empecher les attaquants adverses de s approcher du but', 0, 0, 70, 5),
        (2, 'Attaquant', 'Le but d un attaquant est de marquer des buts pour augmenter le score de l équipe', 0, 0, 5, 70);
    `)
	if err != nil {
		fmt.Println("Error in pushPoste.go :", err)
	}
}
