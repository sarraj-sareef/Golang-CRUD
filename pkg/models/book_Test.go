package models

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type BookTestSuite struct {
	suite.Suite
}

func TestBookTestSuite(t *testing.T) {
	suite.Run(t, new(BookTestSuite))
}
func (ts *BookTestSuite) SetupSuite() {}

func (ts *BookTestSuite) SetupTest() {

}
func (ts *BookTestSuite) TestCreateBook() {
	type testcase struct {
		name       string
		assertions func(*Book)
	}
}
