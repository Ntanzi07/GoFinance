package repository

import (
	"database/sql"
	"fmt"

	"github.com/Ntanzi07/gofinance/internal/models"
)

type Transaction models.Transaction
type TransactionWithUser models.TransactionWithUser

func GetAllTransactions(db *sql.DB) ([]TransactionWithUser, error) {
	rows, err := db.Query("Call GetAllTransactions()")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []TransactionWithUser
	for rows.Next() {
		var t TransactionWithUser
		if err := rows.Scan(&t.ID, &t.Type, &t.Amount, &t.Description, &t.Date, &t.UserName, &t.UserEmail); err != nil {
			return nil, err
		}
		transactions = append(transactions, t)

	}
	return transactions, nil
}

func GetTransactionByID(db *sql.DB, id int) (TransactionWithUser, error) {
	var t TransactionWithUser
	err := db.QueryRow("CALL GetTransactionById(?)", id).Scan(&t.ID, &t.Type, &t.Amount, &t.Description, &t.Date, &t.UserName, &t.UserEmail)
	if err != nil {
		return TransactionWithUser{}, err
	}
	return t, nil
}

func CreateTransaction(db *sql.DB, userID int, tType string, amount float64, description, date string) error {
	_, err := db.Exec("CALL CreateTransaction(?,?,?,?,?)", userID, tType, amount, description, date)
	if err != nil {
		return err
	}
	fmt.Println("Transaction created for user ID:", userID)
	return nil
}

func DeleteTransaction(db *sql.DB, transactionID int) error {
	_, err := db.Exec("CALL DeleteTransaction(?)", transactionID)
	if err != nil {
		return err
	}
	fmt.Println("Transaction deleted with ID:", transactionID)
	return nil
}
