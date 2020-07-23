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
)


func GetIndexT(w http.ResponseWriter, r *http.Request) { // render "wallet" template
	var results []entities.Invoice
	results, err := database.GetAll()
	if err != nil {	panic(err.Error())}
	tmpl, _ := template.ParseFiles("static/index.html")
	err = tmpl.Execute(w, results)
	if err != nil {
		panic(err.Error())
	}
} // render "wallet" template
//
func AddPageT(w http.ResponseWriter, r *http.Request) { // execute "addticker" template
	_, _ = ioutil.ReadAll(r.Body)
	Title := "New Invoice"
	tmpl, _ := template.ParseFiles("static/addPage.html")
	err := tmpl.Execute(w, Title)
	if err != nil {
		panic(err.Error())
	}
} // render addPage.html
//
func AddInvoice(w http.ResponseWriter, r *http.Request) { // add stock buy to the "buys" table in database
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
	password := r.PostFormValue("password")
	hash, _ := database.GetHashByProfile("admin")
	hashStr := string(hash)
	bl := src.ComparePasswords(hashStr,password)
	fmt.Println(password,hashStr)
	if bl == true {
		_ = database.InsertInvoice(inv)
		defer AddPageT(w,r)
	} else {
		fmt.Println("password invalid!")
		defer AddPageT(w,r)
	}
	//AddPageT(w,r)
} // add stock buy to the "buys" table in database
//
//
func GetAll(w http.ResponseWriter, r *http.Request) { // get invoices and returns in json format
	var results []entities.Invoice
	results, err := database.GetAll()
	if err != nil {	panic(err.Error())} else {}
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(results)
	fmt.Println(results)
} // get wallet in json format
//
func GetInvoiceByDocument(w http.ResponseWriter, r *http.Request) { // get invoices and returns in json format
	r.Header.Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	var invoice entities.Invoice  // Unmarshal
	err = json.Unmarshal(b, &invoice)
	var result entities.Invoice
	result, err = database.GetInvoiceByDocument(invoice.Document)
	if err != nil {	panic(err.Error())} else {}
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(result)
	//json.NewEncoder(w).Encode(results)
} // get wallet in json format
//
func InsertInvoice(w http.ResponseWriter, r *http.Request) { // insert invoice
	r.Header.Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	var invoice entities.Invoice  // Unmarshal
	err = json.Unmarshal(b, &invoice)
	if err != nil {
		panic(err.Error())
	}
	err = database.InsertInvoice(invoice)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(invoice)
} // insert invoice
//
func DeleteInvoice(w http.ResponseWriter, r *http.Request) {  // delete invoice
	r.Header.Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	var invoice entities.Invoice  // Unmarshal
	err = json.Unmarshal(b, &invoice)
	idNotInList, err := database.GetInvoiceById(invoice.ID)
	if err != nil {
		panic(err.Error())
	}
	if (idNotInList == entities.Invoice{}) {
		println("id not found!")
		w.WriteHeader(http.StatusNotFound)
	} else {
		err := database.DeleteInvoice(invoice.ID)
		if err != nil {
			panic(err.Error())
		}
		w.WriteHeader(http.StatusNoContent)
	}
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(invoice.ID)
} // delete invoice
//
func UpdateInvoice(w http.ResponseWriter, r *http.Request) {  // update invoice
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
	invoiceExists, _ := database.InvoiceExists(invoice.Document)
	if (invoiceExists == entities.Invoice{})  {
		println("Invoice not found!")
		w.WriteHeader(http.StatusNotFound)
	}	else {
		invoice.ID = invoiceExists.ID
		err := database.UpdateInvoice(invoice)
		if err != nil {
			panic(err.Error())
		}
	}
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(invoice)
} // update invoice
//
func PatchInvoice(w http.ResponseWriter, r *http.Request) {  // partial updates invoice
	r.Header.Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	var invoice, editedInvoice entities.Invoice
	_ = json.Unmarshal(b, &invoice) // Unmarshal
	err = database.PatchInvoice(invoice)
	if err != nil {	panic(err.Error())}
	editedInvoice, err = database.GetInvoiceById(invoice.ID)
	if err != nil {	panic(err.Error())}
	w.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(editedInvoice)
} // partial updates invoice
//
