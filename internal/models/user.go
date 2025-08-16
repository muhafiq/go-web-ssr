package models

import (
	"database/sql"
	"fmt"
	"go-web-ssr/internal/config"
	"log"
)

type User struct {
	ID        string
	Name      string
	Email     string
	Password  string
	CreatedAt int
	UpdatedAt int
}

func (u *User) GetByEmail(email string) (*User, error) {
	query := "SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = ?"
	row := config.DB.QueryRow(query, email)

	result := &User{}
	err := row.Scan(&result.ID, &result.Name, &result.Email, &result.Password, &result.CreatedAt, &result.UpdatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		log.Println("Error fetching user", err)
		return nil, err
	}

	fmt.Print(result)
	return result, nil
}
