// Glasses29@
package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connection() {
	database, err := gorm.Open("mysql", "root:Glasses29@tcp(127.0.0.1:3306)/simplerest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = database
}
func GetDB() *gorm.DB {
	return db
}
