package main

import (
	"database-go-demo/database"
	"database/sql"
	"log"
)

var Connection *sql.DB

func main() {
	connector := database.NewConnectorRepositoryImplement("root:@tcp(localhost:3306)/db_lib_go")
	Connection, err := connector.Connect()
	if err != nil {
		log.Fatal("Error trying to connect db : ", err)
	}

	InsertNewUserDefault(Connection)

}
