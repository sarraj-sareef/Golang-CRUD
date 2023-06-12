package main

import (
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sarraj/go-bookstore/pkg/config"
	"github.com/sarraj/go-bookstore/pkg/routes"
	"log"
	"net/http"
)

func main() {
	config.InitApp()
	r := mux.NewRouter()
	routes.RegisterBookstoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
