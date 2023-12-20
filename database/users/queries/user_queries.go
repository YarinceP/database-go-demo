// Package queries provides pre-defined SQL queries for user-related operations.
package queries

// UserQueries contains pre-defined SQL queries for user-related operations.
var UserQueries = struct {
	// GetUserByID is the SQL query to retrieve a user by their ID.
	// Example usage: db.QueryRowContext(ctx, UserQueries.GetUserByID, userID)
	GetUserByID string
}{
	GetUserByID: "SELECT * FROM users WHERE id = ?",
}
