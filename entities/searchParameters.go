package entities

// SearchParameters struct
type SearchParameters struct {
	Page int `json:"page"`
	OrderBy string `json:"orderBy"`
	Month int `json:"month"`
	Year int `json:"year"`
	Document string `json:"document"`
}
//