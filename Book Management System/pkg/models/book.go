package models

import (
	"BookManagementSystem/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connection()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}
func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}
func GetBookByID(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID = ?", id).Find(&getBook)
	return &getBook, db
}
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID = ?", ID).Delete(book)
	return book
}
