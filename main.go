package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/star-wars-go/config"
	. "github.com/star-wars-go/dao"
	planetRouter "github.com/star-wars-go/router"
)

var dao = PlanetsDAO{}
var config = Config{}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/planets", planetRouter.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/planets/{id}", planetRouter.GetByID).Methods("GET")
	r.HandleFunc("/api/v1/planets", planetRouter.Create).Methods("POST")
	r.HandleFunc("/api/v1/planets/{id}", planetRouter.Update).Methods("PUT")
	r.HandleFunc("/api/v1/planets/{id}", planetRouter.Delete).Methods("DELETE")

	var port = ":3000"
	fmt.Println("Server running in port: ", port)
	log.Fatal(http.ListenAndServe(port, r))
}
