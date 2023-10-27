package database

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	
	"main.go/model"
)
type Client struct{
	FirstName	string	`json:"firstname"`
	LastName	string	`json:"lastname"`
	UserName 	string	`json:"username"`
	Location 	string	`json:"location"`
}

//hashing the password

func HashAndSalt(c *fiber.Ctx, password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password),14)
	if err!=nil{
		c.Send(err.Error())
	}
	return string(bytes),nil
}
//sign up endpoint
func CreateUserAccount(c *fiber.Ctx){
	Customers:=model.Customers{}
	if err:=c.BodyParser(&Customers);err!=nil{
		panic(err)
	}
	db:=ConnectDB()
	defer db.Close()
	hashedpassword,err:=HashAndSalt(c,Customers.Password)
	if err!=nil{
		c.Send(err.Error(),"Failed to hash the password")
	}
	db.AutoMigrate(&Customers)
	db.Create(&model.Customers{
		Model: gorm.Model{},
		FirstName: Customers.FirstName,
		LastName: Customers.LastName,
		UserName: Customers.UserName,
		Location: Customers.Location,
		Password: hashedpassword,
	})
	//calling response user function
	Response:=ResponseUser(c,Customers)
	c.JSON(Response)
}
func ResponseUser(c *fiber.Ctx, response model.Customers)Client{
	Response:=Client{
		FirstName: response.FirstName,
		LastName: response.LastName,
		Location: response.Location,
		UserName: response.UserName,
	}
	return Response
}
	

	