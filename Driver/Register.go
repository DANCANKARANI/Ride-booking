package driver

import (
	"github.com/gofiber/fiber"
	"main.go/database"
	"main.go/model"
)
//What will be send after the driver account is created
type Response struct{
	FirsName 	string	`json:"firstname"`
	LastName	string	`json:"lastname"`
	UserName	string	`json:"username"`
}
//this function registers the Driver to the app
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
		PhoneNumber: Driver.PhoneNumber,
		Location: Driver.Location,
	})
	db.Create(&Driver)
	response:=Response{
		FirsName: Driver.FirstName,
		LastName: Driver.Lastname,
		UserName: Driver.UserName,
	}
	c.JSON(response)
}