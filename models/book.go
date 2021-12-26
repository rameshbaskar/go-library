package models

import (
	"go-library/database"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	ISBN string `gorm:"unique;not null" json:"isbnNumber"`
	Title string `gorm:"not null" json:"bookTitle"`
	Author string `gorm:"not null" json:"bookAuthor"`
	CreatedBy string `json:"createdBy"`
	UpdatedBy string `json:"updatedBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func MigrateBooksTable(table *Book) {
	db := database.GetDbConnection()
	db.AutoMigrate(&table)
	fmt.Println("books table migrated")
}