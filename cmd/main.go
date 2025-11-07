package main

import (
	"fmt"

	"github.com/Ntanzi07/gofinance/internal/database"
	"github.com/Ntanzi07/gofinance/internal/repository"
)

func main() {
	//connect to database
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	if err = repository.CreateTransaction(db, 1, "income", 1917.18, "Freelance project", "15/11/2025"); err != nil {
		panic(err)
	}

	transaction, err := repository.GetTransactionByID(db, 3)
	if err != nil {
		panic(err)
	}
	fmt.Println("Transaction retrieved:", transaction)

	transactions, err := repository.GetAllTransactions(db)
	if err != nil {
		panic(err)
	}
	for _, t := range transactions {
		fmt.Println(t)
	}

	/*
		if err = repository.DeleteTransaction(db, 1); err != nil {
			panic(err)
		}
	*/

}
