package main

import (
	
	"fmt"

	"github.com/gofiber/fiber"

	"main.go/Customer"
	"main.go/database"
	"main.go/Driver"
)

func main() {

	fmt.Println("connected")
	app := fiber.New()
	//cutomers endpoints registration
	app.Get("/viewbikes",customer.ViewBikes)
	app.Get("/hi",HelloWorld)
	app.Post("/signup",database.CreateUserAccount)
	app.Post("/login",database.Login)
	app.Post("/book-ride",customer.BookRide)

	//Driver endpoints registration
	app.Post("register-driver",driver.RegisterDriver)
	app.Listen(":8080")
}
func HelloWorld(c *fiber.Ctx){
	c.Send("Hello Dancan")
	
	
}