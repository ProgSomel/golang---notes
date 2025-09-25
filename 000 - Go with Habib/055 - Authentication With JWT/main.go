package main

import (
	"ecommerce/util"
	"fmt"
)

func main(){
	// secret := []byte("my-secret")
	// message := []byte("Hello World")

	// h := hmac.New(sha256.New, secret) //? Here hmac will use sha256 algorithm to hash
	// h.Write(message)

	// text := h.Sum(nil)

	// fmt.Println(text)

	jwt, err := util.CreateJWT("my-secret", util.Payload{
		Sub: "45",
		FirstName: "Somel",
		LastName: "Ahmed",
		Email: "somelahmed55@gmail.com",
		IsShopOwner: false,
	})

	if err != nil {
		fmt.Println(err)
		return 
	}

	fmt.Println(jwt)

}