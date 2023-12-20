package database

import (
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInitDB(t *testing.T) {
	tests := []struct {
		name        string
		dsn         string
		mockConfig  func(mock sqlmock.Sqlmock)
		wantError   bool
		expectedMsg string
	}{
		{
			name: "OpenConnectionError",
			dsn:  "bad-connection-string",
			mockConfig: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT 1").WillReturnError(fmt.Errorf("expected error"))
			},
			wantError:   true,
			expectedMsg: "failed to open database connection: invalid DSN: missing the slash separating the database name",
		},
		{
			name: "Success",
			dsn:  "root:@tcp(localhost:3306)/db_lib_go",
			mockConfig: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("SELECT 1").WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))
			},
			wantError:   false,
			expectedMsg: "Connected to the database",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Crea una instancia de sqlmock
			mockDB, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("Error al crear instancia de sqlmock: %v", err)
			}
			defer mockDB.Close()

			// Configura el mock según las expectativas del caso de prueba
			tt.mockConfig(mock)

			// Llama a la función que estás probando
			db, err := InitDB(tt.dsn)

			// Verifica que el error sea el esperado
			if tt.name != "Success" {
				assert.EqualError(t, err, tt.expectedMsg)
			}

			// Verifica si se esperaba un error o no
			if tt.wantError {
				assert.Nil(t, db, "No se esperaba una instancia de DB")
			} else {
				assert.NotNil(t, db, "Se esperaba una instancia de DB")
			}
		})
	}
}

func TestCloseDB(t *testing.T) {
	tests := []struct {
		name        string
		mockConfig  func(mock sqlmock.Sqlmock)
		wantError   bool
		expectedMsg string
	}{
		{
			name: "Success",
			mockConfig: func(mock sqlmock.Sqlmock) {
				// Puedes agregar configuraciones específicas de cierre si es necesario
				mock.ExpectClose()
			},
			wantError:   false,
			expectedMsg: "Closed the database connection",
		},
		{
			name: "ErrorOnClose",
			mockConfig: func(mock sqlmock.Sqlmock) {
				mock.ExpectClose().WillReturnError(fmt.Errorf("expected error"))
			},
			wantError:   true,
			expectedMsg: "failed to close database connection: expected error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("Error al crear instancia de sqlmock: %v", err)
			}
			defer mockDB.Close()

			tt.mockConfig(mock)

			db = mockDB // Establece la instancia de db para la prueba

			err = CloseDB()

			if tt.wantError {
				assert.EqualError(t, err, tt.expectedMsg)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
