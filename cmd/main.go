package main

import "github.com/Ntanzi07/gofinance/internal/database"

func main() {
	_, err := database.Connect()
	if err != nil {
		panic(err)
	}
}
