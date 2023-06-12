package models

import (
	"github.com/jinzhu/gorm"
	"github.com/sarraj/go-bookstore/pkg/config"
)

var Ibook IBook

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

func Init() {
	//config.InitApp()
	//db = config.GetDB()
	config.Db.AutoMigrate(&Book{})

}

func (b *Book) CreateBook() *Book {
	config.Db.NewRecord(b)
	config.Db.Create(&b)
	return b
}

func (b *Book) GetAllBooks() []Book {
	var Books []Book
	config.Db.Find(&Books)
	return Books
}
func (b *Book) GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := config.Db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}
func (b *Book) DeleteBook(ID int64) Book {
	var book Book
	config.Db.Where("ID=?", ID).Delete(&book)
	return book
}
