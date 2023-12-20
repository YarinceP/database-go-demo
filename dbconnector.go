package database_go_demo

import (
	"database-go-demo/database"
	"database-go-demo/database/users/repository"
	"fmt"

	_ "database-go-demo/database"
	_ "database-go-demo/database/users/repository"
	"database-go-demo/database/users/service"
)

// Config contiene la configuración de la biblioteca.
type Config struct {
	DatabaseDSN string
}

// DBConnector representa la instancia principal de la biblioteca.
type DBConnector struct {
	UserService *service.UserService
	// Puedes agregar más servicios aquí según sea necesario
}

// NewDBConnector crea una nueva instancia de la biblioteca.
func NewDBConnector(config Config) (*DBConnector, error) {
	// Inicializa la conexión a la base de datos
	db, err := database.InitDB(config.DatabaseDSN)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database: %v", err)
	}

	// Crea instancias de repositorios y servicios
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)

	// Crea la instancia principal de la biblioteca
	connector := &DBConnector{
		UserService: userService,
		// Puedes agregar más servicios aquí según sea necesario
	}

	return connector, nil
}

// CloseDBConnector cierra la conexión a la base de datos y realiza otras tareas de limpieza.
func (connector *DBConnector) CloseDBConnector() error {
	// Cierra la conexión a la base de datos
	err := database.CloseDB()
	if err != nil {
		return fmt.Errorf("failed to close database: %v", err)
	}

	// Realiza otras tareas de limpieza si es necesario

	return nil
}
