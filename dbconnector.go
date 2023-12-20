package database_go_demo

import (
	"fmt"
	"github.com/yarincep/database-go-demo/database"
	service2 "github.com/yarincep/database-go-demo/database/service"
	"github.com/yarincep/database-go-demo/database/users/repository"

	_ "github.com/yarincep/database-go-demo/database"
	_ "github.com/yarincep/database-go-demo/database/users/repository"
	"github.com/yarincep/database-go-demo/database/users/service"
)

// Config contiene la configuración de la biblioteca.
type Config struct {
	DatabaseDSN string
}

// DBConnector representa la instancia principal de la biblioteca.
type DBConnector struct {
	UserService     *service.UserService
	DatabaseService *service2.DatabaseServiceImpl
	// Puedes agregar más servicios aquí según sea necesario
}

var Connector *DBConnector

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
	databaseService := service2.NewDatabaseServiceImpl(db)

	// Crea la instancia principal de la biblioteca
	connector := &DBConnector{
		UserService:     userService,
		DatabaseService: databaseService,
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

func RegisterDBConnector(dbConnector *DBConnector) {
	Connector = dbConnector
}

func ValidateDBConnector() error {
	if Connector == nil {
		return fmt.Errorf("Connector no registered ")
	}
	return nil
}
