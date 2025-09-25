package main

import (
	"encoding/json"
	"fmt"
	"net/http"
) //?Here net is a package and http is a sub-package under net package



func helloHandler (w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Hello World")
}

func aboutHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "I'm Somel. I'm youtuber. I'm software engineer")
}

type Product struct{
	ID int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	ImgURL string `json:"imageURL"`
}

var productList []Product

func getProducts(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Access-Control-Allow-Origin", "*") //? Here * means everyone can access
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet{ //?or r.Method != "GET"
		http.Error(w, "Please give me GET Request", 400)
		return
	}
	encoder := json.NewEncoder(w) //? it will create an encoder
	encoder.Encode(productList)
}

func main(){
	mux := http.NewServeMux() //? Here mux is a Router. it will return an address of object
	mux.HandleFunc("/hello", helloHandler) //? route
	mux.HandleFunc("/about", aboutHandler) //? route
	mux.HandleFunc("/products", getProducts)
	fmt.Println("Server is running on: 3000")

	err := http.ListenAndServe(":3000", mux) //? return an Error(if there is an Error) or nil(if there is no Error)
	if err != nil{
		fmt.Println("Error starting the server: ", err)
	}
}

func init(){
	prdct1 := Product{
		ID: 1,
		Title: "Orange",
		Description: "Orange is Red. I Love Orange",
		Price: 100,
		ImgURL: "https://www.dole.com/sites/default/files/media/2025-01/oranges.png",
	}
	prdct2 := Product{
		ID: 2,
		Title: "Apple",
		Description: "Apple is Green. I eat Apple",
		Price: 40,
		ImgURL: "https://static.vecteezy.com/system/resources/thumbnails/012/086/172/small_2x/green-apple-with-green-leaf-isolated-on-white-background-vector.jpg",
	}
	prdct3 := Product{
		ID: 3,
		Title: "Banana",
		Description: "Banana is Boring. I hate Banana",
		Price: 5,
		ImgURL: "https://www.dole.com/sites/default/files/media/2025-01/banana-cavendish_0.png",
	}
	productList = append(productList, prdct1)
	productList = append(productList, prdct2)
	productList = append(productList, prdct3)

}