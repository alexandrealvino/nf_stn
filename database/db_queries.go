package database

import (
	"fmt"
	"nf_stn/config"
	"nf_stn/entities"
)

//go:generate  go run github.com/golang/mock/mockgen  -package mock -destination=./mock/db_mock.go -source=$GOFILE

type DataBase interface {
	Init()
	GetAll() ([]entities.Invoice, error)
	GetInvoiceByDocument(document string) (entities.Invoice, error)
	GetAccountByProfile(profile string) (string, []byte, error)
	GetInvoiceByID(id int) (entities.Invoice, error)
	InsertInvoice(invoice entities.Invoice) error
	DeleteInvoice(id int) error
	UpdateInvoice(invoice entities.Invoice) error
	PatchInvoice(invoice entities.Invoice) error
	InvoiceExists(document string) (entities.Invoice, error)
	ClearTable()
	Pagination(offset int) ([]entities.Invoice, error)
	PaginationOrderByMonth(offset int) ([]entities.Invoice, error)
	PaginationOrderByYear(offset int) ([]entities.Invoice, error)
	PaginationOrderByDocument(offset int) ([]entities.Invoice, error)
	PaginationByMonth(offset, referenceMonth int) ([]entities.Invoice, error)
	PaginationByYear(offset, referenceYear int) ([]entities.Invoice, error)
	PaginationByDocument(offset int, document string) ([]entities.Invoice, error)
}

// Instantiating the App object for the db connection



type MySql struct {
	Config config.DataBaseConfig
}

var db config.App

// init initializes the db connection
func (ms *MySql) Init()  {
	db.Initialize(ms.Config.Dbdriver(), ms.Config.Dbuser(), ms.Config.Dbpass(), ms.Config.Dbname())
}

// GetAll gets all the rows of the invoices db TODO paginação passando parametros
func (ms *MySql) GetAll() ([]entities.Invoice, error) { // get list of all invoices
	results, err := db.Db.Query("SELECT id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices ORDER BY id ASC")
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
	fmt.Println("Successfully got invoice list!")
	return invoicesList, err
}

// GetInvoiceByDocument gets the invoice by the document value
func (ms *MySql) GetInvoiceByDocument(document string) (entities.Invoice, error) { // get invoice by document
	result, err := db.Db.Query("SELECT id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices WHERE document = ?;", document)
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
	fmt.Println("Successfully got invoice!")
	return inv, err
}

// GetAccountByProfile gets the account data by the given profile
func (ms *MySql) GetAccountByProfile(profile string) (string, []byte, error) { // get acc data by profile
	result, err := db.Db.Query("SELECT profile, hash FROM nf_stn.hashes WHERE profile = ?;", profile)
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
	fmt.Println("Successfully got hash!")
	return acc.Profile, acc.Hash, err
}

// GetInvoiceByID gets the invoice by the given ID
func (ms *MySql) GetInvoiceByID(id int) (entities.Invoice, error) { // get ticker by id
	result, err := db.Db.Query("SELECT id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices WHERE id = ?", id)
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
		fmt.Println("Successfully got invoice!")
	}
	return inv, err
}

// InsertInvoice inserts the given invoice to invoices db
func (ms *MySql) InsertInvoice(invoice entities.Invoice) error { // insert invoice
	//monthDay,month,hour,min,sec,year := time.Now().Day(),time.Now().Month(),time.Now().Hour(),time.Now().Minute(),time.Now().Second(),time.Now().Year()
	//date := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(monthDay)
	//clock := strconv.Itoa(hour) + ":" + strconv.Itoa(min) + ":" + strconv.Itoa(sec)
	//now := date + " " + clock
	//fmt.Println(now)
	_, err := db.Db.Query("INSERT INTO nf_stn.invoices (referenceMonth, referenceYear, document , description, amount, createdAt, deactivatedAt) VALUES (?, ?, ?, ?, ?, ?, ?);",
		invoice.ReferenceMonth, invoice.ReferenceYear, invoice.Document, invoice.Description, invoice.Amount, invoice.CreatedAt, invoice.DeactivatedAt)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfully inserted invoice!")
	return err
}

