package database

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int
	Name      string
	Email     string
	Password  string
	CreatedAt string
}

func GetAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("CALL GetAllUsers()")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	fmt.Println("Users loaded:", len(users))
	return users, nil
}

func CreateUser(db *sql.DB, name, email, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = db.Exec("CALL CreateUser(?,?,?)", name, email, string(hashed))
	if err != nil {
		return err
	}
	fmt.Println("User created:", email)
	return nil
}
