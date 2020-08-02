package adapter

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math"
	"net/http"
	"nf_stn/authentication"
	"nf_stn/database"
	"nf_stn/entities"
	"nf_stn/lib"
	"strconv"
	//"github.com/golang/mock"
)

// Routes struct
type Routes struct {
	Db database.DataBase
	Au authentication.Token
}

// GetInvoices gets page list of invoices by the select parameters conditions, 10 invoices per page
func (rr *Routes) GetInvoices(w http.ResponseWriter, r *http.Request) { // get invoices and returns in json format
	var err error
	var params entities.SearchParameters
	var results []entities.Invoice
	var info entities.Info
	var totalPages int

	// getting request parameters
	params.Page, _ = strconv.Atoi(r.FormValue("page"))
	params.OrderBy = r.FormValue("orderBy")
	params.Month, _ = strconv.Atoi(r.FormValue("month"))
	params.Year, _ = strconv.Atoi(r.FormValue("year"))
	params.Document = r.FormValue("document")
	params.Deletes, _ = strconv.Atoi(r.FormValue("deletes"))

	// checking page parameter
	if params.Page <= 0 {
		info.Page = "1"
	} else {
		info.Page = strconv.Itoa(params.Page)
	}

	// fetching db data
	results, info.InvoicesFound, err = rr.Db.GetInvoices(params)
	if err != nil {
		log.Error(err.Error())
	}

	// calculating number of pages
	if info.InvoicesFound%10 != 0 {
		totalPages = int(math.Round(float64(info.InvoicesFound/10))) + 1
	} else {
		totalPages = info.InvoicesFound / 10
	}

	// setting the info struct
	info.AuthenticationStatus = "authorized"
	info.RequestMethod = r.Method
	info.ContentType = "application/json"
	info.TotalPages = strconv.Itoa(totalPages)
	info.Invoices = results
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	if params.Page*10 <= info.InvoicesFound+9 {
		_ = encoder.Encode(info)
	} else {
		info.Page = "not found"
		_ = encoder.Encode(info)
	}
}

// GetInvoiceByDocument gets invoice by document value and returns the invoice in json format
func (rr *Routes) GetInvoiceByDocument(w http.ResponseWriter, r *http.Request) {
	var info entities.Info
	var invoice, result entities.Invoice
	params := mux.Vars(r)
	invoice.Document = params["document"]
	result, err := rr.Db.GetInvoiceByDocument(invoice.Document)
	if err != nil {
		log.Error("invoice not found")
	} else {
	}

	page, _ := strconv.Atoi(r.FormValue("page"))
	if page <= 0 {
		info.Page = "1"
	}

	// setting the info struct
	info.Page = strconv.Itoa(page)
	info.AuthenticationStatus = "authorized"
	info.RequestMethod = r.Method
	info.ContentType = "application/json"
	info.TotalPages = "1"
	info.InvoicesFound = 1
	info.Invoices = append(info.Invoices, result)
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(info)
}

// InsertInvoice inserts the given invoice data in the rr.Db
func (rr *Routes) InsertInvoice(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("request body incorrect format")
	}
	var invoice entities.Invoice // Unmarshal
	err = json.Unmarshal(b, &invoice)
	if err != nil {
		log.Error("request body incorrect format")
	}
	if invoice.ReferenceMonth > 12 || invoice.ReferenceMonth < 0 {
		log.Error("invalid month!")
		return
	}
	err = rr.Db.InsertInvoice(invoice)
	if err != nil {
		log.Error(err)
		w.WriteHeader(http.StatusConflict)
		return
	}

	// setting the info struct
	var info entities.Info
	info.Page = strconv.Itoa(1)
	info.AuthenticationStatus = "authorized"
	info.RequestMethod = r.Method
	info.ContentType = "application/json"
	info.TotalPages = "1"
	info.InvoicesFound = 1
	info.Invoices = append(info.Invoices, invoice)
	w.WriteHeader(http.StatusCreated)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(info)
}

