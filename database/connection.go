package database

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type ConnectorRepository interface {
	Connect() (*sql.DB, error)
	Disconnect() error
}

var connection *sql.DB

type ConnectorRepositoryImplement struct {
	ConnectionString string
	db               *sql.DB
}

// NewConnectorRepositoryImplement crea una nueva instancia de ConnectorRepositoryImplement.
//
// NewConnectorRepositoryImplement toma una cadena de conexión como parámetro
// y devuelve una nueva instancia de ConnectorRepositoryImplement con la
// cadena de conexión proporcionada. Esta función se utiliza para inicializar
// la estructura del repositorio de conectores con la información necesaria
// para establecer una conexión a la base de datos.
//
// Parámetros:
// - connectionString: La cadena de conexión que se utilizará para conectarse a la base de datos.
//
// Devuelve:
// - Una nueva instancia de ConnectorRepositoryImplement inicializada con la cadena de conexión.
func NewConnectorRepositoryImplement(connectionString string) ConnectorRepositoryImplement {
	return ConnectorRepositoryImplement{ConnectionString: connectionString}
}

func (c *ConnectorRepositoryImplement) Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", c.ConnectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	c.db = db
	log.Println("Conexión exitosa a la base de datos")
	return db, nil
}

// Disconnect cierra la conexión a la base de datos.
//
// Disconnect es responsable de cerrar la conexión activa a la base de datos
// mantenida por la instancia de ConnectorRepositoryImplement. Primero verifica
// si hay una conexión activa para cerrar. Si no se encuentra ninguna conexión,
// devuelve un error indicando la ausencia de una conexión de base de datos activa.
//
// Si hay una conexión presente, la función intenta cerrarla. Cualquier error
// que ocurra durante el proceso de cierre se devuelve con un mensaje de contexto
// adicional. Si la conexión se cierra con éxito, se imprime un mensaje de registro
// para indicar una desconexión exitosa.
//
// Devuelve:
// - nil en caso de desconexión exitosa.
// - Error si no hay una conexión activa o si ocurre un error durante el cierre.
func (c *ConnectorRepositoryImplement) Disconnect() error {
	if c.db == nil {
		return errors.New("No hay conexión de base de datos activa para cerrar")
	}

	err := c.db.Close()
	if err != nil {
		return fmt.Errorf("Error al cerrar la conexión de la base de datos: %v", err)
	}

	log.Println("Desconexión exitosa de la base de datos")
	return nil
}

func (c *ConnectorRepositoryImplement) RegisterConnectionDB() {
	if c.db != nil {
		connection = c.db
	}

}
