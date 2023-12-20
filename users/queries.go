package users

var UserQueries = struct {
	InsertUser string
}{
	InsertUser: "INSERT INTO users (ID, Name) VALUES (?, ?)",
}
