package models

type Transaction struct {
	ID          int
	UserID      int
	Type        string
	Amount      float64
	Description string
	Date        string
}

type TransactionCreate struct {
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

// Notes:
// - Transaction uses `Date` as a string to keep the code simple; the handlers
//   normalize incoming date strings before passing them to the repository.
// - TransactionWithUser is used when joining transaction rows with user info
//   (for responses that include the user's name/email).
