package dbPack

import (
	"database/sql"
	"fmt"
)

func PushJoueur(db *sql.DB) {
	fmt.Println("pushing players...")
	_, err := db.Exec(`
        INSERT INTO joueur (nom_joueur, prenom_joueur, age, id_poste, vitesse, endurance, force, technique, blesse, date_blessure, matchs_absence, id_equipe) VALUES
        ('Ine', 'Tart', 20, 2, 60, 60, 30, 70, false, null, 0, 1),
		('Bon', 'Jean', 36, 1, 30, 70, 70, 20, false, null, 0, 1),
		('Vire', 'Lenna', 39, 2, 60, 60, 30, 70, false, null, 0, 1),
		('Auto', 'Riz', 25, 0, 5, 5, 20, 5, false, null, 0, 1),
		('Hillage', 'Mac', 19, 1, 30, 70, 70, 20, false, null, 0, 1),

		('Hental', 'Aime', 21, 0, 5, 5, 18, 5, false, null, 0, 2),
		('Fort', 'Beau', 37, 1, 30, 70, 70, 20, false, null, 0, 2),
		('Embert', 'Cam', 40, 2, 60, 60, 30, 70, false, null, 0, 2),
		('Fort', 'Rock', 40, 1, 30, 70, 70, 20, false, null, 0, 2),
		('Dsavoie', 'Tom', 38, 2, 60, 60, 30, 70, false, null, 0, 2),

		('Ramplou', 'Denis', 31, 1, 30, 70, 70, 20, false, null, 0, 3),
		('Garrix', 'Martin', 33, 1, 30, 70, 70, 20, false, null, 0, 3),
		('Cuviller', 'Annalisa', 43, 2, 60, 60, 30, 70, false, null, 0, 3),
		('Fab', 'Léo', 25, 0, 5, 5, 20, 5, false, null, 0, 3),
		('Prussel', 'Loïc', 35, 2, 60, 60, 30, 70, false, null, 0, 3),

		('Mayeur', 'Lucie', 24, 0, 5, 5, 20, 5, false, null, 0, 4),
		('Cichy', 'Dominique', 64, 2, 60, 60, 30, 70, false, null, 0, 4),
		('Himana', 'Iqbal', 31, 1, 30, 70, 70, 20, false, null, 0, 4),
		('Leroux', 'Jean-Claude', 59, 2, 60, 60, 30, 70, false, null, 0, 4),
		('Belhadji', 'Chadia', 45, 1, 30, 70, 70, 20, false, null, 0, 4),

		('Grade', 'Julien', 30, 0, 5, 5, 20, 5, false, null, 0, 5), -- vane sur du sql
		('Lenglet', 'Julien', 30, 1, 30, 70, 70, 20, false, null, 0, 5),
		('Lallain', 'Clément', 30, 2, 60, 60, 30, 70, false, null, 0, 5),
		('Laplume', 'Etienne', 33, 2, 60, 60, 30, 70, false, null, 0, 5), -- attaque le réseau
		('Galice', 'Lucie', 35, 1, 30, 70, 70, 20, false, null, 0, 5);
    `)
	if err != nil {
		fmt.Println("Error in pushJoueur.go :", err)
	}
}
