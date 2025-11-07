package repository

import (
	"database/sql"
	"fmt"

	"github.com/Ntanzi07/gofinance/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UsersRepository struct {
	DB *sql.DB
}

func NewUsersRepository(db *sql.DB) *UsersRepository {
	return &UsersRepository{DB: db}
}

// GetAllUsers retrieves all users from the database.
func (r *UsersRepository) GetAllUsers() ([]models.User, error) {
	rows, err := r.DB.Query("CALL GetAllUsers()")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	fmt.Println("Users loaded:", len(users))
	return users, nil
}

// GetUserByID retrieves a user by their ID.
func (r *UsersRepository) GetUserByID(id int) (models.User, error) {
	var u models.User
	err := r.DB.QueryRow("CALL GetUserById(?)", id).Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.CreatedAt)
	if err != nil {
		return u, err
	}
	return u, nil
}

// CreateUser creates a new user with hashed password.
func (r *UsersRepository) CreateUser(name, email, password string) error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = r.DB.Exec("CALL CreateUser(?,?,?)", name, email, string(hashed))
	if err != nil {
		return err
	}
	fmt.Println("User created:", email)
	return nil
}

// DeleteUser deletes a user by ID after retrieving and printing their information.
func (r *UsersRepository) DeleteUser(userID int) error {
	user, err := r.GetUserByID(userID)
	if err != nil {
		return err
	}

	_, err = r.DB.Exec("CALL DeleteUser(?)", userID)
	if err != nil {
		return err
	}

	fmt.Printf("User deleted with ID: %d", user.ID)
	return nil
}
