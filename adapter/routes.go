package adapter

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"nf_stn/database"
	"nf_stn/entities"
	"nf_stn/src"
	"strconv"
	//"github.com/golang/mock"
)




type Routes struct {
	Db database.DataBase
}

// GetIndexT gets the list of all invoices in the database
// and renders the index.html with a table containing
// the results from the GET method
func (rr *Routes) GetIndexT(w http.ResponseWriter, r *http.Request) { // render index.html
	var results []entities.Invoice
	//var profile string
	results, err := rr.Db.GetAll()
	if err != nil {
		panic(err.Error())
	}
	tmpl, _ := template.ParseFiles("static/index.html")
	err = tmpl.Execute(w, results)
	if err != nil {
		panic(err.Error())
	}
}

// LoginT renders the login.html page
func (rr *Routes) LoginT(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("static/login.html")
	var dummy string
	err := tmpl.Execute(w, dummy)
	if err != nil {
		panic(err.Error())
	}
}

// Login verifies the given account information, compares the profile
// and the respective hash of the password with the stored accounts data in database
// and redirects to index.html if true
func (rr *Routes) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err.Error()) // Handle error here via logging and then return
	}
	var acc entities.Account
	acc.Profile = r.PostFormValue("user")
	password := r.PostFormValue("password")
	profile, hash, _ := rr.Db.GetAccountByProfile(acc.Profile)
	hashStr := string(hash)
	bl := src.ComparePasswords(hashStr, password)
	if (bl == true) && (acc.Profile == profile) {
		defer rr.GetIndexT(w, r)
	} else {
		fmt.Println("password invalid!")
		defer rr.LoginT(w, r)
	}
}

// AddPageT renders the addPage.html with the form for insert invoice to database
func (rr *Routes) AddPageT(w http.ResponseWriter, r *http.Request) {
	_, _ = ioutil.ReadAll(r.Body)
	Title := "New Invoice"
	tmpl, _ := template.ParseFiles("static/addPage.html")
	err := tmpl.Execute(w, Title)
	if err != nil {
		panic(err.Error())
	}
}

// AddInvoice collect data from the form in addPage.html and calls
// database.InsertInvoice to insert the invoice to the database
func (rr *Routes) AddInvoice(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		panic(err.Error()) // Handle error here via logging and then return
	}
	var inv entities.Invoice
	inv.Document = r.PostFormValue("document")
	inv.Description = r.PostFormValue("description")
	inv.Amount, _ = strconv.ParseFloat(r.PostFormValue("amount"), 64)
	inv.CreatedAt = r.PostFormValue("createdAt")
	inv.DeactivatedAt = r.PostFormValue("deactivatedAt")
	inv.ReferenceMonth, _ = strconv.Atoi(r.PostFormValue("referenceMonth"))
	inv.ReferenceYear, _ = strconv.Atoi(r.PostFormValue("referenceYear"))
	_ = rr.Db.InsertInvoice(inv)
	rr.AddPageT(w, r)
}

// GetAll gets all invoices and returns in json format
func (rr *Routes) GetAll(w http.ResponseWriter, r *http.Request) { // get invoices and returns in json format
	var results []entities.Invoice
	results, err := rr.Db.GetAll()
	if err != nil {
		panic(err.Error())
	} else {
	}
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
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(invoice)
}

// DeleteInvoice makes the logic deletion of the invoice by the given ID in the rr.Db
func (rr *Routes) DeleteInvoice(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	var invoice entities.Invoice // Unmarshal
	err = json.Unmarshal(b, &invoice)
	idNotInList, err := rr.Db.GetInvoiceByID(invoice.ID)
	if err != nil {
		panic(err.Error())
	}
	if (idNotInList == entities.Invoice{}) {
		println("id not found!")
		w.WriteHeader(http.StatusNotFound)
	} else {
		err := rr.Db.DeleteInvoice(invoice.ID)
		if err != nil {
			panic(err.Error())
		}
		w.WriteHeader(http.StatusNoContent)
	}
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(invoice.ID)
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
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(editedInvoice)
}

//
func (rr *Routes) Pagination(w http.ResponseWriter, r *http.Request) { // get invoices and returns in json format
	var results []entities.Invoice
	results, err := rr.Db.Pagination()
	if err != nil {
		panic(err.Error())
	} else {
	}
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(results)
}