// DeleteInvoice makes the logic deletion setting isActive = 0 by the given ID
func (ms *MySql) DeleteInvoice(id int) error { // set isActive = 0 for logic deletion
	_, err := db.Db.Exec("UPDATE nf_stn.invoices SET isActive = ? WHERE id = ?;", 0, id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfully deleted invoice!")
	return err
}

// UpdateInvoice updates database values from the row of the given invoice
func (ms *MySql) UpdateInvoice(invoice entities.Invoice) error { // update invoice
	_, err := db.Db.Exec("UPDATE nf_stn.invoices  SET referenceMonth=?, referenceYear=?, document=?, description=?, amount=?, isActive=?, createdAt=?, deactivatedAt=? WHERE id = ?;", invoice.ReferenceMonth, invoice.ReferenceYear, invoice.Document, invoice.Description, invoice.Amount, invoice.IsActive, invoice.CreatedAt, invoice.DeactivatedAt, invoice.ID)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfully updated invoice!")
	return err
}

// PatchInvoice partially updates database values from the row of the given invoice
func (ms *MySql) PatchInvoice(invoice entities.Invoice) error { // update invoice
	_, err := db.Db.Exec("UPDATE nf_stn.invoices  SET referenceMonth=?, referenceYear=?, description=?, amount=? WHERE id = ?;", invoice.ReferenceMonth, invoice.ReferenceYear, invoice.Description, invoice.Amount, invoice.ID)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfully patched invoice!")
	return err
}

// InvoiceExists checks if the given invoice document exists
func (ms *MySql) InvoiceExists(document string) (entities.Invoice, error) { // checks if invoice is already exists
	result, err := db.Db.Query("SELECT id FROM nf_stn.invoices WHERE document = ?", document)
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
func (ms *MySql) ClearTable() {
	_, _ = db.Db.Exec("TRUNCATE TABLE nf_stn.invoices")
}

// Pagination gets invoices list order by id and responds with 10 items limit
func (ms *MySql) Pagination(offset int) ([]entities.Invoice, error) {
	results, err := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices LIMIT 10 OFFSET ?;", offset)
	if err != nil {
		panic(err.Error())
	}
	inv := entities.Invoice{}
	invoicesList := []entities.Invoice{}
		for results.Next() {
		err := results.Scan(&inv.ID, &inv.ReferenceMonth, &inv.ReferenceYear, &inv.Document, &inv.Description, &inv.Amount, &inv.IsActive, &inv.CreatedAt, &inv.DeactivatedAt)
		if err != nil {
			panic(err.Error())
		}
		invoicesList = append(invoicesList, inv)
	}
	lines, _ := db.Db.Query("SELECT FOUND_ROWS();")
	var total int
	for lines.Next() {
		_ = lines.Scan(&total)
	}
	if offset > total {
		fmt.Println("Index out of range!")
	}
	fmt.Println("Successfully got invoice list!")
	return invoicesList, err
}

// Pagination gets invoices list order by id and responds with 10 items limit
func (ms *MySql) PaginationOrderByMonth(offset int) ([]entities.Invoice, error) {
	results, err := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices ORDER BY referenceMonth LIMIT 10 OFFSET ?;", offset)
	if err != nil {
		panic(err.Error())
	}
	inv := entities.Invoice{}
	invoicesList := []entities.Invoice{}
	for results.Next() {
		err := results.Scan(&inv.ID, &inv.ReferenceMonth, &inv.ReferenceYear, &inv.Document, &inv.Description, &inv.Amount, &inv.IsActive, &inv.CreatedAt, &inv.DeactivatedAt)
		if err != nil {
			panic(err.Error())
		}
		invoicesList = append(invoicesList, inv)
	}
	lines, _ := db.Db.Query("SELECT FOUND_ROWS();")
	var total int
	for lines.Next() {
		_ = lines.Scan(&total)
	}
	if offset > total {
		fmt.Println("Index out of range!")
	}
	fmt.Println("Successfully got invoice list!")
	return invoicesList, err
}

// Pagination gets invoices list order by id and responds with 10 items limit
func (ms *MySql) PaginationOrderByYear(offset int) ([]entities.Invoice, error) {
	results, err := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices ORDER BY referenceYear LIMIT 10 OFFSET ?;", offset)
	if err != nil {
		panic(err.Error())
	}
	inv := entities.Invoice{}
	invoicesList := []entities.Invoice{}
	for results.Next() {
		err := results.Scan(&inv.ID, &inv.ReferenceMonth, &inv.ReferenceYear, &inv.Document, &inv.Description, &inv.Amount, &inv.IsActive, &inv.CreatedAt, &inv.DeactivatedAt)
		if err != nil {
			panic(err.Error())
		}
		invoicesList = append(invoicesList, inv)
	}
	lines, _ := db.Db.Query("SELECT FOUND_ROWS();")
	var total int
	for lines.Next() {
		_ = lines.Scan(&total)
	}
	if offset > total {
		fmt.Println("Index out of range!")
	}
	fmt.Println("Successfully got invoice list!")
	return invoicesList, err
}

// Pagination gets invoices list order by id and responds with 10 items limit
func (ms *MySql) PaginationOrderByDocument(offset int) ([]entities.Invoice, error) {
	results, err := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices ORDER BY document LIMIT 10 OFFSET ?;", offset)
	if err != nil {
		panic(err.Error())
	}
	inv := entities.Invoice{}
	invoicesList := []entities.Invoice{}
	for results.Next() {
		err := results.Scan(&inv.ID, &inv.ReferenceMonth, &inv.ReferenceYear, &inv.Document, &inv.Description, &inv.Amount, &inv.IsActive, &inv.CreatedAt, &inv.DeactivatedAt)
		if err != nil {
			panic(err.Error())
		}
		invoicesList = append(invoicesList, inv)
	}
	lines, _ := db.Db.Query("SELECT FOUND_ROWS();")
	var total int
	for lines.Next() {
		_ = lines.Scan(&total)
	}
	if offset > total {
		fmt.Println("Index out of range!")
	}
	fmt.Println("Successfully got invoice list!")
	return invoicesList, err
}

//
func (ms *MySql) PaginationByMonth(offset, referenceMonth int) ([]entities.Invoice, error) {
	results, err := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices WHERE referenceMonth = ? LIMIT 10 OFFSET ?;", referenceMonth, offset)
	if err != nil {
		panic(err.Error())
	}
	inv := entities.Invoice{}
	invoicesList := []entities.Invoice{}
	for results.Next() {
		err := results.Scan(&inv.ID, &inv.ReferenceMonth, &inv.ReferenceYear, &inv.Document, &inv.Description, &inv.Amount, &inv.IsActive, &inv.CreatedAt, &inv.DeactivatedAt)
		if err != nil {
			panic(err.Error())
		}
		invoicesList = append(invoicesList, inv)
	}
	lines, _ := db.Db.Query("SELECT FOUND_ROWS();")
	var total int
	for lines.Next() {
		_ = lines.Scan(&total)
	}
	if offset > total {
		fmt.Println("Index out of range!")
	}
	fmt.Println("Successfully got invoice by month list!")
	return invoicesList, err
}

//
func (ms *MySql) PaginationByYear(offset, referenceYear int) ([]entities.Invoice, error) {
	results, err := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices WHERE referenceYear = ? LIMIT 10 OFFSET ?;", referenceYear, offset)
	if err != nil {
		panic(err.Error())
	}
	inv := entities.Invoice{}
	invoicesList := []entities.Invoice{}
	for results.Next() {
		err := results.Scan(&inv.ID, &inv.ReferenceMonth, &inv.ReferenceYear, &inv.Document, &inv.Description, &inv.Amount, &inv.IsActive, &inv.CreatedAt, &inv.DeactivatedAt)
		if err != nil {
			panic(err.Error())
		}
		invoicesList = append(invoicesList, inv)
	}
	lines, _ := db.Db.Query("SELECT FOUND_ROWS();")
	var total int
	for lines.Next() {
		_ = lines.Scan(&total)
	}
	if offset > total {
		fmt.Println("Index out of range!")
	}
	fmt.Println("Successfully got invoice by year list!")
	return invoicesList, err
}

//
func (ms *MySql) PaginationByDocument(offset int, document string) ([]entities.Invoice, error) {
	results, err := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices WHERE document = ? LIMIT 10 OFFSET ?;", document, offset)
	if err != nil {
		panic(err.Error())
	}
	inv := entities.Invoice{}
	invoicesList := []entities.Invoice{}
	for results.Next() {
		err := results.Scan(&inv.ID, &inv.ReferenceMonth, &inv.ReferenceYear, &inv.Document, &inv.Description, &inv.Amount, &inv.IsActive, &inv.CreatedAt, &inv.DeactivatedAt)
		if err != nil {
			panic(err.Error())
		}
		invoicesList = append(invoicesList, inv)
	}
	lines, _ := db.Db.Query("SELECT FOUND_ROWS();")
	var total int
	for lines.Next() {
		_ = lines.Scan(&total)
	}
	if offset > total {
		fmt.Println("Index out of range!")
	}
	fmt.Println("Successfully got invoice by year list!")
	return invoicesList, err
}