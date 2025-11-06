package main

import "github.com/Ntanzi07/gofinance/internal/database"

func main() {
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	err = database.CreateUser(db, "Nathan", "nathan@example.com", "123456")
	if err != nil {
		panic(err)
	}

	users, err := database.GetAllUsers(db)
	if err != nil {
		panic(err)
	}

	for _, user := range users {
		println("User:", user.Name, "Email:", user.Email)
	}
}
