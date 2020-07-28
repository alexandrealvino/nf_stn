package main

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"nf_stn/adapter"
	"nf_stn/config"
	"nf_stn/database"
	"nf_stn/entities"
	"nf_stn/src/middleware"

	//"github.com/sirupsen/logrus"
	//"github.com/spf13/viper"
)

type handler struct {}

//A sample use
var user = entities.User{
	ID:             1,
	Username: "username",
	Password: "password",
}

func main() {


	router := mux.NewRouter() // init router
	log.Println("Server started on: http://localhost:8000")


	cfg :=config.Config{}
	db := database.MySql{
		Config: &cfg,
	}
	db.Init()
	routes := adapter.Routes{
		Db: &db,
	}
	// router handlers

	router.Handle("/api", middleware.Logger(middleware.Authentication(routes.GetAll))).Methods("GET")                              // get invoices list and returns data in json format
	router.Handle("/api/getInvoiceByDocument/", middleware.Logger(middleware.Authentication(routes.GetInvoiceByDocument))).Methods("GET") // get invoice by document and returns data in json format
	router.Handle("/api/insertInvoice", middleware.Logger(middleware.Authentication(routes.InsertInvoice))).Methods("POST")               // insert invoice
	router.Handle("/api/del", middleware.Logger(middleware.Authentication(routes.DeleteInvoice))).Methods("DELETE")                     // set isActive = 0 for logic deletion
	router.Handle("/api/up/", middleware.Logger(middleware.Authentication(routes.UpdateInvoice))).Methods("PUT")                          // update invoice
	router.Handle("/api/patch/", middleware.Logger(middleware.Authentication(routes.PatchInvoice))).Methods("PATCH")                      // patch invoice
	router.Handle("/api/pagination", middleware.Logger(middleware.Authentication(routes.Pagination))).Methods("GET")
	router.Handle("/api/pagination/{offset}/month", middleware.Logger(middleware.Authentication(routes.PaginationOrderByMonth))).Methods("GET")
	router.Handle("/api/pagination/{offset}/year", middleware.Logger(middleware.Authentication(routes.PaginationOrderByYear))).Methods("GET")
	router.Handle("/api/pagination/{offset}/document", middleware.Logger(middleware.Authentication(routes.PaginationOrderByDocument))).Methods("GET")
	router.Handle("/api/pagination/{offset}/month/year/", middleware.Logger(middleware.Authentication(routes.PaginationOrderByMonthYear))).Methods("GET")
	router.Handle("/api/pagination/{offset}/month/document/", middleware.Logger(middleware.Authentication(routes.PaginationOrderByMonthDocument))).Methods("GET")
	router.Handle("/api/pagination/{offset}/year/document/", middleware.Logger(middleware.Authentication(routes.PaginationOrderByYearDocument))).Methods("GET")
	router.Handle("/api/pagination/{offset}/month/{referenceMonth}", middleware.Logger(middleware.Authentication(routes.PaginationByMonth))).Methods("GET")
	router.Handle("/api/pagination/{offset}/year/{referenceYear}", middleware.Logger(middleware.Authentication(routes.PaginationByYear))).Methods("GET")
	router.Handle("/api/pagination/{offset}/document/{document}", middleware.Logger(middleware.Authentication(routes.PaginationByDocument))).Methods("GET")

	//router.HandleFunc("/api/login", routes.Login).Methods("GET")

	router.Handle("/metrics", promhttp.Handler())				// get metrics


	log.Fatal(http.ListenAndServe(":8000", router)) // if error return fatal log
}

//
