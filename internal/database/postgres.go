// internal/database/postgres.go
package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Config struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
    SSLMode  string
}

func NewPostgresDB(config Config) (*sql.DB, error) {
    // Construction de la chaîne de connexion
    connStr := fmt.Sprintf(
        "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
        config.Host, 
        config.Port, 
        config.User, 
        config.Password, 
        config.DBName, 
        config.SSLMode,
    )
    
    // Ouverture de la connexion
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        return nil, fmt.Errorf("error opening database: %w", err)
    }

    // Configuration du pool de connexions
    db.SetMaxOpenConns(25)                // Nombre maximum de connexions ouvertes
    db.SetMaxIdleConns(25)                // Nombre maximum de connexions inactives
    db.SetConnMaxLifetime(5 * time.Minute) // Durée de vie maximum d'une connexion

    // Test de la connexion
    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("error connecting to the database: %w", err)
    }

    return db, nil
}

// Fonction utilitaire pour gérer les transactions
func BeginTx(db *sql.DB) (*sql.Tx, error) {
    tx, err := db.Begin()
    if err != nil {
        return nil, fmt.Errorf("error beginning transaction: %w", err)
    }
    return tx, nil
}

// Fonction helper pour vérifier l'état de la base de données
func CheckConnection(db *sql.DB) error {
    return db.Ping()
}

// Fonction pour fermer proprement la connexion
func CloseDB(db *sql.DB) error {
    return db.Close()
}