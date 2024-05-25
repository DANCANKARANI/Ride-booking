package model

import "github.com/jinzhu/gorm"

type Customers struct {
	gorm.Model
	FirstName 	string 	`json:"firstname"`
	LastName  	string	`json:"lastname"`
	UserName	string	`json:"username"`
	PhoneNumber	string	`json:"phone_number"`
	Password	string	`json:"password"`
}
type Driver struct{
	gorm.Model
	FirstName 	string	`json:"firstname"`
	Lastname 	string	`json:"lastname"`
	UserName	string	`json:"username"`
	PhoneNumber	string	`json:"phone_number"`
	Password 	string	`json:"password"`
	Location	string	`json:"location"`
	
}
type Cars struct{
	gorm.Model
	ID			uint64	`json:"_" gorm:"primaryKey;autoIncrement:true"`
	Location	string	`json:"location"`	
	DriverID	uint64	//it holds the foreign key from the driver
	Driver 		Driver	`gorm:"foreignKey:DriverId"`
}


type Login struct{
	gorm.Model
	UserName 	string 	`json:"username"`
	Password 	string 	`json:"password"`
}
type Admin struct{
	ID 	uint64			`json:"-" gorm:"primaryKey;autoIncrement:true"`
	Email string		`json:"email"`
	Password string		`json:"password"`
}
type AvailableCars struct {
	gorm.Model
	CostPerKm 	float64		`json:"cost"`
	CarID		float64		
	Cars		Cars		`gorm:"foreignKey:CarID"`
}
type TravellingCustomer struct{
	gorm.Model
	Distance uint64			`json:"distance"`
	CurrentLocation	string	`json:"currentlocation"`
	Destinition 	string	`json:"destinition"`
	CustomerID uint64
	Customers Customers `gorm:"foreignKey:CustomerID"`
}
