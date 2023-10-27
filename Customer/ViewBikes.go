package customer

import (
	"fmt"

	"github.com/gofiber/fiber"
	"main.go/database"
	"main.go/model"
	//"main.go/model"
	//"main.go/model"
)
type Cars struct{
	ID	uint64	`json:"id"`
	Location	string 	`json:"location"`
}

//view bikes endpoint
//a user will be able to view all bikes available in the database
func ViewBikes(c *fiber.Ctx) {
    db := database.ConnectDB()
    defer db.Close()
    
     Bikes:=[]model.Customers{}
    if err := db.Raw("SELECT * FROM customers").Scan(&Bikes).Error; err != nil {
        c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
        return
    }
	    // Log the query and the results
		fmt.Println("Result:", Bikes)
    
    c.JSON(Bikes)
}