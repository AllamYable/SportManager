package dbPack

import (
	"database/sql"
	"fmt"
)

func PushEquipe(db *sql.DB) {
	fmt.Println("pushing team...")
	_, err := db.Exec(`
        INSERT INTO equipe (nom_equipe, coach, nb_victoires, nb_defaites, nb_matchs, niveau_global, date_creation) VALUES
        ('Paprika', 'M. Poivre', 6, 1, 7, 2, '2018-01-01'),
        ('Parmesan', 'M. Bri', 5, 3, 8, 2, '2019-06-20'),
        ('Epitech', 'M. Jérôme', 2, 16, 18, 0, '2024-08-09'),
        ('Epsi', 'M. Croquette', 4, 0, 5, 3, '1961-03-30');
    `)
	if err != nil {
		fmt.Println("Error in pushEquipe.go :", err)
	}
	_, err = db.Exec(`
        INSERT INTO equipe (nom_equipe, coach, nb_victoires, nb_defaites, nb_matchs, niveau_global) VALUES
        ('Cesi', 'M. Allan & M. Yannis', 0, 0, 0, 2);
    `)
	if err != nil {
		fmt.Println("Error in pushEquipe.go 2 :", err)
	}
}


// p('Cesi', 'M. Allan & M. Yannis', 0, 0, 0, 2),
