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

/*package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//hitting the safaricom api end points
	url := "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"
	method := "POST"
	payloadData := map[string]interface{}{
		"BusinessShortCode": 174379,
		"Password": "MTc0Mzc5YmZiMjc5ZjlhYTliZGJjZjE1OGU5N2RkNzFhNDY3Y2QyZTBjODkzMDU5YjEwZjc4ZTZiNzJhZGExZWQyYzkxOTIwMjMxMjE3MTUzMTMy",
		"Timestamp": "20231217153132",
		"TransactionType": "CustomerPayBillOnline",
		"Amount": 1,
		"PartyA": 254708374149,
		"PartyB": 174379,
		"PhoneNumber": 254715118565,
		"CallBackURL": "http://192.168.0.101",
		"AccountReference": "Penta Drive",
		"TransactionDesc": "Bike Ride" ,
	}
//converting the payload into array of bytes
	payloadBytes, err := json.Marshal(payloadData)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	payload := bytes.NewReader(payloadBytes)
	//client for making request to the api
	client := &http.Client{}
	//creating the request to the specified url,storing the results in the request
	req, err := http.NewRequest(method, url, payload)

	//handling errors
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}
	//adding the access token to the request header
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer F77zCa3Zo99ubr9cnt96OGnxDCh3")

	//making request using client.Do
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer res.Body.Close()

	//getting the request results from the request body and storing in the body variable
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
	//printin the result from the body
	fmt.Println(string(body))
}
*/