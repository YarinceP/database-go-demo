package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/yarincep/database-go-demo/database/users/model"
	"github.com/yarincep/database-go-demo/database/users/queries"
	"reflect"
	"testing"
)

func TestNewUserRepository(t *testing.T) {
	type args struct {
		db *sql.DB
	}
	tests := []struct {
		name string
		args args
		want *UserRepositoryImplement
	}{
		// TODO: Add test cases.
		{
			name: "OK",
			args: args{
				db: &sql.DB{},
			},
			want: &UserRepositoryImplement{
				db: &sql.DB{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRepository(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepositoryImplement_GetUserByID(t *testing.T) {
	// Configura el mock de la base de datos
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error al configurar el mock de la base de datos: %v", err)
	}
	defer db.Close()

	// Crea una instancia del repositorio con el mock de la base de datos
	repo := NewUserRepository(db)

	// Configura el mock para la consulta y el resultado
	mock.ExpectQuery(queries.UserQueries.GetUserByID).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(1, "John Doe"))

	// Llama a la función que queremos probar
	ctx := context.Background()
	user, err := repo.GetUserByID(ctx, 1)

	// Verifica que no haya errores
	assert.NoError(t, err)

	// Verifica que el resultado sea el esperado
	expectedUser := &model.User{
		ID:   1,
		Name: "John Doe",
	}
	assert.Equal(t, expectedUser, user)

	// Configura el mock para simular un error al escanear la fila
	mock.ExpectQuery(queries.UserQueries.GetUserByID).
		WithArgs(2).
		WillReturnError(errors.New("some error"))

	// Llama a la función con un ID que provocará un error al escanear la fila
	user, err = repo.GetUserByID(ctx, 2)

	// Verifica que haya un error esperado al escanear la fila
	assert.Error(t, err)
	assert.Nil(t, user)

	// Verifica que no haya llamadas no esperadas al mock
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
