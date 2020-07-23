package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"nf_stn/adapter"
	"strconv"
	"time"
)

func main() {

	router := mux.NewRouter()  // init router
	log.Println("Server started on: http://localhost:8000")
	//year, month, day := time.Now().Date()
	//hour, min, sec :=  time.Now().Clock()
	monthDay := time.Now().Day()
	month := time.Now().Month()
	hour := time.Now().Hour()
	min := time.Now().Minute()
	sec := time.Now().Second()
	year := time.Now().Year()
	date := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(monthDay)
	clock := strconv.Itoa(hour) + ":" + strconv.Itoa(min) + ":" + strconv.Itoa(sec)
	now := date + " " + clock
	fmt.Println(now)

	// router handlers

	router.HandleFunc("/api", adapter.GetIndexT).Methods("GET") // render index.html
	router.HandleFunc("/api/addPage", adapter.AddPageT).Methods("GET") // render addPage.html
	router.HandleFunc("/api/addPage", adapter.AddInvoice).Methods("POST") // insert invoice from addPage.html

	router.HandleFunc("/api/getAll", adapter.GetAll).Methods("GET") // get invoices list and returns data in json format
	router.HandleFunc("/api/getInvoiceByDocument/", adapter.GetInvoiceByDocument).Methods("GET") // get invoice by document and returns data in json format
	router.HandleFunc("/api/insertInvoice", adapter.InsertInvoice).Methods("POST") // insert invoice
	router.HandleFunc("/api/del/", adapter.DeleteInvoice).Methods("DELETE") // set isActive = 0 for logic deletion
	router.HandleFunc("/api/up/", adapter.UpdateInvoice).Methods("PUT") // update invoice
	router.HandleFunc("/api/patch/", adapter.PatchInvoice).Methods("PATCH") // patch invoice

	log.Fatal(http.ListenAndServe(":8000", router)) // if error return fatal log
}
//