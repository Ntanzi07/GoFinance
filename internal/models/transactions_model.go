package models

type Transaction struct {
	ID          int
	UserID      int
	Type        string
	Amount      float64
	Description string
	Date        string
}

type TransactionWithUser struct {
	ID          int
	Type        string
	Amount      float64
	Description string
	Date        string
	UserName    string
	UserEmail   string
}
