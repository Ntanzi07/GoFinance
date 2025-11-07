package database

import (
	"database/sql"
	"fmt"

	"github.com/Ntanzi07/gofinance/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type User models.User

// GetAllUsers retrieves all users from the database.
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

// GetUserByID retrieves a user by their ID.
func GetUserByID(db *sql.DB, id int) (User, error) {
	var u User
	err := db.QueryRow("CALL GetUserById(?)", id).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt)
	if err != nil {
		return u, err
	}
	return u, nil
}

// CreateUser creates a new user with hashed password.
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

// DeleteUser deletes a user by ID after retrieving and printing their information.
func DeleteUser(db *sql.DB, userID int) error {
	user, err := GetUserByID(db, userID)
	if err != nil {
		return err
	}

	_, err = db.Exec("CALL DeleteUser(?)", userID)
	if err != nil {
		return err
	}

	fmt.Println("User deleted:")
	fmt.Println(printUser(user))
	return nil
}

func printUser(u User) string {
	userInfos := fmt.Sprintf("%d \t| %s \t| %s \t| %s \t| %s \t|\n", u.ID, u.Name, u.Email, u.Password, u.CreatedAt)
	return userInfos
}
