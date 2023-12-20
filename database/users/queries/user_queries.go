package queries

var UserQueries = struct {
	GetUserByID string
}{
	GetUserByID: "SELECT id, name FROM users WHERE id = ?",
}
