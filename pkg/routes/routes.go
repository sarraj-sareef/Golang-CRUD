package routes

import (
	"github.com/gorilla/mux"
	"github.com/sarraj/go-bookstore/pkg/controllers"
	"github.com/sarraj/go-bookstore/pkg/models"
)

var RegisterBookstoreRoutes = func(router *mux.Router) {

	bkInstance := controllers.NewbookController(&models.Book{}) // Create an instance of the bk struct
	router.HandleFunc("/book/", bkInstance.CreateBook).Methods("POST")
	router.HandleFunc("/book/", bkInstance.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", bkInstance.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", bkInstance.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}", bkInstance.DeleteBook).Methods("DELETE")
}
