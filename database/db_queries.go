package database

import (
	"fmt"
	"nf_stn/config"
	"nf_stn/entities"
)

//ReferenceMonth int `json:"referenceMonth"`
//ReferenceYear  int `json:"referenceYear"`
//Document string `json:"document"`
//Description string `json:"description"`
//Amount float64 `json:"amount"`
//IsActive int `json:"isActive"`
//CreatedAt time.Time `json:"createdAt"`
//DeactivatedAt time.Time `json:"deactivatedAt"`
//}

// Instantiating the App object for the db connection
var db config.App

// init initializes the db connection
func init() {
	db.Initialize(config.Dbdriver, config.Dbuser, config.Dbpass, config.Dbname)
}

// GetAll gets all the rows of the invoices db
func GetAll() ([]entities.Invoice, error) { // get list of all invoices
	results, err := db.Db.Query("SELECT id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM invoices ORDER BY id ASC")
	if err != nil {
		panic(err.Error())
	}
	inv := entities.Invoice{}
	invoicesList := []entities.Invoice{}
	for results.Next() {
		err = results.Scan(&inv.ID, &inv.ReferenceMonth, &inv.ReferenceYear, &inv.Document, &inv.Description, &inv.Amount, &inv.IsActive, &inv.CreatedAt, &inv.DeactivatedAt)
		if err != nil {
			panic(err.Error())
		}
		invoicesList = append(invoicesList, inv)
	}
	fmt.Println("Successfuly got invoice list!")
	return invoicesList, err
}

// GetInvoiceByDocument gets the invoice by the document value
func GetInvoiceByDocument(document string) (entities.Invoice, error) { // get invoice by document
	result, err := db.Db.Query("SELECT id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM invoices WHERE document = ?;", document)
	if err != nil {
		panic(err.Error())
	}
	inv := entities.Invoice{}
	for result.Next() {
		err = result.Scan(&inv.ID, &inv.ReferenceMonth, &inv.ReferenceYear, &inv.Document, &inv.Description, &inv.Amount, &inv.IsActive, &inv.CreatedAt, &inv.DeactivatedAt)
		if err != nil {
			panic(err.Error())
		}
	}
	fmt.Println("Successfuly got invoice!")
	return inv, err
}

// GetAccountByProfile gets the account data by the given profile
func GetAccountByProfile(profile string) (string, []byte, error) { // get acc data by profile
	result, err := db.Db.Query("SELECT profile, hash FROM hashes WHERE profile = ?;", profile)
	if err != nil {
		panic(err.Error())
	}
	var acc entities.Account
	for result.Next() {
		err = result.Scan(&acc.Profile, &acc.Hash)
		if err != nil {
			panic(err.Error())
		}
	}
	fmt.Println("Successfuly got hash!")
	return acc.Profile, acc.Hash, err
}

// GetInvoiceByID gets the invoice by the given ID
func GetInvoiceByID(id int) (entities.Invoice, error) { // get ticker by id
	result, err := db.Db.Query("SELECT id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM invoices WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
	inv := entities.Invoice{}
	for result.Next() {
		err = result.Scan(&inv.ID, &inv.ReferenceMonth, &inv.ReferenceYear, &inv.Document, &inv.Description, &inv.Amount, &inv.IsActive, &inv.CreatedAt, &inv.DeactivatedAt)
		if err != nil {
			panic(err.Error())
		}
	}
	if (inv != entities.Invoice{}) {
		fmt.Println("Successfuly got invoice!")
	}
	return inv, err
}

// InsertInvoice inserts the given invoice to invoices db
func InsertInvoice(invoice entities.Invoice) error { // insert invoice
	//monthDay,month,hour,min,sec,year := time.Now().Day(),time.Now().Month(),time.Now().Hour(),time.Now().Minute(),time.Now().Second(),time.Now().Year()
	//date := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(monthDay)
	//clock := strconv.Itoa(hour) + ":" + strconv.Itoa(min) + ":" + strconv.Itoa(sec)
	//now := date + " " + clock
	//fmt.Println(now)
	_, err := db.Db.Query("INSERT INTO invoices (referenceMonth, referenceYear, document , description, amount, createdAt, deactivatedAt) VALUES (?, ?, ?, ?, ?, ?, ?);",
		invoice.ReferenceMonth, invoice.ReferenceYear, invoice.Document, invoice.Description, invoice.Amount, invoice.CreatedAt, invoice.DeactivatedAt)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfuly inserted invoice!")
	return err
}

// DeleteInvoice makes the logic deletion setting isActive = 0 by the given ID
func DeleteInvoice(id int) error { // set isActive = 0 for logic deletion
	_, err := db.Db.Exec("UPDATE invoices SET isActive = ? WHERE id = ?;", 0, id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfuly deleted invoice!")
	return err
}

// UpdateInvoice updates database values from the row of the given invoice
func UpdateInvoice(invoice entities.Invoice) error { // update invoice
	_, err := db.Db.Exec("UPDATE invoices  SET referenceMonth=?, referenceYear=?, document=?, description=?, amount=?, isActive=?, createdAt=?, deactivatedAt=? WHERE id = ?;", invoice.ReferenceMonth, invoice.ReferenceYear, invoice.Document, invoice.Description, invoice.Amount, invoice.IsActive, invoice.CreatedAt, invoice.DeactivatedAt, invoice.ID)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfuly updated invoice!")
	return err
}

// PatchInvoice partially updates database values from the row of the given invoice
func PatchInvoice(invoice entities.Invoice) error { // update invoice
	_, err := db.Db.Exec("UPDATE invoices  SET referenceMonth=?, referenceYear=?, description=?, amount=? WHERE id = ?;", invoice.ReferenceMonth, invoice.ReferenceYear, invoice.Description, invoice.Amount, invoice.ID)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfuly updated invoice!")
	return err
}

// InvoiceExists checks if the given invoice document exists
func InvoiceExists(document string) (entities.Invoice, error) { // checks if invoice is already exists
	result, err := db.Db.Query("SELECT id FROM invoices WHERE document = ?", document)
	if err != nil {
		panic(err.Error())
	}
	invoice := entities.Invoice{}
	for result.Next() {
		err = result.Scan(&invoice.ID)
		if err != nil {
			panic(err.Error())
		}
	}
	if (invoice != entities.Invoice{}) {
		fmt.Println("Invoice exists!")
	}
	return invoice, err
}

// ClearTable truncates the invoices table
func ClearTable() {
	_, _ = db.Db.Exec("TRUNCATE TABLE invoices")
}

//
