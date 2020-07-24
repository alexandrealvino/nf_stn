package entities


// Invoice class for data storage
type Invoice struct {
	ID     int `json:"id"`
	ReferenceMonth int `json:"referenceMonth"`
	ReferenceYear  int `json:"referenceYear"`
	Document string `json:"document"`
	Description string `json:"description"`
	Amount float64 `json:"amount"`
	IsActive int `json:"isActive"`
	CreatedAt string `json:"createdAt"`
	DeactivatedAt string `json:"deactivatedAt"`
}


//# Invoice
//#     ReferenceMonth : INTEGER
//#     ReferenceYear : INTEGER
//#     Document : VARCHAR(14)
//#     Description : VARCHAR(256)
//#     Amount : DECIMAL(16, 2)
//#     IsActive : TINYINT
//#     CreatedAt  : DATETIME
//#     DeactivatedAt : DATETIME