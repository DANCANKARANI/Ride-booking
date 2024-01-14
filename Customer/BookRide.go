package customer

import (
	"github.com/gofiber/fiber"
	"main.go/database"
	"main.go/model"
)
type ResponseBooking struct{
	Distance		uint64	`json:"distance"`
	CurrentLocation	string	`json:"currentlocation"`
	Destinition		string	`json:"destinition"`
}

type AvailableDrivers struct{
	Location 	string	`json:"location"`
	Cost		float64	`json:"cost"`
	PhoneNumber	string	`json:"phone_number"`
}

type Order struct{
	FirstName	string	`json:"firstname"`
	Lastname	string	`json:"lastname"`
	Location	string	`json:"location"`
	PhoneNumber	string	`json:"phone_number"`
}
func BookRide(c *fiber.Ctx){
	db:=database.ConnectDB()
	defer db.Close()
	TravellingCustomer:=model.TravellingCustomer{}
	if err:=c.BodyParser(&TravellingCustomer);err!=nil{
		c.Send(err)

	}
	db.AutoMigrate(&TravellingCustomer)
	db.Create(&model.TravellingCustomer{
		Distance: TravellingCustomer.Distance,
		CurrentLocation: TravellingCustomer.CurrentLocation,
		Destinition: TravellingCustomer.Destinition,
	})
	Response:=ResponseBooking{
		Distance: uint64(TravellingCustomer.Distance),
		CurrentLocation: TravellingCustomer.CurrentLocation,
		Destinition: TravellingCustomer.Destinition,
	}
	c.JSON(Response)
}
func CostCalculator(CostPerKm float64,Distance int)float64{
	return	CostPerKm*float64(Distance)
}
func ViewAvailableDrivers(c *fiber.Ctx,location string)AvailableDrivers{
	db:=database.ConnectDB()
	defer db.Close()
	AvailableDrivers:=AvailableDrivers{}
	db.Raw("SELECT * FROM DRIVERS WHERE location=?",location).Scan(&AvailableDrivers)
	return AvailableDrivers
}