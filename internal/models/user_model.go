package models

// User represents a user in the system.
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
