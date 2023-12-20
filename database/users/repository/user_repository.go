// Package repository provides functionality for interacting with user data in the database.
package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/yarincep/database-go-demo/database/users/model"
	"github.com/yarincep/database-go-demo/database/users/queries"
)

// UserRepository defines the methods for interacting with user data.
type UserRepository interface {
	// GetUserByID retrieves a user by their ID from the database.
	// It returns a pointer to the User and an error, if any.
	GetUserByID(ctx context.Context, userID int) (*model.User, error)
}

// UserRepositoryImplement represents the concrete implementation of the UserRepository interface.
// It encapsulates a *sql.DB instance for database interactions related to user data.
type UserRepositoryImplement struct {
	// db is a pointer to the *sql.DB instance representing the database connection.
	db *sql.DB
}

// NewUserRepository creates and returns a new instance of UserRepositoryImplement
// with the provided *sql.DB instance for database interactions.
//
// Parameters:
//   - db: A pointer to the *sql.DB instance representing the database connection.
//
// Returns:
//   - A pointer to a new UserRepositoryImplement instance.
//
// Example:
//
//	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/database")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer db.Close()
//
//	userRepository := NewUserRepository(db)
//	// Now you can use userRepository to interact with user data in the database.
func NewUserRepository(db *sql.DB) *UserRepositoryImplement {
	return &UserRepositoryImplement{
		db: db,
	}
}

// GetUserByID retrieves a user from the database based on the provided user ID.
// It queries the database using the associated *sql.DB instance in the UserRepositoryImplement.
//
// Parameters:
//   - ctx: The context.Context to handle the cancellation signal and deadlines.
//   - userID: An integer representing the unique identifier of the user to retrieve.
//
// Returns:
//   - A pointer to a model.User instance representing the retrieved user.
//   - An error, if any, indicating the success or failure of the operation.
//
// Example:
//
//	db, err := sql.Open("mysql", "user:password@tcp(localhost:3306)/database")
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer db.Close()
//
//	userRepository := NewUserRepository(db)
//	user, err := userRepository.GetUserByID(context.Background(), 123)
//	if err != nil {
//	    log.Printf("Error retrieving user: %v", err)
//	    return
//	}
//	fmt.Printf("Retrieved user: %+v\n", user)
func (r *UserRepositoryImplement) GetUserByID(ctx context.Context, userID int) (*model.User, error) {
	// Query the database for the user by ID
	row := r.db.QueryRowContext(ctx, queries.UserQueries.GetUserByID, userID)

	// Create a new User instance
	user := &model.User{}

	// Scan the row data into the User instance
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		// Return an error if the scanning process fails
		return nil, fmt.Errorf("failed to get user by ID: %v", err)
	}

	// Return the User instance and nil error on success
	return user, nil
}
