package entities


// Account class for account data
type Account struct {
	Profile string `json:"profile"`
	Hash []byte `json:"hash"`
}
