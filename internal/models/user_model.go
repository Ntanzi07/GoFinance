package models

// User represents an application user stored in the database.
//
// Notes:
//   - `Password` holds the hashed password in DB structs. Never return
//     this field in API responses; handlers should omit or sanitize it.
//   - `CreatedAt` is a string here to keep the model simple; consider using
//     time.Time if you need date arithmetic.
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	IsAdmin   bool   `json:"is_admin"`
}

// UserSingUp is the payload accepted when registering a new user.
// Handlers should hash the `Password` before calling repository methods.
type UserSingUp struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserLogin represents credential data retrieved from the DB for login
// verification (email, hashed password and admin flag).
type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	IsAdmin  bool   `json:"is_admin"`
}

// UserCreds is the JSON body shape expected from login requests.
type UserCreds struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// TODO: Consider adding a sanitized UserResponse type that omits Password
// and any sensitive fields when returning user data through the API.
