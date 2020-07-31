package adapter

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"nf_stn/authentication"
	"nf_stn/database"
	"nf_stn/entities"
	//"nf_stn/src"
	"strconv"
	//"github.com/golang/mock"
)

// Routes struct
type Routes struct {
	Db database.DataBase
	Au authentication.Token
}

// GetAll gets all invoices and returns in json format
func (rr *Routes) GetAll(w http.ResponseWriter, r *http.Request) { // get invoices and returns in json format
	var results []entities.Invoice
	results, err := rr.Db.GetAll()
	if err != nil {
		panic(err.Error())
	} else {
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(results)
}
// GetInvoiceByDocument gets invoice by document value and returns the invoice in json format
func (rr *Routes) GetInvoiceByDocument(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	var invoice entities.Invoice // Unmarshal
	err = json.Unmarshal(b, &invoice)
	var result entities.Invoice
	result, err = rr.Db.GetInvoiceByDocument(invoice.Document)
	if err != nil {
		panic(err.Error())
	} else {
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(result)
}
// InsertInvoice inserts the given invoice data in the rr.Db
func (rr *Routes) InsertInvoice(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	var invoice entities.Invoice // Unmarshal
	err = json.Unmarshal(b, &invoice)
	if err != nil {
		panic(err.Error())
	}
	err = rr.Db.InsertInvoice(invoice)
	w.WriteHeader(http.StatusCreated)
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(invoice)
}
// DeleteInvoice makes the logic deletion of the invoice by the given ID in the rr.Db
func (rr *Routes) DeleteInvoice(w http.ResponseWriter, r *http.Request) {
	ID , _ := strconv.Atoi(r.FormValue("ID"))
	idNotInList, err := rr.Db.GetInvoiceByID(ID)
	if err != nil {
		panic(err.Error())
	}
	if (idNotInList == entities.Invoice{}) {
		println("id not found!")
		w.WriteHeader(http.StatusNotFound)
	} else {
		err := rr.Db.DeleteInvoice(ID)
		if err != nil {
			panic(err.Error())
		}
		w.WriteHeader(http.StatusOK)
	}
	deletedID := map[string]int{
		"ID": idNotInList.ID,
	}
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(deletedID)
}
// UpdateInvoice updates database values from the row of the given invoice
func (rr *Routes) UpdateInvoice(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	var invoice entities.Invoice
	err = json.Unmarshal(b, &invoice) // Unmarshal
	if err != nil {
		panic(err.Error())
	}
	invoiceExists, _ := rr.Db.InvoiceExists(invoice.Document)
	if (invoiceExists == entities.Invoice{}) {
		println("Invoice not found!")
		w.WriteHeader(http.StatusNotFound)
	} else {
		invoice.ID = invoiceExists.ID
		err := rr.Db.UpdateInvoice(invoice)
		if err != nil {
			panic(err.Error())
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(invoice)
}
// PatchInvoice partially updates rr.Db values from the row of the given invoice
func (rr *Routes) PatchInvoice(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	var invoice, editedInvoice entities.Invoice
	_ = json.Unmarshal(b, &invoice) // Unmarshal
	err = rr.Db.PatchInvoice(invoice)
	if err != nil {
		panic(err.Error())
	}
	editedInvoice, err = rr.Db.GetInvoiceByID(invoice.ID)
	if err != nil {
		panic(err.Error())
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(editedInvoice)
}
// Pagination gets page list of invoices ordered by id, 10 invoices per page
func (rr *Routes) Pagination(w http.ResponseWriter, r *http.Request) { // get invoices and returns in json format
	//params := mux.Vars(r)
	//page, err := strconv.Atoi(params["page"])
	page , _ := strconv.Atoi(r.FormValue("page"))
	results, err := rr.Db.Pagination((page-1)*10)
	if err != nil {
		panic(err.Error())
	} else {
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(results)
}
// PaginationByMonth gets page list of invoices by month, 10 invoices per page
func (rr *Routes) PaginationByMonth(w http.ResponseWriter, r *http.Request) { // get invoices and returns in json format
	params := mux.Vars(r)
	page, err := strconv.Atoi(params["page"])
	referenceMonth, err := strconv.Atoi(params["referenceMonth"])
	results, err := rr.Db.PaginationByMonth((page-1)*10,referenceMonth)
	if err != nil {
		panic(err.Error())
	} else {
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(results)
}
// PaginationByYear gets page list of invoices by year, 10 invoices per page
func (rr *Routes) PaginationByYear(w http.ResponseWriter, r *http.Request) { // get invoices and returns in json format
	params := mux.Vars(r)
	page, err := strconv.Atoi(params["page"])
	referenceYear, err := strconv.Atoi(params["referenceYear"])
	results, err := rr.Db.PaginationByYear((page-1)*10,referenceYear)
	if err != nil {
		panic(err.Error())
	} else {
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(results)
}
// PaginationByDocument gets page list of invoices by document, 10 invoices per page
func (rr *Routes) PaginationByDocument(w http.ResponseWriter, r *http.Request) { // get invoices and returns in json format
	params := mux.Vars(r)
	page, err := strconv.Atoi(params["page"])
	document := params["document"]
	results, err := rr.Db.PaginationByDocument((page-1)*10,document)
	if err != nil {
		panic(err.Error())
	} else {
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(results)
}
// PaginationOrderByMonth gets page list of invoices ordered by month, 10 invoices per page
func (rr *Routes) PaginationOrderByMonth(w http.ResponseWriter, r *http.Request) { // get invoices and returns in json format
	params := mux.Vars(r)
	page, err := strconv.Atoi(params["page"])
	results, err := rr.Db.PaginationOrderByMonth((page-1)*10)
	if err != nil {
		panic(err.Error())
	} else {
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(results)
}
// PaginationOrderByYear gets page list of invoices ordered by year, 10 invoices per page
func (rr *Routes) PaginationOrderByYear(w http.ResponseWriter, r *http.Request) { // get invoices and returns in json format
	params := mux.Vars(r)
	page, err := strconv.Atoi(params["page"])
	results, err := rr.Db.PaginationOrderByYear((page-1)*10)
	if err != nil {
		panic(err.Error())
	} else {
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(results)
}
// PaginationOrderByDocument gets page list of invoices ordered by document, 10 invoices per page
func (rr *Routes) PaginationOrderByDocument(w http.ResponseWriter, r *http.Request) { // get invoices and returns in json format
	params := mux.Vars(r)
	page, err := strconv.Atoi(params["page"])
	results, err := rr.Db.PaginationOrderByDocument((page-1)*10)
	if err != nil {
		panic(err.Error())
	} else {
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(results)
}
// PaginationOrderByMonthYear gets page list of invoices ordered by month and year, 10 invoices per page
func (rr *Routes) PaginationOrderByMonthYear(w http.ResponseWriter, r *http.Request) { // get invoices and returns in json format
	params := mux.Vars(r)
	page, err := strconv.Atoi(params["page"])
	results, err := rr.Db.PaginationOrderByMonthYear((page-1)*10)
	if err != nil {
		panic(err.Error())
	} else {
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(results)
}
// PaginationOrderByMonthDocument gets page list of invoices ordered by month and document, 10 invoices per page
func (rr *Routes) PaginationOrderByMonthDocument(w http.ResponseWriter, r *http.Request) { // get invoices and returns in json format
	params := mux.Vars(r)
	page, err := strconv.Atoi(params["page"])
	results, err := rr.Db.PaginationOrderByMonthDocument((page-1)*10)
	if err != nil {
		panic(err.Error())
	} else {
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(results)
}
// PaginationOrderByYearDocument gets page list of invoices ordered by year and document, 10 invoices per page
func (rr *Routes) PaginationOrderByYearDocument(w http.ResponseWriter, r *http.Request) { // get invoices and returns in json format
	params := mux.Vars(r)
	page, err := strconv.Atoi(params["page"])
	results, err := rr.Db.PaginationOrderByYearDocument((page-1)*10)
	if err != nil {
		panic(err.Error())
	} else {
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(results)
}
// GenerateToken generates token for authenticated user
func (rr *Routes) GenerateToken(w http.ResponseWriter, r *http.Request) {
	var u entities.User
	var user entities.User
	var err error
	u.Username = r.Header.Get("username")
	u.Password = r.Header.Get("password")
	user.ID,user.Username, user.Password, err = rr.Db.GetUser(u.Username,u.Password)
	if err != nil {
		panic(err)
		return
	}
	//compare the user from the request, with the one defined in database:
	if user.Username != u.Username || user.Password != u.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	token, err := rr.Au.CreateToken(uint64(user.ID), user.Username)
	if err != nil {
		panic(err)
		return
	}
	saveErr := rr.Au.CreateAuth(uint64(user.ID),token)
	if saveErr != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Println(saveErr)
	}
	tokens := map[string]string{
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(tokens)
}
//
