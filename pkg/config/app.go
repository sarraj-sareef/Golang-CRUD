package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Db *gorm.DB
)

func InitApp() {
	d, err := gorm.Open("mysql", "sarraj:Password@123@/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	Db = d
}

//
//func GetDB() *gorm.DB {
//	return Db
//}
