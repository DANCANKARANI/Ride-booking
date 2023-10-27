package database

import (
	"fmt"

	"github.com/gofiber/fiber"
	"golang.org/x/crypto/bcrypt"
	authentication "main.go/Authentication"
	"main.go/model"
)

//this is a login endpoint
func Login(c *fiber.Ctx){
	Customers:=model.Customers{}
	if err:=c.BodyParser(&Customers);err!=nil{
		c.Send(err.Error())
	}
	db:=ConnectDB()
	var DbUsernames [] string
	db.Table("Customers").Select("user_name").Scan(&DbUsernames)
	IsValid:=CompareUserNames(c,Customers.UserName)
	//checking if the passwword is wrong or correct
	/*if IsValid{
		fmt.Println(Customers.UserName)
		var hashedpassword [] model.Customers
		db.Raw("SELECT PASSWORD FROM CUSTOMERS WHERE User_name=?",Customers.UserName).Scan(&hashedpassword)

		fmt.Println(Customers.Password,"",hashedpassword)
	 for _,Pass:=range hashedpassword{
	 	err:=ComparePasswords(c,Customers.Password,Pass.Password)
		if err!=nil{
			fmt.Println(Customers.Password,"",Pass.Password)
			c.Send("Wrong password")
			fmt.Println("Wrong")
			break
		}else{
			fmt.Println(Customers.Password,"",Pass.Password)
			c.Send("logged in successfully")
		}
		}
	}*/
	if IsValid{
		var HashedPassword [] model.Customers
		db.Raw("select password  from customers where user_name=?",Customers.UserName).Scan(&HashedPassword)
		
		for _,Pass:=range HashedPassword{
			err:=ComparePasswords(c,Customers.Password,Pass.Password)
			if err!=nil{
				
				fmt.Printf("up")
				fmt.Println(Customers.Password," ",Pass.Password)
				c.Send("invalid password")
				break
			}else{
			
				if err:=db.Where("user_name=?",Customers.UserName).First(&Customers).Error;err!=nil{
					panic(err)
				}
				Token, _:=authentication.GenerateJWT(c,uint64(Customers.ID))
				c.BaseURL()
				fmt.Println(Token)
				fmt.Println(Customers.ID)
				
				//c.Send("Logged in successfully")
				
				break			
			}	
		}
	}else{
		c.Send("The account doesn't exist")
	}

	defer db.Close()
	
}
//the function return true when the username is found in the database and false when the username does not exist i  the database
func CompareUserNames(c *fiber.Ctx,username string)bool{
db:=ConnectDB()
return !db.Where("User_name=?",username).First((&model.Customers{})).RecordNotFound()
}

// funtion to compare the entered password if it is correct or wrong
func ComparePasswords(c *fiber.Ctx,plaintext string,hashedPassword string)error{
	passwordbytes:=[]byte(plaintext)
	hashedPasswordBytes:=[]byte(hashedPassword)
	err:=bcrypt.CompareHashAndPassword(hashedPasswordBytes,passwordbytes)
	if err!=nil{
		return err
	}
	return nil
}