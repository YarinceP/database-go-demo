package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Importa el driver de MySQL (o el que estés utilizando)
)

var db *sql.DB

// InitDB inicializa la conexión a la base de datos.
func InitDB(dsn string) (*sql.DB, error) {
	var err error

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %v", err)
	}

	// Intenta conectar para asegurarse de que la conexión es válida
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	fmt.Println("Connected to the database")

	return db, nil
}

// CloseDB cierra la conexión a la base de datos.
func CloseDB() error {
	if db != nil {
		if err := db.Close(); err != nil {
			return fmt.Errorf("failed to close database connection: %v", err)
		}
		fmt.Println("Closed the database connection")
	}
	return nil
}
