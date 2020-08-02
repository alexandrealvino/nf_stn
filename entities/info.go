package entities

// Info struct
type Info struct {
	AuthenticationStatus string    `json:"authenticationStatus"`
	RequestMethod        string    `json:"requestMethod"`
	ContentType          string    `json:"contentType"`
	Page                 string    `json:"page"`
	TotalPages           string    `json:"totalPages"`
	InvoicesFound        int       `json:"numberOfInvoices"`
	Invoices             []Invoice `json:"invoices"`
}

//
