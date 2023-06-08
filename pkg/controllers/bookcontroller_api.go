package controllers

import "net/http"

type Ibc interface {
	GetBook(w http.ResponseWriter, r *http.Request)
	GetBookById(w http.ResponseWriter, r *http.Request)
	CreateBook(w http.ResponseWriter, r *http.Request)
	DeleteBook(w http.ResponseWriter, r *http.Request)
	UpdateBook(w http.ResponseWriter, r *http.Request)
}
