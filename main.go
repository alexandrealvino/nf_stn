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
	cfg :=config.Config{}
	db := database.MySql{
		Config: &cfg,
	}
	routes := adapter.Routes{
		Db: &db,
	}
	// router handlers

	router.HandleFunc("/login", routes.LoginT).Methods("GET")            // render login.html
	router.HandleFunc("/login", routes.Login).Methods("POST")            // try to login and redirects to index.html if true
	router.HandleFunc("/api", routes.GetIndexT).Methods("GET")           // render index.html
	router.HandleFunc("/api/addPage", routes.AddPageT).Methods("GET")    // render addPage.html
	router.HandleFunc("/api/addPage", routes.AddInvoice).Methods("POST") // insert invoice from addPage.html

	router.HandleFunc("/api/getAll", routes.GetAll).Methods("GET")                              // get invoices list and returns data in json format
	router.HandleFunc("/api/getInvoiceByDocument/", routes.GetInvoiceByDocument).Methods("GET") // get invoice by document and returns data in json format
	router.HandleFunc("/api/insertInvoice", routes.InsertInvoice).Methods("POST")               // insert invoice
	router.HandleFunc("/api/del/", routes.DeleteInvoice).Methods("DELETE")                      // set isActive = 0 for logic deletion
	router.HandleFunc("/api/up/", routes.UpdateInvoice).Methods("PUT")                          // update invoice
	router.HandleFunc("/api/patch/", routes.PatchInvoice).Methods("PATCH")                      // patch invoice

	router.HandleFunc("/api/pagination", routes.Pagination).Methods("GET")

	router.Handle("/metrics", promhttp.Handler())				// get metrics

	log.Fatal(http.ListenAndServe(":8000", router)) // if error return fatal log
}

//