// DeleteInvoice makes the logic deletion of the invoice by the given ID in the rr.Db
func (rr *Routes) DeleteInvoice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	ID, _ := strconv.Atoi(params["ID"])

	inv, err := rr.Db.GetInvoiceByID(ID)
	if err != nil {
		log.Error("id not found!")
		return
	}
	if (inv == entities.Invoice{}) {
		log.Error("id not found!")
		w.WriteHeader(http.StatusNotFound)
	} else {
		err := rr.Db.DeleteInvoice(ID)
		if err != nil {
			log.Error("delete unsuccessful")
		}
		inv.IsActive = 0
		w.WriteHeader(http.StatusOK)
	}

	// setting the info struct
	var info entities.Info
	info.Page = strconv.Itoa(1)
	info.AuthenticationStatus = "authorized"
	info.RequestMethod = r.Method
	info.ContentType = "application/json"
	info.TotalPages = "1"
	info.InvoicesFound = 1
	info.Invoices = append(info.Invoices, inv)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(info)
}

// UpdateInvoice updates database values from the row of the given invoice
func (rr *Routes) UpdateInvoice(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("request body incorrect format")
	}
	var invoice entities.Invoice
	err = json.Unmarshal(b, &invoice) // Unmarshal
	if err != nil {
		log.Error("request body incorrect format")
	}
	invoiceExists, _ := rr.Db.InvoiceExists(invoice.Document)
	if (invoiceExists == entities.Invoice{}) {
		log.Info("Invoice not found!")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	invoice.ID = invoiceExists.ID
	err = rr.Db.UpdateInvoice(invoice)
	if err != nil {
			log.Error("update unsuccessful")
		}

	// setting the info struct
	var info entities.Info
	info.Page = strconv.Itoa(1)
	info.AuthenticationStatus = "authorized"
	info.RequestMethod = r.Method
	info.ContentType = "application/json"
	info.TotalPages = "1"
	info.InvoicesFound = 1
	invoice.Amount /= 100 // handling bigint
	info.Invoices = append(info.Invoices, invoice)
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(info)
}

// PatchInvoice partially updates rr.Db values from the row of the given invoice
func (rr *Routes) PatchInvoice(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Content-Type", "application/json")
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Error("request body incorrect format")
	}
	var invoice, editedInvoice entities.Invoice
	_ = json.Unmarshal(b, &invoice) // Unmarshal
	invoiceExists, _ := rr.Db.InvoiceExists(invoice.Document)
	if (invoiceExists == entities.Invoice{}) {
		log.Info("Invoice not found!")
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err = rr.Db.PatchInvoice(invoice)
	if err != nil {
		log.Error("patch unsuccessful")
	}
	editedInvoice, err = rr.Db.GetInvoiceByID(invoice.ID)
	if err != nil {
		log.Error("patch unsuccessful")
	}

	// setting the info struct
	var info entities.Info
	info.Page = strconv.Itoa(1)
	info.AuthenticationStatus = "authorized"
	info.RequestMethod = r.Method
	info.ContentType = "application/json"
	info.TotalPages = "1"
	info.InvoicesFound = 1
	info.Invoices = append(info.Invoices, editedInvoice)
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(info)
}

// GenerateToken generates token for authenticated user
func (rr *Routes) GenerateToken(w http.ResponseWriter, r *http.Request) {
	var user, u entities.User
	var err error
	u.Username = r.Header.Get("username")
	pwd := r.Header.Get("password")
	//u.Hash = r.Header.Get("password")
	user.ID, user.Username, user.Hash, err = rr.Db.GetUser(u.Username)
	if err != nil {
		log.Error("user status: not found!")
		return
	}
	//compare the user from the request, with the one defined in database:
	isOk := lib.ComparePasswords(user.Hash, pwd)
	if isOk != true {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	token, err := rr.Au.CreateToken(uint64(user.ID), user.Username)
	if err != nil {
		panic(err)
		return
	}
	saveErr := rr.Au.CreateAuth(uint64(user.ID), token)
	if saveErr != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnprocessableEntity)
		log.Println(saveErr)
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
