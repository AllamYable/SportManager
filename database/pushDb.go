package database

import (
	"database/sql"
	"fmt"
	"sportmanager/database/dbPack"
)

func PushDatabase(db *sql.DB) {
	fmt.Println("pushing in DB...")
	dbPack.PushEquipe(db)
	dbPack.PushJoueur(db)
	dbPack.PushMatch(db)
	dbPack.PushPerf(db)
	dbPack.PushPoste(db)
}
