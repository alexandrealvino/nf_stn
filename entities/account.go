package entities

// User credentials struct
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Hash     string `json:"hash"`
}

//
