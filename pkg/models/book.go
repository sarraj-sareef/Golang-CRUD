package models

import (
	"github.com/jinzhu/gorm"
	"github.com/sarraj/go-bookstore/pkg/config"
)

var Ibook IBook
var db *gorm.DB

type IBook interface {
	CreateBook() *Book
	GetAllBooks() []Book
	GetBookById(Id int64) (*Book, *gorm.DB)
	DeleteBook(ID int64) Book
}

type Book struct {
	//gorm.Model
	Name   string `gorm:""json:"name"`
	Author string `json:"author"`
	ID     int    `json:"id"`
}

func Newbook(name string, author string, id int) IBook {

	Ibook = &Book{
		Name:   name,
		Author: author,
		ID:     id,
	}
	return Ibook
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})

}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func (b *Book) GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}
func (b *Book) GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}
func (b *Book) DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(&book)
	return book
}
