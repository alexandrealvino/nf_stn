package entities

// User credentials struct
type User struct {
	ID uint64            `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
