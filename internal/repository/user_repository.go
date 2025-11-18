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

// GetUserByID retrieves a user by their ID.
func (r *UsersRepository) GetUserByID(id int) (models.User, error) {
	var u models.User
	err := r.DB.QueryRow("CALL GetUserById(?)", id).Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.Password,
		&u.CreatedAt,
	)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (r *UsersRepository) GetUserByName(name string) (models.User, error) {
	var u models.User
	err := r.DB.QueryRow("CALL GetUserByName(?)", name).Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.Password,
		&u.CreatedAt,
		&u.IsAdmin,
	)
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
	return nil
}

// DeleteUser deletes a user by ID after retrieving and printing their information.
func (r *UsersRepository) DeleteUser(userID int) error {
	_, err := r.GetUserByID(userID)
	if err != nil {
		return err
	}

	_, err = r.DB.Exec("CALL DeleteUser(?)", userID)
	if err != nil {
		return err
	}

	return nil
}

func (r *UsersRepository) UserLogin(email string) (models.UserLogin, error) {

	var user models.UserLogin
	err := r.DB.QueryRow("CALL GetUserCredentials(?)", email).Scan(&user.Email, &user.Password, &user.IsAdmin)
	if err != nil {
		return user, err
	}

	return user, nil

}

func (r *UsersRepository) GetAllUserTransactions(name string) ([]models.Transaction, error) {
	rows, err := r.DB.Query("call GetAllTransactionsByUser(?)", name)
	if err != nil {
		return []models.Transaction{}, err
	}

	var allTransactions []models.Transaction
	for rows.Next() {
		var t models.Transaction
		err := rows.Scan(&t.ID, &t.UserID, &t.Type, &t.Amount, &t.Description, &t.Date)
		if err != nil {
			return []models.Transaction{}, err
		}
		allTransactions = append(allTransactions, t)
	}

	return allTransactions, nil
}

func (r *UsersRepository) CreateUserTransaction(UserName string, tType string, amount float64, description, date string) error {
	user, err := r.GetUserByName(UserName)
	if err != nil {
		return err
	}

	_, err = r.DB.Exec("CALL CreateTransaction(?,?,?,?,?)", user.ID, tType, amount, description, date)
	if err != nil {
		return err
	}
	return nil
}

func (r *UsersRepository) UpdateUserTransaction(userName string, transactionID int, tType string, amount float64, description, date string) error {
	user, err := r.GetUserByName(userName)
	if err != nil {
		return err
	}

	var transaction models.Transaction
	err = r.DB.QueryRow("CALL GetTransactionById(?)", transactionID).Scan(
		&transaction.ID,
		&transaction.UserID,
		&transaction.Type,
		&transaction.Amount,
		&transaction.Description,
		&transaction.Date,
	)
	if err != nil {
		return err
	}

	if user.ID != transaction.UserID {
		return fmt.Errorf("you do not have permission to update this transaction")
	}

	_, err = r.DB.Exec("CALL UpdateUserTransaction(?,?,?,?,?)", transactionID, tType, amount, description, date)
	if err != nil {
		return err
	}
	return nil
}

func (r *UsersRepository) DeleteUserTransaction(userName string, transactionID int) error {
	user, err := r.GetUserByName(userName)
	if err != nil {
		return err
	}

	var transaction models.Transaction
	err = r.DB.QueryRow("CALL GetTransactionById(?)", transactionID).Scan(
		&transaction.ID,
		&transaction.UserID,
		&transaction.Type,
		&transaction.Amount,
		&transaction.Description,
		&transaction.Date,
	)
	if err != nil {
		return err
	}

	if user.ID != transaction.ID {
		return fmt.Errorf("you do not have permission to delete this transaction")
	}

	_, err = r.DB.Exec("CALL DeleteTransaction(?)", transactionID)
	if err != nil {
		return err
	}
	return nil
}
