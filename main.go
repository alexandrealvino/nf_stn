package main

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"net/http"
	"nf_stn/adapter"
	"nf_stn/authentication"
	"nf_stn/config"
	"nf_stn/database"
	"nf_stn/middleware"
)

func main() {
	router := mux.NewRouter() // init router
	log.Println("Server started on: http://localhost:8000")

	// initializing app
	rd := config.RedisCfg{}
	cfg :=config.Config{}
	db := database.MySQL{
		Config: &cfg,
	}
	tk := authentication.Auth{
		RedisCg: &rd,
	}
	tk.Init()
	db.Init()
	routes := adapter.Routes{
		Db: &db,
		Au: &tk,
	}
	// router handlers

	router.Handle("/api",middleware.JoinMiddleWares(routes.GetAll)).Methods("GET")                                                             // get invoices list and returns data in json format
	router.Handle("/api/", middleware.JoinMiddleWares(routes.InsertInvoice)).Methods("POST")                                       // insert invoice
	router.Handle("/api/", middleware.JoinMiddleWares(routes.UpdateInvoice)).Methods("PUT")                                                  // update invoice
	router.Handle("/api/", middleware.JoinMiddleWares(routes.PatchInvoice)).Methods("PATCH")                                              // patch invoice
	router.Handle("/api/{document}", middleware.JoinMiddleWares(routes.GetInvoiceByDocument)).Methods("GET")                         // get invoice by document and returns data in json format
	router.Handle("/api/{ID}", middleware.JoinMiddleWares(routes.DeleteInvoice)).Methods("DELETE")                                               // set isActive = 0 for logic deletion
	router.Handle("/api/pagination", middleware.JoinMiddleWares(routes.Pagination)).Methods("GET")                                              // paginates by id, 10 invoices per page
	router.Handle("/api/pagination/{offset}/month/{referenceMonth}", middleware.JoinMiddleWares(routes.PaginationByMonth)).Methods("GET")       // paginates by month, 10 invoices per page
	router.Handle("/api/pagination/{offset}/year/{referenceYear}", middleware.JoinMiddleWares(routes.PaginationByYear)).Methods("GET")          // paginates by year, 10 invoices per page
	router.Handle("/api/pagination/{offset}/document/{document}", middleware.JoinMiddleWares(routes.PaginationByDocument)).Methods("GET")       // paginates by document, 10 invoices per page
	router.Handle("/api/pagination/{offset}/month", middleware.JoinMiddleWares(routes.PaginationOrderByMonth)).Methods("GET")                   // paginates by month, 10 invoices per page
	router.Handle("/api/pagination/{offset}/year", middleware.JoinMiddleWares(routes.PaginationOrderByYear)).Methods("GET")                     // paginates by year, 10 invoices per page
	router.Handle("/api/pagination/{offset}/document", middleware.JoinMiddleWares(routes.PaginationOrderByDocument)).Methods("GET")             // paginates by document, 10 invoices per page
	router.Handle("/api/pagination/{offset}/month/year/", middleware.JoinMiddleWares(routes.PaginationOrderByMonthYear)).Methods("GET")         // paginates by month and year, 10 invoices per page
	router.Handle("/api/pagination/{offset}/month/document/", middleware.JoinMiddleWares(routes.PaginationOrderByMonthDocument)).Methods("GET") // paginates by month and document, 10 invoices per page
	router.Handle("/api/pagination/{offset}/year/document/", middleware.JoinMiddleWares(routes.PaginationOrderByYearDocument)).Methods("GET")   // paginates by year and document, 10 invoices per page


	router.HandleFunc("/api/login", routes.GenerateToken).Methods("POST")   // generates token for authenticated user
	router.HandleFunc("/api/logout", tk.Logout).Methods("POST") // logout user

	router.Handle("/metrics", promhttp.Handler())				// get metrics for future metrics handler

	//arrumar rotas, erros, log erros
	log.Fatal(http.ListenAndServe(":8000", router)) // if error return fatal log
}

//
