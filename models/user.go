package models

import (
	"go-library/database"
	"time"
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Username string `gorm:"unique;not null" json:"username"`
	EncPassword string `gorm:"not null" json:"encPassword"`
	FullName string `gorm:"not null" json:"fullName"`
	Email string `gorm:"not null" json:"email"`
	CreatedBy string `json:"createdBy"`
	UpdatedBy string `json:"updatedBy"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func MigrateUsersTable() {
	db := database.GetDbConnection()
	db.AutoMigrate(&User{})
	fmt.Println("users table migrated")
}