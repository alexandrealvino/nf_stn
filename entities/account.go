package entities

// User credentials struct
type User struct {
	ID 		 int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
