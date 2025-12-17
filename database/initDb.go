package database

import "database/sql"

// CreateSchema crée les tables nécessaires dans la base de données si elles n'existent pas déjà.
func InitDatabase(db *sql.DB) error {
	schema :=

	`
	CREATE TABLE IF NOT EXISTS equipe (
    id_equipe INTEGER PRIMARY KEY AUTOINCREMENT,
    nom_equipe TEXT NOT NULL UNIQUE,
    coach TEXT NOT NULL,
    nb_victoires INTEGER DEFAULT 0,
    nb_defaites INTEGER DEFAULT 0,
    nb_matchs INTEGER DEFAULT 0,
    niveau_global INTEGER DEFAULT 0,
    date_creation DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS poste (
    id_poste INTEGER PRIMARY KEY AUTOINCREMENT,
    nom_poste TEXT NOT NULL UNIQUE,
    description_poste TEXT,
    min_vitesse INTEGER,
    min_endurance INTEGER,
    min_force INTEGER,
    min_technique INTEGER
);

CREATE TABLE IF NOT EXISTS joueur (
    id_joueur INTEGER PRIMARY KEY AUTOINCREMENT,
    nom_joueur TEXT NOT NULL,
    prenom_joueur TEXT NOT NULL,
    age INTEGER NOT NULL,
    id_poste INTEGER NOT NULL,
    vitesse INTEGER,
    endurance INTEGER,
    force INTEGER,
    technique INTEGER,
    blesse INTEGER DEFAULT 0,
    date_blessure DATETIME,
    matchs_absence INTEGER DEFAULT 0,
    id_equipe INTEGER NOT NULL,
    date_creation DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_equipe) REFERENCES equipe(id_equipe) ON DELETE CASCADE,
    FOREIGN KEY (id_poste) REFERENCES poste(id_poste) ON DELETE RESTRICT
);

CREATE TABLE IF NOT EXISTS match (
    id_match INTEGER PRIMARY KEY AUTOINCREMENT,
    date_match DATETIME NOT NULL,
    id_equipe_domicile INTEGER NOT NULL,
    id_equipe_exterieur INTEGER NOT NULL,
    score_domicile INTEGER DEFAULT 0,
    score_exterieur INTEGER DEFAULT 0,
    gagnant INTEGER,
    date_creation DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_equipe_domicile) REFERENCES equipe(id_equipe) ON DELETE CASCADE,
    FOREIGN KEY (id_equipe_exterieur) REFERENCES equipe(id_equipe) ON DELETE CASCADE,
    FOREIGN KEY (gagnant) REFERENCES equipe(id_equipe) ON DELETE SET NULL,
    CHECK (id_equipe_domicile != id_equipe_exterieur)
);

CREATE TABLE IF NOT EXISTS performance (
    id_performance INTEGER PRIMARY KEY AUTOINCREMENT,
    id_joueur INTEGER NOT NULL,
    id_match INTEGER NOT NULL,
    note_performance INTEGER,
    buts INTEGER DEFAULT 0,
    passe_decisives INTEGER DEFAULT 0,
    fautes INTEGER DEFAULT 0,
    temps_jeu INTEGER DEFAULT 90,
    date_creation DATETIME DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (id_joueur, id_match),
    FOREIGN KEY (id_joueur) REFERENCES joueur(id_joueur) ON DELETE CASCADE,
    FOREIGN KEY (id_match) REFERENCES match(id_match) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS dispute (
    id_dispute INTEGER PRIMARY KEY AUTOINCREMENT,
    id_match INTEGER NOT NULL,
    id_equipe INTEGER NOT NULL,
    date_creation DATETIME DEFAULT CURRENT_TIMESTAMP,
    UNIQUE (id_match, id_equipe),
    FOREIGN KEY (id_match) REFERENCES match(id_match) ON DELETE CASCADE,
    FOREIGN KEY (id_equipe) REFERENCES equipe(id_equipe) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS blessure (
    id_blessure INTEGER PRIMARY KEY AUTOINCREMENT,
    id_joueur INTEGER NOT NULL,
    id_match INTEGER NOT NULL,
    date_blessure DATETIME DEFAULT CURRENT_TIMESTAMP,
    matchs_restants INTEGER DEFAULT 3,
    description TEXT,
    date_guerison DATETIME,
    FOREIGN KEY (id_joueur) REFERENCES joueur(id_joueur) ON DELETE CASCADE,
    FOREIGN KEY (id_match) REFERENCES match(id_match) ON DELETE CASCADE
`
	_, err := db.Exec(schema)
	return err
}
