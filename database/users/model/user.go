package model

import "fmt"

type User struct {
	ID   int
	Name string
}

// Validate method
func (u *User) Validate() error {
	// Lógica de validación según tus requisitos
	if u.Name == "" {
		return fmt.Errorf("el nombre de usuario no puede estar vacío")
	}
	return nil
}
