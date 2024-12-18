// Just returns DB

package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:0101@tcp(127.0.0.1:3306)/bookstore?charset=utf8&parseTime=True&loc=Local") // Opening a connection to the database
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
