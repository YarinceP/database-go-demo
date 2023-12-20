package main

import (
	"context"
	"database-go-demo/users"
	"database/sql"
	"fmt"
	"time"
)

func InsertNewUserDefault(c *sql.DB) {
	user := users.User{
		Name: "Example Name",
	}

	repository := users.NewUserRepositoryImplement(c)

	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	err := repository.Save(timeout, &user)
	if err != nil {
		fmt.Printf("InsertNewUserDefault error: %v ", err)
	}
}
