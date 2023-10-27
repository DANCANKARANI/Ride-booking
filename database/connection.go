package database

import (
	"fmt"
	_"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)
func ConnectDB()*gorm.DB{
	db, err := gorm.Open("mysql", "root:Karanidancan120@gmail.com@tcp(localhost:3306)/tyson?charset=utf8&parseTime=True")
	if err != nil {
		panic("failed to connect to the database")
	}
	fmt.Println("Connected...")
	return db
}