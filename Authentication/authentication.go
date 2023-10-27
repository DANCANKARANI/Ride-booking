package authentication

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber"
	"github.com/golang-jwt/jwt"
)

func GenerateJWT(c *fiber.Ctx, user_id uint64)(string,error){
	secretkey := "secretkey"
	var mySigningKey = []byte(secretkey)
	token := jwt.New(jwt.SigningMethodHS256)
	//claims := token.Claims.(jwt.MapClaims)
	
	//create map claims
	claims:=jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	token.Claims=claims
	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		c.Send(err.Error())
		return  "",err
	}
	//setting jwt in cookie
	cookie:=new(fiber.Cookie)
	cookie.Name="jwt"
	cookie.Value=tokenString
	cookie.Expires=time.Now().Add(time.Minute * 30)
	c.Cookie(cookie)
	c.Send(tokenString)
	return tokenString,nil
}

func ExtractClaims(c *fiber.Ctx,tokenString string)(int,error){
	token,err:=jwt.Parse(tokenString,func (token *jwt.Token)(interface{},error)  {
		return []byte("your security key"),nil
	})
	if err!=nil{
		return 0,err
	}
	if claims,ok :=token.Claims.(jwt.MapClaims);ok && token.Valid{
		if id,ok:=claims["id"].(string);ok{
			//convert the Id claim from string to int
			idInt,err:=strconv.Atoi(id)
			if err!=nil{
				return 0,err
			}
			return idInt, nil
		}
	}
	return 0, err
}