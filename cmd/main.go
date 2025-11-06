package main

import (
	"fmt"

	"github.com/Ntanzi07/gofinance/internal/database"
)

func main() {
	//connect to database
	db, err := database.Connect()
	if err != nil {
		panic(err)
	}

	//create a new user
	err = database.CreateUser(db, "Teste3", "teste3@example.com", "123456")
	if err != nil {
		panic(err)
	}

	//get user by ID
	user, err := database.GetUserByID(db, 1)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User ID: %d \t| %s \t| %s \t| %s \t| %s \t|\n", user.ID, user.Name, user.Email, user.Password, user.CreatedAt)

	//get all users
	users, err := database.GetAllUsers(db)
	if err != nil {
		panic(err)
	}

	for _, u := range users {
		fmt.Printf("%d \t| %s \t| %s \t| %s \t| %s \t|\n", u.ID, u.Name, u.Email, u.Password, u.CreatedAt)
	}

	//delete user by ID
	err = database.DeleteUser(db, 3)
	if err != nil {
		panic(err)
	}

	//get all users after deletion
	users, err = database.GetAllUsers(db)
	if err != nil {
		panic(err)
	}

	for _, u := range users {
		fmt.Printf("%d \t| %s \t| %s \t| %s \t| %s \t|\n", u.ID, u.Name, u.Email, u.Password, u.CreatedAt)
	}

}
