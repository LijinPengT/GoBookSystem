package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Book struct {
	Model

	BookID     int    `json:"book_id" gorm:"index"`
	Name       string `json:"name"`
	Org        string `json:"org"`
	Phone      string `json:"phone"`
	Desc       string `json:"desc"`
	Date       string `json:"date"`
	Last       int    `json:"Last"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func ExistBookById(book_id int) bool {
	var book Book
	db.Select("id").Where("book_id = ?", book_id).First(&book)

	if book.ID > 0 {
		return true
	}

	return false
}

func GetBookTotal(maps interface{}) (count int) {
	db.Model(&Book{}).Where(maps).Count(&count)

	return
}

func GetBooks(pageNum int, pageSize int, maps interface{}) (books []Book) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&books)

	return
}

func GetBook(book_id int) (book Book) {
	db.Where("book_id = ?", book_id).First(&book)

	return
}

func EditBook(book_id int, data interface{}) bool {
	db.Model(&Book{}).Where("book_id = ?", book_id).Updates(data)

	return true
}

func AddBook(data map[string]interface{}) bool {
	db.Create(&Book{
		BookID:    data["book_id"].(int),
		Name:      data["name"].(string),
		Org:       data["org"].(string),
		Phone:     data["phone"].(string),
		Date:      data["date"].(string),
		Last:      data["last"].(int),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})

	return true
}

func DeleteBook(book_id int) bool {
	db.Where("book_id = ?", book_id).Delete(Book{})

	return true
}

func (book *Book) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreateOn", time.Now().Unix())

	return nil
}

func (book *Book) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
