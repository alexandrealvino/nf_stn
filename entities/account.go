package entities

// User credentials

type User struct {
	ID uint64            `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
