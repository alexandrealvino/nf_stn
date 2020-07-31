package main

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"nf_stn/adapter"
	"nf_stn/config"
	"nf_stn/database"
	"nf_stn/src/middleware"
	//"github.com/sirupsen/logrus"
	//"github.com/spf13/viper"
)

func main() {
	router := mux.NewRouter() // init router
	log.Println("Server started on: http://localhost:8000")

	// initializing app

	cfg :=config.Config{}
	db := database.MySQL{
		Config: &cfg,
	}
	db.Init()
	routes := adapter.Routes{
		Db: &db,
	}

	// router handlers

	router.Handle("/api", middleware.Logger(middleware.TokenAuthMiddleware(routes.GetAll))).Methods("GET")                              // get invoices list and returns data in json format
	router.Handle("/api/getInvoiceByDocument/", middleware.Logger(middleware.TokenAuthMiddleware(routes.GetInvoiceByDocument))).Methods("GET") // get invoice by document and returns data in json format
	router.Handle("/api/insertInvoice", middleware.Logger(middleware.TokenAuthMiddleware(routes.InsertInvoice))).Methods("POST")               // insert invoice
	router.Handle("/api/del", middleware.Logger(middleware.TokenAuthMiddleware(routes.DeleteInvoice))).Methods("DELETE")                     // set isActive = 0 for logic deletion
	router.Handle("/api/up/", middleware.Logger(middleware.TokenAuthMiddleware(routes.UpdateInvoice))).Methods("PUT")                          // update invoice
	router.Handle("/api/patch/", middleware.Logger(middleware.TokenAuthMiddleware(routes.PatchInvoice))).Methods("PATCH")                      // patch invoice
	router.Handle("/api/pagination", middleware.Logger(middleware.TokenAuthMiddleware(routes.Pagination))).Methods("GET") // paginates by id, 10 invoices per page
	router.Handle("/api/pagination/{offset}/month/{referenceMonth}", middleware.Logger(middleware.TokenAuthMiddleware(routes.PaginationByMonth))).Methods("GET") // paginates by month, 10 invoices per page
	router.Handle("/api/pagination/{offset}/year/{referenceYear}", middleware.Logger(middleware.TokenAuthMiddleware(routes.PaginationByYear))).Methods("GET") // paginates by year, 10 invoices per page
	router.Handle("/api/pagination/{offset}/document/{document}", middleware.Logger(middleware.TokenAuthMiddleware(routes.PaginationByDocument))).Methods("GET") // paginates by document, 10 invoices per page
	router.Handle("/api/pagination/{offset}/month", middleware.Logger(middleware.TokenAuthMiddleware(routes.PaginationOrderByMonth))).Methods("GET")  // paginates by month, 10 invoices per page
	router.Handle("/api/pagination/{offset}/year", middleware.Logger(middleware.TokenAuthMiddleware(routes.PaginationOrderByYear))).Methods("GET") // paginates by year, 10 invoices per page
	router.Handle("/api/pagination/{offset}/document", middleware.Logger(middleware.TokenAuthMiddleware(routes.PaginationOrderByDocument))).Methods("GET") // paginates by document, 10 invoices per page
	router.Handle("/api/pagination/{offset}/month/year/", middleware.Logger(middleware.TokenAuthMiddleware(routes.PaginationOrderByMonthYear))).Methods("GET") // paginates by month and year, 10 invoices per page
	router.Handle("/api/pagination/{offset}/month/document/", middleware.Logger(middleware.TokenAuthMiddleware(routes.PaginationOrderByMonthDocument))).Methods("GET") // paginates by month and document, 10 invoices per page
	router.Handle("/api/pagination/{offset}/year/document/", middleware.Logger(middleware.TokenAuthMiddleware(routes.PaginationOrderByYearDocument))).Methods("GET") // paginates by year and document, 10 invoices per page



	router.HandleFunc("/api/login", routes.GenerateToken).Methods("POST") // generates token for authenticated user
	router.HandleFunc("/api/logout", middleware.Logout).Methods("POST") // logout user

	router.Handle("/metrics", promhttp.Handler())				// get metrics

	//router.HandleFunc("/api/createTodo", middleware.CreateTodo).Methods("POST") // creates todo

	//arrumar rotas, erros, log erros, redis fetch, struct
	log.Fatal(http.ListenAndServe(":8000", router)) // if error return fatal log
}

//
