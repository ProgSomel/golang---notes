package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ecommerce/database"
	"ecommerce/util"
)

func CreateProduct(w http.ResponseWriter, r *http.Request){

	/*
	    1) Take body information(description, imageURL, price, titile) from r.body
		2) Create an instance using Product struct with the body information
		3) Append the instance to the productList
	*/ 

	var newProduct database.Product
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil{
		fmt.Println(err)
		http.Error(w, "Plz give me a valid json", 400)
		return
	}
	newProduct.ID = len(database.ProductList)+1
	database.ProductList = append(database.ProductList, newProduct)

	util.SendData(w, newProduct, 201)
	
}