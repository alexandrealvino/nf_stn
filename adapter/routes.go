package adapter

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"nf_stn/database"
	"nf_stn/entities"
	//"github.com/golang/mock"
)

type Routes struct {
	Db database.DataBase
}


// GetAll gets all invoices and returns in json format
func (rr *Routes) GetAll(w http.ResponseWriter, r *http.Request) { // get invoices and returns in json format
	var results []entities.Invoice
	results, err := rr.Db.GetAll()
	if err != nil {
		panic(err.Error())
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
	r.Header.Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	var interval []int
	err = json.Unmarshal(b, &interval) // Unmarshal
	if err != nil {
		panic(err.Error())
	}
	results, err := rr.Db.Pagination(interval[0],interval[1])
	if err != nil {
		panic(err.Error())
	} else {
	}
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(results)
}
