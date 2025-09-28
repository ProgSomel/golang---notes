package product

import (
	"ecommerce/repo"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

type ReqCreateProduct struct{
	Title string `json:"title"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	ImgURL string `json:"imageURL"`
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request){
	
	var req ReqCreateProduct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil{
		fmt.Println(err)
		http.Error(w, "Plz give me a valid json", 400)
		return
	}

	createdProduct, err := h.productRepo.Create(repo.Product{
		Title: req.Title,
		Description: req.Description,
		Price: req.Price,
		ImgURL: req.ImgURL,
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	util.SendData(w, createdProduct, http.StatusCreated)
	
}