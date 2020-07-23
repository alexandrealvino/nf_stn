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

var db config.App

func init() {
	db.Initialize(config.Dbdriver,config.Dbuser,config.Dbpass,config.Dbname)
}

//func GetAll() ([]entities.Invoice, error){  // get list of all invoices
//	db := config.DbConn()
//	defer db.Close()
//	results, err := db.Query("SELECT id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM invoices ORDER BY id ASC")
//	if err != nil {
//		panic(err.Error())
//	}
//	inv := entities.Invoice{}
//	invoicesList := []entities.Invoice{}
//	for results.Next() {
//		err = results.Scan(&inv.ID, &inv.ReferenceMonth, &inv.ReferenceYear, &inv.Document, &inv.Description, &inv.Amount, &inv.IsActive, &inv.CreatedAt, &inv.DeactivatedAt)
//		if err != nil {
//			panic(err.Error())
//		}
//		invoicesList = append(invoicesList, inv)
//	}
//	fmt.Println("Successfuly got invoice list!")
//	return invoicesList, err
//}
////
//func GetInvoiceById(id int) (entities.Invoice, error){  // get ticker by id
//	db := config.DbConn()
//	defer db.Close()
//	result, err := db.Query("SELECT id, referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt FROM invoices WHERE id = ?", id)
//	if err != nil {
//		panic(err.Error())
//	}
//	inv := entities.Invoice{}
//	for result.Next() {
//		err = result.Scan(&inv.ID, &inv.ReferenceMonth, &inv.ReferenceYear, &inv.Document, &inv.Description, &inv.Amount, &inv.IsActive, &inv.CreatedAt, &inv.DeactivatedAt)
//		if err != nil {
//			panic(err.Error())
//		}
//	}
//	if (inv!= entities.Invoice{}) {
//		fmt.Println("Successfuly got invoice!")
//	}
//	return inv, err
//}
////
//func InsertInvoice(invoice entities.Invoice) error {  // insert invoice
//	db := config.DbConn()
//	defer db.Close()
//	_, err := db.Query("INSERT INTO invoices (ReferenceMonth, ReferenceYear, Document, Description, Amount, IsActive, CreatedAt, DeactivatedAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?);",
//		invoice.ReferenceMonth, invoice.ReferenceYear, invoice.Document, invoice.Description, invoice.Amount, invoice.IsActive, invoice.CreatedAt, invoice.DeactivatedAt)
//	if err != nil {
//		panic(err.Error())
//	}
//	fmt.Println("Successfuly inserted invoice!")
//	return err
//}
////
//func DeleteInvoice(id int) (error) { // delete invoice
//	db := config.DbConn()
//	defer db.Close()
//	_, err := db.Exec("DELETE FROM invoices WHERE id = ?;", id)
//	if err != nil {
//		panic(err.Error())
//	}
//	fmt.Println("Successfuly deleted invoice!")
//	return err
//}
////
//func UpdateInvoice(invoice entities.Invoice) (error) {  // update invoice
//	db := config.DbConn()
//	defer db.Close()
//	_, err := db.Exec("UPDATE invoices  SET ReferenceMonth=?, ReferenceYear=?, Document=?, Description=?, Amount=?, IsActive=?, CreatedAt=?, DeactivatedAt=? WHERE id = ?;", invoice.ReferenceMonth, invoice.ReferenceYear, invoice.Document, invoice.Description, invoice.Amount, invoice.IsActive, invoice.CreatedAt, invoice.DeactivatedAt, invoice.ID)
//	if err != nil {
//		panic(err.Error())
//	}
//	fmt.Println("Successfuly updated invoice!")
//	return err
//}
////
//func InvoiceExists(document string) (entities.Invoice, error){  // checks if invoice is already exists
//	db := config.DbConn()
//	defer db.Close()
//	result, err := db.Query("SELECT id FROM invoices WHERE document = ?", document)
//	if err != nil {
//		panic(err.Error())
//	}
//	invoice := entities.Invoice{}
//	for result.Next() {
//		err = result.Scan(&invoice.ID)
//		if err != nil {
//			panic(err.Error())
//		}
//	}
//	if (invoice != entities.Invoice{}) {
//		fmt.Println("Invoice exists!")
//	}
//	//else {
//	//	fmt.Println("Invoice not found!")
//	//}
//	return invoice, err
//}
////
//func ClearTable() {
//	db := config.DbConn()
//	defer db.Close()
//	_,_ = db.Exec("TRUNCATE TABLE invoices")
//}
////

func GetAll() ([]entities.Invoice, error){  // get list of all invoices
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
//
func GetInvoiceByDocument(document string) (entities.Invoice, error){  // get invoice by document
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
//
func GetHashByProfile(profile string) ([]byte, error){  // get invoice by profile
	result, err := db.Db.Query("SELECT hash FROM hashes WHERE profile = ?;", profile)
	if err != nil {
		panic(err.Error())
	}
	var hash []byte
	for result.Next() {
		err = result.Scan(&hash)
		if err != nil {
			panic(err.Error())
		}
	}
	fmt.Println("Successfuly got hash!")
	return hash, err
}
//
func GetInvoiceById(id int) (entities.Invoice, error){  // get ticker by id
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
	if (inv!= entities.Invoice{}) {
		fmt.Println("Successfuly got invoice!")
	}
	return inv, err
}
//
func InsertInvoice(invoice entities.Invoice) error {  // insert invoice
	_, err := db.Db.Query("INSERT INTO invoices (referenceMonth, referenceYear, document , description, amount, isActive, createdAt, deactivatedAt) VALUES (?, ?, ?, ?, ?, ?, ?, ?);",
		invoice.ReferenceMonth, invoice.ReferenceYear, invoice.Document, invoice.Description, invoice.Amount, invoice.IsActive, invoice.CreatedAt, invoice.DeactivatedAt)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfuly inserted invoice!")
	return err
}
//
func DeleteInvoice(id int) (error) { // set isActive = 0 for logic deletion
	_, err := db.Db.Exec("UPDATE invoices SET isActive = ? WHERE id = ?;",0, id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfuly deleted invoice!")
	return err
}
//
func UpdateInvoice(invoice entities.Invoice) (error) {  // update invoice
	_, err := db.Db.Exec("UPDATE invoices  SET referenceMonth=?, referenceYear=?, document=?, description=?, amount=?, isActive=?, createdAt=?, deactivatedAt=? WHERE id = ?;", invoice.ReferenceMonth, invoice.ReferenceYear, invoice.Document, invoice.Description, invoice.Amount, invoice.IsActive, invoice.CreatedAt, invoice.DeactivatedAt, invoice.ID)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfuly updated invoice!")
	return err
}
//
func PatchInvoice(invoice entities.Invoice) (error) {  // update invoice
	_, err := db.Db.Exec("UPDATE invoices  SET referenceMonth=?, referenceYear=?, description=?, amount=? WHERE id = ?;", invoice.ReferenceMonth, invoice.ReferenceYear, invoice.Description, invoice.Amount, invoice.ID)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Successfuly updated invoice!")
	return err
}
//
func InvoiceExists(document string) (entities.Invoice, error){  // checks if invoice is already exists
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
//
func ClearTable() {
	_,_ = db.Db.Exec("TRUNCATE TABLE invoices")
}
//