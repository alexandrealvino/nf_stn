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
	log.Println("Server starting on: http://localhost:8000")

	// initializing app
	rd := config.RedisCfg{}
	cfg := config.Config{}
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
	router.Handle("/api", middleware.JoinMiddleWares(routes.GetInvoices)).Methods("GET")                     // get invoices list and returns data in json format
	router.Handle("/api", middleware.JoinMiddleWares(routes.InsertInvoice)).Methods("POST")                  // insert invoice
	router.Handle("/api", middleware.JoinMiddleWares(routes.UpdateInvoice)).Methods("PUT")                   // update invoice
	router.Handle("/api", middleware.JoinMiddleWares(routes.PatchInvoice)).Methods("PATCH")                  // patch invoice
	router.Handle("/api/{document}", middleware.JoinMiddleWares(routes.GetInvoiceByDocument)).Methods("GET") // get invoice by document and returns data in json format
	router.Handle("/api/{ID}", middleware.JoinMiddleWares(routes.DeleteInvoice)).Methods("DELETE")           // set isActive = 0 for logic deletion

	router.HandleFunc("/api/login", routes.GenerateToken).Methods("POST") // generates token for authenticated user
	router.HandleFunc("/api/logout", tk.Logout).Methods("POST")           // logout user
	router.HandleFunc("/api/refresh", tk.RefreshToken).Methods("POST")	// refreshes token

	router.Handle("/metrics", promhttp.Handler()) // get metrics for future metrics handler

	log.Println("Server started on: http://localhost:8000")

	log.Fatal(http.ListenAndServe(":8000", router)) // if error return fatal log
}

//
