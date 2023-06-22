package controllers

import (
	"encoding/json"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/suite"

	"github.com/sarraj/go-bookstore/pkg/models"
)

type BookControllerTestSuite struct {
	suite.Suite
	controller *Bk
	bookMock   *models.Mockibook
}

func TestBookController(t *testing.T) {
	suite.Run(t, new(BookControllerTestSuite))
}

func (suite *BookControllerTestSuite) SetupTest() {
	ctrl := gomock.NewController(suite.T())
	suite.bookMock = models.NewMockibook(ctrl)
	suite.controller = &Bk{
		ebook: suite.bookMock,
	}

}

func (suite *BookControllerTestSuite) TestGetBook() {
	// Prepare mock data and expectations
	books := []models.Book{
		{ID: 1, Name: "Book 1", Author: "Author 1"},
		{ID: 2, Name: "Book 2", Author: "Author 2"},
	}
	suite.bookMock.EXPECT().GetAllBooks().Return(books)

	// Perform the request
	req, err := http.NewRequest("GET", "/books", nil)
	suite.Require().NoError(err)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(suite.controller.GetBook)
	handler.ServeHTTP(recorder, req)

	// Check the response
	suite.Equal(http.StatusOK, recorder.Code)
	suite.Equal("pkglication/json", recorder.Header().Get("Content-Type"))

	var responseBooks []models.Book
	err = json.Unmarshal(recorder.Body.Bytes(), &responseBooks)
	suite.Require().NoError(err)
	suite.ElementsMatch(books, responseBooks)
}

func (suite *BookControllerTestSuite) TestGetBookById() {
	// Prepare mock data and expectations
	bookID := 1
	book := &models.Book{ID: 1, Name: "Book 1", Author: "Author 1"}
	suite.bookMock.EXPECT().GetBookById(int64(bookID)).Return(book, nil)

	// Perform the request
	req, err := http.NewRequest("GET", "/books/1", nil)
	suite.Require().NoError(err)
	req = mux.SetURLVars(req, map[string]string{"bookId": strconv.Itoa(bookID)})

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(suite.controller.GetBookById)
	handler.ServeHTTP(recorder, req)

	// Check the response
	suite.Equal(http.StatusOK, recorder.Code)
	suite.Equal("pkglication/json", recorder.Header().Get("Content-Type"))

	var responseBook models.Book
	err = json.Unmarshal(recorder.Body.Bytes(), &responseBook)
	suite.Require().NoError(err)
	suite.Equal(book, &responseBook)
}

// Add test cases for other controller methods (CreateBook, DeleteBook, UpdateBook) in a similar way.
