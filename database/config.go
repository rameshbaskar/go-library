package database

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

func GetDbConnection() *gorm.DB {
	dbHost := os.Getenv("LIBRARY_DB_HOST")
	dbUser := os.Getenv("LIBRARY_DB_USER")
	dbPassword := os.Getenv("LIBRARY_DB_PASSWORD")
	dbName := os.Getenv("LIBRARY_DB_NAME")
	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbName)

	sqlDb, err := sql.Open("mysql", connString)
	if err != nil {
		panic(err)
	}
	sqlDb.SetConnMaxLifetime(time.Minute * 3)
	sqlDb.SetMaxOpenConns(10)
	sqlDb.SetMaxIdleConns(10)

	dbConnection, _ := gorm.Open(mysql.New(mysql.Config{
  	Conn: sqlDb,
	}), &gorm.Config{})
	
	return dbConnection
}