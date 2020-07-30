package database

import (
	"fmt"
	"nf_stn/config"
	"nf_stn/entities"
	"nf_stn/src"
)

//go:generate  go run github.com/golang/mock/mockgen  -package mock -destination=./mock/db_mock.go -source=$GOFILE

// DataBase interface
type DataBase interface {
	Init()
	GetAll() ([]entities.Invoice, error)
	GetInvoiceByDocument(document string) (entities.Invoice, error)
	GetUser(username, password string) (string, string, error)
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
	PaginationOrderByMonthYear(offset int) ([]entities.Invoice, error)
	PaginationOrderByMonthDocument(offset int) ([]entities.Invoice, error)
	PaginationOrderByYearDocument(offset int) ([]entities.Invoice, error)
}

// MySql struct
type MySql struct {
	Config config.DataBaseConfig
}

// db instantiation
var db config.App

// Init initializes db and redis connections
func (ms *MySql) Init()  {
	db.Initialize(ms.Config.DbDriver(), ms.Config.DbUser(), ms.Config.DbPass(), ms.Config.DbName())
}
// GetAll gets all the rows of the invoices db
func (ms *MySql) GetAll() ([]entities.Invoice, error) { // get list of all invoices
	results, err := db.Db.Query("SELECT id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices ORDER BY id")
	if err != nil {
		panic(err.Error())
	}
	inv := entities.Invoice{}
	var invoicesList []entities.Invoice
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
// GetUser gets the user credentials, if exists, by the given profile in the request
func (ms *MySql) GetUser(username, password string) (string, string, error) { // get acc data by profile
	result, err := db.Db.Query("SELECT username, password FROM nf_stn.users WHERE username = ?;", username)
	if err != nil {
		panic(err.Error())
	}
	var u entities.User
	for result.Next() {
		err = result.Scan(&u.Username, &u.Password)
		if err != nil {
			panic(err.Error())
		}
	}
	fmt.Println("Successfully got user!")
	return u.Username, u.Password, err
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
	//monthDay, month, hour, min, sec, year := time.Now().Day(), time.Now().Month(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Year()
	//date := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(monthDay)
	//clock := strconv.Itoa(hour) + ":" + strconv.Itoa(min) + ":" + strconv.Itoa(sec)
	//now := date + " " + clock
	now:=src.Now()
	_, err := db.Db.Query("INSERT INTO nf_stn.invoices (referenceMonth, referenceYear, document , description, amount, createdAt) VALUES (?, ?, ?, ?, ?, ?);",
		invoice.ReferenceMonth, invoice.ReferenceYear, invoice.Document, invoice.Description, invoice.Amount, now)
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
// Pagination returns page list of invoices ordered by id, 10 invoices per page
func (ms *MySql) Pagination(offset int) ([]entities.Invoice, error) {
	results, err := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices LIMIT 10 OFFSET ?;", offset)
	if err != nil {
		panic(err.Error())
	}
	inv := entities.Invoice{}
	var invoicesList []entities.Invoice
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
// PaginationByMonth returns page list of invoices ordered by month, 10 invoices per page
func (ms *MySql) PaginationByMonth(offset, referenceMonth int) ([]entities.Invoice, error) {
	results, err := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices WHERE referenceMonth = ? LIMIT 10 OFFSET ?;", referenceMonth, offset)
	if err != nil {
		panic(err.Error())
	}
	inv := entities.Invoice{}
	var invoicesList []entities.Invoice
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
// PaginationByYear returns page list of invoices ordered by year, 10 invoices per page
func (ms *MySql) PaginationByYear(offset, referenceYear int) ([]entities.Invoice, error) {
	results, err := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices WHERE referenceYear = ? LIMIT 10 OFFSET ?;", referenceYear, offset)
	if err != nil {
		panic(err.Error())
	}
	inv := entities.Invoice{}
	var invoicesList []entities.Invoice
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
// PaginationByDocument returns page list of invoices ordered by document, 10 invoices per page
func (ms *MySql) PaginationByDocument(offset int, document string) ([]entities.Invoice, error) {
	results, err := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices WHERE document = ? LIMIT 10 OFFSET ?;", document, offset)
	if err != nil {
		panic(err.Error())
	}
	inv := entities.Invoice{}
	var invoicesList []entities.Invoice
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
// PaginationOrderByMonth returns page list of invoices ordered by month, 10 invoices per page
func (ms *MySql) PaginationOrderByMonth(offset int) ([]entities.Invoice, error) {
	results, err := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices ORDER BY referenceMonth LIMIT 10 OFFSET ?;", offset)
	if err != nil {
		panic(err.Error())
	}
	inv := entities.Invoice{}
	var invoicesList []entities.Invoice
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
// PaginationOrderByYear returns page list of invoices ordered by year, 10 invoices per page
func (ms *MySql) PaginationOrderByYear(offset int) ([]entities.Invoice, error) {
	results, err := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices ORDER BY referenceYear LIMIT 10 OFFSET ?;", offset)
	if err != nil {
		panic(err.Error())
	}
	inv := entities.Invoice{}
	var invoicesList []entities.Invoice
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
// PaginationOrderByDocument returns page list of invoices ordered by document, 10 invoices per page
func (ms *MySql) PaginationOrderByDocument(offset int) ([]entities.Invoice, error) {
	results, err := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices ORDER BY document LIMIT 10 OFFSET ?;", offset)
	if err != nil {
		panic(err.Error())
	}
	inv := entities.Invoice{}
	var invoicesList []entities.Invoice
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
// PaginationOrderByMonthYear returns page list of invoices ordered by month and year, 10 invoices per page
func (ms *MySql) PaginationOrderByMonthYear(offset int) ([]entities.Invoice, error) {
	results, err := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices ORDER BY referenceMonth, referenceYear LIMIT 10 OFFSET ?;", offset)
	if err != nil {
		panic(err.Error())
	}
	inv := entities.Invoice{}
	var invoicesList []entities.Invoice
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
// PaginationOrderByMonthDocument returns page list of invoices ordered by month and document, 10 invoices per page
func (ms *MySql) PaginationOrderByMonthDocument(offset int) ([]entities.Invoice, error) {
	results, err := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices ORDER BY referenceMonth, document LIMIT 10 OFFSET ?;", offset)
	if err != nil {
		panic(err.Error())
	}
	inv := entities.Invoice{}
	var invoicesList []entities.Invoice
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
// PaginationOrderByYearDocument returns page list of invoices ordered by year and document, 10 invoices per page
func (ms *MySql) PaginationOrderByYearDocument(offset int) ([]entities.Invoice, error) {
	results, err := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices ORDER BY document, referenceYear LIMIT 10 OFFSET ?;", offset)
	if err != nil {
		panic(err.Error())
	}
	inv := entities.Invoice{}
	var invoicesList []entities.Invoice
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