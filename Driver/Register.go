package driver

import (
	"github.com/gofiber/fiber"
	"main.go/database"
	"main.go/model"
)
type Response struct{
	FirsName 	string	`json:"firstname"`
	LastName	string	`json:"lastname"`
	UserName	string	`json:"username"`
}
func RegisterDriver(c *fiber.Ctx) {
	Driver:=model.Driver{}

	if err:=c.BodyParser(&Driver);err!=nil{
		panic(err)
	}
	db:=database.ConnectDB()
	defer db.Close()
	db.AutoMigrate(&model.Driver{
		FirstName: Driver.FirstName,
		Lastname: Driver.Lastname,
		UserName: Driver.UserName,
	})
	db.Create(&Driver)
	response:=Response{
		FirsName: Driver.FirstName,
		LastName: Driver.Lastname,
		UserName: Driver.UserName,
	}
	c.JSON(response)
}