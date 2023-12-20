package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// InitDB inicializa la conexión a la base de datos.
func InitDB(dsn string) (*sql.DB, error) {
	var err error

	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %v", err)
	}

	//// Maximum Idle Connections
	//db.SetMaxIdleConns(getMaxIdleConn())
	//// Maximum Open Connections
	//db.SetMaxOpenConns(getMaxOpenConn())
	//// Idle Connection Timeout
	//maxIdleTime := getMaxMaxIdleTime()
	//db.SetConnMaxIdleTime(time.Duration(maxIdleTime) * time.Second)
	//// Connection Lifetime
	//lifetime := getMaxMaxLifetime()
	//db.SetConnMaxLifetime(time.Duration(lifetime) * time.Second)
	//
	//contextTime := getContextTime()
	//ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(contextTime)*time.Second)
	//defer cancelFunc()
	//if err = db.PingContext(ctx); err != nil {
	//	return nil, err
	//}

	fmt.Println("Connected to the database")

	return db, nil
}

// CloseDB closes the connection to the database.
// It checks if the database connection is not nil before attempting to close it.
//
// Returns:
//   - An error, if any, indicating the success or failure of the database connection closure.
//     Returns nil if the database connection was successfully closed.
//
// Example:
//
//	err := CloseDB()
//	if err != nil {
//	    log.Printf("Error closing database connection: %v", err)
//	    return
//	}
//	fmt.Println("Successfully closed the database connection")
func CloseDB() error {
	// Check if the database connection is not nil
	if db != nil {
		// Attempt to close the database connection
		if err := db.Close(); err != nil {
			// Return an error if closing the database connection fails
			return fmt.Errorf("failed to close database connection: %v", err)
		}
		fmt.Println("Closed the database connection")
	}

	// Return nil if the database connection was successfully closed
	return nil
}
