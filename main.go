package main

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"nf_stn/adapter"
	"nf_stn/config"
	"nf_stn/database"
	//"github.com/sirupsen/logrus"
	//"github.com/spf13/viper"
)

func main() {


	router := mux.NewRouter() // init router
	log.Println("Server started on: http://localhost:8000")

	//monthDay, month, hour, min, sec, year := time.Now().Day(), time.Now().Month(), time.Now().Hour(), time.Now().Minute(), time.Now().Second(), time.Now().Year()
	//date := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(monthDay)
	//clock := strconv.Itoa(hour) + ":" + strconv.Itoa(min) + ":" + strconv.Itoa(sec)
	//now := date + " " + clock
	//fmt.Println(now)
	//var ms = database.MySql{}
	//ms.Init()

	cfg :=config.Config{}
	db := database.MySql{
		Config: &cfg,
	}
	db.Init()
	routes := adapter.Routes{
		Db: &db,
	}
	// router handlers

	router.HandleFunc("/api", routes.GetAll).Methods("GET")                              // get invoices list and returns data in json format
	router.HandleFunc("/api/getInvoiceByDocument/", routes.GetInvoiceByDocument).Methods("GET") // get invoice by document and returns data in json format
	router.HandleFunc("/api/insertInvoice", routes.InsertInvoice).Methods("POST")               // insert invoice
	router.HandleFunc("/api/del/", routes.DeleteInvoice).Methods("DELETE")                      // set isActive = 0 for logic deletion
	router.HandleFunc("/api/up/", routes.UpdateInvoice).Methods("PUT")                          // update invoice
	router.HandleFunc("/api/patch/", routes.PatchInvoice).Methods("PATCH")                      // patch invoice
	router.HandleFunc("/api/pagination/{offset}", routes.Pagination).Methods("GET")
	router.HandleFunc("/api/pagination/{offset}/month", routes.PaginationOrderByMonth).Methods("GET")
	router.HandleFunc("/api/pagination/{offset}/year", routes.PaginationOrderByYear).Methods("GET")
	router.HandleFunc("/api/pagination/{offset}/document", routes.PaginationOrderByDocument).Methods("GET")
	router.HandleFunc("/api/pagination/{offset}/month/year/", routes.PaginationOrderByMonthYear).Methods("GET")
	router.HandleFunc("/api/pagination/{offset}/month/document/", routes.PaginationOrderByMonthDocument).Methods("GET")
	router.HandleFunc("/api/pagination/{offset}/year/document/", routes.PaginationOrderByYearDocument).Methods("GET")
	router.HandleFunc("/api/pagination/{offset}/month/{referenceMonth}", routes.PaginationByMonth).Methods("GET")
	router.HandleFunc("/api/pagination/{offset}/year/{referenceYear}", routes.PaginationByYear).Methods("GET")
	router.HandleFunc("/api/pagination/{offset}/document/{document}", routes.PaginationByDocument).Methods("GET")

	router.Handle("/metrics", promhttp.Handler())				// get metrics

	log.Fatal(http.ListenAndServe(":8000", router)) // if error return fatal log
}

//
