package entities

// Account class for account data
type Account struct {
	Profile string `json:"profile"`
	Hash    []byte `json:"hash"`
}

type User struct {
	ID uint64            `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
