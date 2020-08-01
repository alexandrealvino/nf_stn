package database

import (
    log "github.com/sirupsen/logrus"
	"nf_stn/config"
	"nf_stn/entities"
	"nf_stn/lib"
)

//go:generate  go run github.com/golang/mock/mockgen  -package mock -destination=./mock/db_mock.go -source=$GOFILE

// DataBase interface
type DataBase interface {
	Init()
	GetAll() ([]entities.Invoice, error)
	GetInvoiceByDocument(document string) (entities.Invoice, error)
	GetUser(username string) (int, string, string, error)
	GetInvoiceByID(id int) (entities.Invoice, error)
	InsertInvoice(invoice entities.Invoice) error
	DeleteInvoice(id int) error
	UpdateInvoice(invoice entities.Invoice) error
	PatchInvoice(invoice entities.Invoice) error
	InvoiceExists(document string) (entities.Invoice, error)
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

	PaginationTEST(query string, referenceMonth int) ([]entities.Invoice, error)
}

// MySQL struct
type MySQL struct {
	Config config.DataBaseConfig
}

// db instantiation
var db config.App

// Init initializes db connection
func (ms *MySQL) Init()  {
	//db.Initialize(ms.Config.DbDriver(), ms.Config.DbUser(), ms.Config.DbPass(), ms.Config.DbName())
	db.Initialize(ms.Config.DbDriver(), ms.Config.Conn())
	log.Info("db connected")
}
// GetAll gets all the rows of the invoices db
func (ms *MySQL) GetAll() ([]entities.Invoice, error) { // get list of all invoices
	results, err := db.Db.Query("SELECT id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices ORDER BY id")
	if err != nil {
		log.Error("null db values")
		return nil, err
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
	log.Println("successfully got invoice list!")
	return invoicesList, err
}
// GetInvoiceByDocument gets the invoice by the document value
func (ms *MySQL) GetInvoiceByDocument(document string) (entities.Invoice, error) { // get invoice by document
	inv := entities.Invoice{}
	result, err := db.Db.Query("SELECT id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices WHERE document = ?;", document)
	if err != nil {
		log.Println("db error")
		return inv, err
	}
	for result.Next() {
		err = result.Scan(&inv.ID, &inv.ReferenceMonth, &inv.ReferenceYear, &inv.Document, &inv.Description, &inv.Amount, &inv.IsActive, &inv.CreatedAt, &inv.DeactivatedAt)
		if err != nil {
			log.Println("db error")
		}
	}
	log.Println("successfully got invoice!")
	return inv, err
}
// GetUser gets the user credentials, if exists, by the given profile in the request
func (ms *MySQL) GetUser(username string) (int,string, string, error) { // get acc data by profile
	var u entities.User
	result, err := db.Db.Query("SELECT id,username, hash FROM nf_stn.users WHERE username = ?;", username)
	if err != nil {
		log.Println("db error")
		return u.ID,u.Username,u.Hash, err
	}
	for result.Next() {
		err = result.Scan(&u.ID,&u.Username, &u.Hash)
		if err != nil {
			log.Println("db error")
		}
	}
	log.Println("successfully got user!")
	return u.ID,u.Username, u.Hash, err
}
// GetInvoiceByID gets the invoice by the given ID
func (ms *MySQL) GetInvoiceByID(id int) (entities.Invoice, error) { // get ticker by id
	inv := entities.Invoice{}
	result, err := db.Db.Query("SELECT id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices WHERE id = ?", id)
	if err != nil {
		log.Println("db error")
		return inv, err
	}
	for result.Next() {
		err = result.Scan(&inv.ID, &inv.ReferenceMonth, &inv.ReferenceYear, &inv.Document, &inv.Description, &inv.Amount, &inv.IsActive, &inv.CreatedAt, &inv.DeactivatedAt)
		if err != nil {
			log.Println("db error")
		}
	}
	if (inv != entities.Invoice{}) {
		log.Println("successfully got invoice!")
	}
	return inv, err
}
// InsertInvoice inserts the given invoice to invoices db
func (ms *MySQL) InsertInvoice(invoice entities.Invoice) error { // insert invoice
	now:= lib.Now()
	_, err := db.Db.Query("INSERT INTO nf_stn.invoices (referenceMonth, referenceYear, document , description, amount, createdAt) VALUES (?, ?, ?, ?, ?, ?);",
		invoice.ReferenceMonth, invoice.ReferenceYear, invoice.Document, invoice.Description, invoice.Amount, now)
	if err != nil {
		log.Println("db error")
	}
	log.Println("successfully inserted invoice!")
	return err
}
// DeleteInvoice makes the logic deletion setting isActive = 0 by the given ID
func (ms *MySQL) DeleteInvoice(id int) error { // set isActive = 0 for logic deletion
	_, err := db.Db.Exec("UPDATE nf_stn.invoices SET isActive = ? WHERE id = ?;", 0, id)
	if err != nil {
		log.Println("db error")
	}
	log.Println("successfully deleted invoice!")
	return err
}
// UpdateInvoice updates database values from the row of the given invoice
func (ms *MySQL) UpdateInvoice(invoice entities.Invoice) error { // update invoice
	_, err := db.Db.Exec("UPDATE nf_stn.invoices  SET referenceMonth=?, referenceYear=?, document=?, description=?, amount=?, isActive=?, createdAt=?, deactivatedAt=? WHERE id = ?;", invoice.ReferenceMonth, invoice.ReferenceYear, invoice.Document, invoice.Description, invoice.Amount, invoice.IsActive, invoice.CreatedAt, invoice.DeactivatedAt, invoice.ID)
	if err != nil {
		log.Println("db error")
	}
	log.Println("successfully updated invoice!")
	return err
}
// PatchInvoice partially updates database values from the row of the given invoice
func (ms *MySQL) PatchInvoice(invoice entities.Invoice) error { // update invoice
	_, err := db.Db.Exec("UPDATE nf_stn.invoices  SET referenceMonth=?, referenceYear=?, description=?, amount=? WHERE id = ?;", invoice.ReferenceMonth, invoice.ReferenceYear, invoice.Description, invoice.Amount, invoice.ID)
	if err != nil {
		log.Println("db error")
	}
	log.Println("successfully patched invoice!")
	return err
}
// InvoiceExists checks if the given invoice document exists
func (ms *MySQL) InvoiceExists(document string) (entities.Invoice, error) { // checks if invoice is already exists
	invoice := entities.Invoice{}
	result, err := db.Db.Query("SELECT id FROM nf_stn.invoices WHERE document = ?", document)
	if err != nil {
		log.Println("db error")
		return invoice, err
	}
	for result.Next() {
		err = result.Scan(&invoice.ID)
		if err != nil {
			log.Println("db error")
		}
	}
	if (invoice != entities.Invoice{}) {
		log.Println("Invoice exists!")
	}
	return invoice, err
}
// Pagination returns page list of invoices ordered by id, 10 invoices per page
func (ms *MySQL) Pagination(offset int) ([]entities.Invoice, error) {
	results, err := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices LIMIT 10 OFFSET ?;", offset)
	if err != nil {
		log.Println("db error")
		return nil, err
	}
	inv := entities.Invoice{}
	var invoicesList []entities.Invoice
		for results.Next() {
		err := results.Scan(&inv.ID, &inv.ReferenceMonth, &inv.ReferenceYear, &inv.Document, &inv.Description, &inv.Amount, &inv.IsActive, &inv.CreatedAt, &inv.DeactivatedAt)
		if err != nil {
			log.Println("db error")
		}
		invoicesList = append(invoicesList, inv)
	}
	lines, _ := db.Db.Query("SELECT FOUND_ROWS();")
	var total int
	for lines.Next() {
		_ = lines.Scan(&total)
	}
	if offset > total {
		log.Println("Index out of range!")
	}
	log.Println("successfully got invoice list!")
	return invoicesList, err
}
// PaginationByMonth returns page list of invoices ordered by month, 10 invoices per page
func (ms *MySQL) PaginationByMonth(offset, referenceMonth int) ([]entities.Invoice, error) {
	results, err := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices WHERE referenceMonth = ? LIMIT 10 OFFSET ?;", referenceMonth, offset)
	if err != nil {
		log.Println("db error")
		return nil, err
	}
	inv := entities.Invoice{}
	var invoicesList []entities.Invoice
	for results.Next() {
		err := results.Scan(&inv.ID, &inv.ReferenceMonth, &inv.ReferenceYear, &inv.Document, &inv.Description, &inv.Amount, &inv.IsActive, &inv.CreatedAt, &inv.DeactivatedAt)
		if err != nil {
			log.Println("db error")
		}
		invoicesList = append(invoicesList, inv)
	}
	lines, _ := db.Db.Query("SELECT FOUND_ROWS();")
	var total int
	for lines.Next() {
		_ = lines.Scan(&total)
	}
	if offset > total {
		log.Println("Index out of range!")
	}
	log.Println("successfully got invoice by month list!")
	return invoicesList, err
}
// PaginationByYear returns page list of invoices ordered by year, 10 invoices per page
func (ms *MySQL) PaginationByYear(offset, referenceYear int) ([]entities.Invoice, error) {
	results, err := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices WHERE referenceYear = ? LIMIT 10 OFFSET ?;", referenceYear, offset)
	if err != nil {
		log.Println("db error")
		return nil, err
	}
	inv := entities.Invoice{}
	var invoicesList []entities.Invoice
	for results.Next() {
		err := results.Scan(&inv.ID, &inv.ReferenceMonth, &inv.ReferenceYear, &inv.Document, &inv.Description, &inv.Amount, &inv.IsActive, &inv.CreatedAt, &inv.DeactivatedAt)
		if err != nil {
			log.Println("db error")
		}
		invoicesList = append(invoicesList, inv)
	}
	lines, _ := db.Db.Query("SELECT FOUND_ROWS();")
	var total int
	for lines.Next() {
		_ = lines.Scan(&total)
	}
	if offset > total {
		log.Println("Index out of range!")
	}
	log.Println("successfully got invoice by year list!")
	return invoicesList, err
}
// PaginationByDocument returns page list of invoices ordered by document, 10 invoices per page
func (ms *MySQL) PaginationByDocument(offset int, document string) ([]entities.Invoice, error) {
	results, err := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices WHERE document = ? LIMIT 10 OFFSET ?;", document, offset)
	if err != nil {
		log.Println("db error")
		return nil, err
	}
	inv := entities.Invoice{}
	var invoicesList []entities.Invoice
	for results.Next() {
		err := results.Scan(&inv.ID, &inv.ReferenceMonth, &inv.ReferenceYear, &inv.Document, &inv.Description, &inv.Amount, &inv.IsActive, &inv.CreatedAt, &inv.DeactivatedAt)
		if err != nil {
			log.Println("db error")
		}
		invoicesList = append(invoicesList, inv)
	}
	lines, _ := db.Db.Query("SELECT FOUND_ROWS();")
	var total int
	for lines.Next() {
		_ = lines.Scan(&total)
	}
	if offset > total {
		log.Println("Index out of range!")
	}
	log.Println("successfully got invoice by year list!")
	return invoicesList, err
}
// PaginationOrderByMonth returns page list of invoices ordered by month, 10 invoices per page
func (ms *MySQL) PaginationOrderByMonth(offset int) ([]entities.Invoice, error) {
	results, err := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices ORDER BY referenceMonth LIMIT 10 OFFSET ?;", offset)
	if err != nil {
		log.Println("db error")
		return nil, err
	}
	inv := entities.Invoice{}
	var invoicesList []entities.Invoice
	for results.Next() {
		err := results.Scan(&inv.ID, &inv.ReferenceMonth, &inv.ReferenceYear, &inv.Document, &inv.Description, &inv.Amount, &inv.IsActive, &inv.CreatedAt, &inv.DeactivatedAt)
		if err != nil {
			log.Println("db error")
		}
		invoicesList = append(invoicesList, inv)
	}
	lines, _ := db.Db.Query("SELECT FOUND_ROWS();")
	var total int
	for lines.Next() {
		_ = lines.Scan(&total)
	}
	if offset > total {
		log.Println("Index out of range!")
	}
	log.Println("successfully got invoice list!")
	return invoicesList, err
}
// PaginationOrderByYear returns page list of invoices ordered by year, 10 invoices per page
func (ms *MySQL) PaginationOrderByYear(offset int) ([]entities.Invoice, error) {
	results, err := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices ORDER BY referenceYear LIMIT 10 OFFSET ?;", offset)
	if err != nil {
		log.Println("db error")
		return nil, err
	}
	inv := entities.Invoice{}
	var invoicesList []entities.Invoice
	for results.Next() {
		err := results.Scan(&inv.ID, &inv.ReferenceMonth, &inv.ReferenceYear, &inv.Document, &inv.Description, &inv.Amount, &inv.IsActive, &inv.CreatedAt, &inv.DeactivatedAt)
		if err != nil {
			log.Println("db error")
		}
		invoicesList = append(invoicesList, inv)
	}
	lines, _ := db.Db.Query("SELECT FOUND_ROWS();")
	var total int
	for lines.Next() {
		_ = lines.Scan(&total)
	}
	if offset > total {
		log.Println("Index out of range!")
	}
	log.Println("successfully got invoice list!")
	return invoicesList, err
}
// PaginationOrderByDocument returns page list of invoices ordered by document, 10 invoices per page
func (ms *MySQL) PaginationOrderByDocument(offset int) ([]entities.Invoice, error) {
	results, err := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices ORDER BY document LIMIT 10 OFFSET ?;", offset)
	if err != nil {
		log.Println("db error")
		return nil, err
	}
	inv := entities.Invoice{}
	var invoicesList []entities.Invoice
	for results.Next() {
		err := results.Scan(&inv.ID, &inv.ReferenceMonth, &inv.ReferenceYear, &inv.Document, &inv.Description, &inv.Amount, &inv.IsActive, &inv.CreatedAt, &inv.DeactivatedAt)
		if err != nil {
			log.Println("db error")
		}
		invoicesList = append(invoicesList, inv)
	}
	lines, _ := db.Db.Query("SELECT FOUND_ROWS();")
	var total int
	for lines.Next() {
		_ = lines.Scan(&total)
	}
	if offset > total {
		log.Println("Index out of range!")
	}
	log.Println("successfully got invoice list!")
	return invoicesList, err
}
// PaginationOrderByMonthYear returns page list of invoices ordered by month and year, 10 invoices per page
func (ms *MySQL) PaginationOrderByMonthYear(offset int) ([]entities.Invoice, error) {
	results, err := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices ORDER BY referenceMonth, referenceYear LIMIT 10 OFFSET ?;", offset)
	if err != nil {
		log.Println("db error")
		return nil, err
	}
	inv := entities.Invoice{}
	var invoicesList []entities.Invoice
	for results.Next() {
		err := results.Scan(&inv.ID, &inv.ReferenceMonth, &inv.ReferenceYear, &inv.Document, &inv.Description, &inv.Amount, &inv.IsActive, &inv.CreatedAt, &inv.DeactivatedAt)
		if err != nil {
			log.Println("db error")
		}
		invoicesList = append(invoicesList, inv)
	}
	lines, _ := db.Db.Query("SELECT FOUND_ROWS();")
	var total int
	for lines.Next() {
		_ = lines.Scan(&total)
	}
	if offset > total {
		log.Println("Index out of range!")
	}
	log.Println("successfully got invoice list!")
	return invoicesList, err
}
// PaginationOrderByMonthDocument returns page list of invoices ordered by month and document, 10 invoices per page
func (ms *MySQL) PaginationOrderByMonthDocument(offset int) ([]entities.Invoice, error) {
	results, err := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices ORDER BY referenceMonth, document LIMIT 10 OFFSET ?;", offset)
	if err != nil {
		log.Println("db error")
		return nil, err
	}
	inv := entities.Invoice{}
	var invoicesList []entities.Invoice
	for results.Next() {
		err := results.Scan(&inv.ID, &inv.ReferenceMonth, &inv.ReferenceYear, &inv.Document, &inv.Description, &inv.Amount, &inv.IsActive, &inv.CreatedAt, &inv.DeactivatedAt)
		if err != nil {
			log.Println("db error")
		}
		invoicesList = append(invoicesList, inv)
	}
	lines, _ := db.Db.Query("SELECT FOUND_ROWS();")
	var total int
	for lines.Next() {
		_ = lines.Scan(&total)
	}
	if offset > total {
		log.Println("Index out of range!")
	}
	log.Println("successfully got invoice list!")
	return invoicesList, err
}
// PaginationOrderByYearDocument returns page list of invoices ordered by year and document, 10 invoices per page
func (ms *MySQL) PaginationOrderByYearDocument(offset int) ([]entities.Invoice, error) {
	results, err := db.Db.Query("SELECT SQL_CALC_FOUND_ROWS id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM nf_stn.invoices ORDER BY document, referenceYear LIMIT 10 OFFSET ?;", offset)
	if err != nil {
		log.Println("db error")
		return nil, err
	}
	inv := entities.Invoice{}
	var invoicesList []entities.Invoice
	for results.Next() {
		err := results.Scan(&inv.ID, &inv.ReferenceMonth, &inv.ReferenceYear, &inv.Document, &inv.Description, &inv.Amount, &inv.IsActive, &inv.CreatedAt, &inv.DeactivatedAt)
		if err != nil {
			log.Println("db error")
		}
		invoicesList = append(invoicesList, inv)
	}
	lines, _ := db.Db.Query("SELECT FOUND_ROWS();")
	var total int
	for lines.Next() {
		_ = lines.Scan(&total)
	}
	if offset > total {
		log.Println("Index out of range!")
	}
	log.Println("successfully got invoice list!")
	return invoicesList, err
}
//

// PaginationByMonth returns page list of invoices ordered by month, 10 invoices per page
func (ms *MySQL) PaginationTEST(query string, referenceMonth int) ([]entities.Invoice, error) {
	results, err := db.Db.Query(query, referenceMonth)
	if err != nil {
		log.Println("db error")
		return nil, err
	}
	inv := entities.Invoice{}
	var invoicesList []entities.Invoice
	for results.Next() {
		err := results.Scan(&inv.ID, &inv.ReferenceMonth, &inv.ReferenceYear, &inv.Document, &inv.Description, &inv.Amount, &inv.IsActive, &inv.CreatedAt, &inv.DeactivatedAt)
		if err != nil {
			log.Println("db error")
		}
		invoicesList = append(invoicesList, inv)
	}
	lines, _ := db.Db.Query("SELECT FOUND_ROWS();")
	var total int
	for lines.Next() {
		_ = lines.Scan(&total)
	}
	//if offset > total {
	//	log.Println("Index out of range!")
	//}
	log.Println("successfully got invoice by month list!")
	return invoicesList, err
}