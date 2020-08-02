package entities

// Invoice struct
type Invoice struct {
	ID             int     `json:"id"`
	ReferenceMonth int     `json:"referenceMonth"`
	ReferenceYear  int     `json:"referenceYear"`
	Document       string  `json:"document"`
	Description    string  `json:"description"`
	Amount         float64 `json:"amount"`
	IsActive       int     `json:"isActive"`
	CreatedAt      string  `json:"createdAt"`
	DeactivatedAt  string  `json:"deactivatedAt"`
}

//
